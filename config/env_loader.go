/* Code duplication from https://github.com/koding/multiconfig/blob/master/env.go
With following changes:
  - Drop support for camelcase
  - Add support for default value
*/

package config

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/koding/multiconfig"
)

type EnvLoader struct {
	Prefix string
}

var _ multiconfig.Loader = (*EnvLoader)(nil)

// Load loads the source into the config defined by struct s
func (e *EnvLoader) Load(s interface{}) error {
	strct := structs.New(s)
	strctMap := strct.Map()
	prefix := e.getPrefix(strct)

	for key, val := range strctMap {
		field := strct.Field(key)

		if err := e.processField(prefix, field, key, val); err != nil {
			return err
		}
	}

	return nil
}

func (e *EnvLoader) getPrefix(s *structs.Struct) string {
	if e.Prefix != "" {
		return e.Prefix
	}

	return s.Name()
}

// processField gets leading name for the env variable and combines the current
// field's name and generates environment variable names recursively
func (e *EnvLoader) processField(prefix string, field *structs.Field, name string, strctMap interface{}) error {
	fieldName := e.generateFieldName(prefix, name)

	strctMap, ok := strctMap.(map[string]interface{})
	if ok {
		for key, val := range strctMap.(map[string]interface{}) {
			field := field.Field(key)

			if err := e.processField(fieldName, field, key, val); err != nil {
				return err
			}
		}
	} else {
		v := os.Getenv(fieldName)
		if v == "" {
			if tmp := field.Tag(_fieldDefault); tmp != "" {
				v = tmp
			} else {
				return nil
			}
		}

		if err := fieldSet(field, v); err != nil {
			return err
		}
	}

	return nil
}

// PrintEnvs prints the generated environment variables to the std out.
func (e *EnvLoader) PrintEnvs(s interface{}) {
	strct := structs.New(s)
	strctMap := strct.Map()
	prefix := e.getPrefix(strct)

	keys := make([]string, 0, len(strctMap))
	for key := range strctMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		field := strct.Field(key)
		e.printField(prefix, field, key, strctMap[key])
	}
}

// printField prints the field of the config struct for the flag.Usage
func (e *EnvLoader) printField(prefix string, field *structs.Field, name string, strctMap interface{}) {
	fieldName := e.generateFieldName(prefix, name)

	smap, ok := strctMap.(map[string]interface{})
	if ok {
		keys := make([]string, 0, len(smap))
		for key := range smap {
			keys = append(keys, key)
		}

		sort.Strings(keys)
		for _, key := range keys {
			field := field.Field(key)
			e.printField(fieldName, field, key, smap[key])
		}
	} else {
		fmt.Println("  ", fieldName)
	}
}

// generateFieldName generates the field name combined with the prefix and the
// struct's field name
func (e *EnvLoader) generateFieldName(prefix string, name string) string {
	fieldName := strings.ToUpper(name)
	return strings.ToUpper(prefix) + "_" + fieldName
}

// fieldSet sets field value from the given string value. It converts the
// string value in a sane way and is usefulf or environment variables or flags
// which are by nature in string types.
func fieldSet(field *structs.Field, v string) error {
	switch f := field.Value().(type) {
	case flag.Value:
		if v := reflect.ValueOf(field.Value()); v.IsNil() {
			typ := v.Type()
			if typ.Kind() == reflect.Ptr {
				typ = typ.Elem()
			}

			if err := field.Set(reflect.New(typ).Interface()); err != nil {
				return err
			}

			f = field.Value().(flag.Value)
		}

		return f.Set(v)
	}

	// TODO: add support for other types
	switch field.Kind() {
	case reflect.Bool:
		val, err := strconv.ParseBool(v)
		if err != nil {
			return err
		}

		if err := field.Set(val); err != nil {
			return err
		}
	case reflect.Int:
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}

		if err := field.Set(i); err != nil {
			return err
		}
	case reflect.String:
		if err := field.Set(v); err != nil {
			return err
		}
	case reflect.Slice:
		switch t := field.Value().(type) {
		case []string:
			if err := field.Set(strings.Split(v, ",")); err != nil {
				return err
			}
		case []int:
			var list []int
			for _, in := range strings.Split(v, ",") {
				i, err := strconv.Atoi(in)
				if err != nil {
					return err
				}

				list = append(list, i)
			}

			if err := field.Set(list); err != nil {
				return err
			}
		default:
			return fmt.Errorf("multiconfig: field '%s' of type slice is unsupported: %s (%T)",
				field.Name(), field.Kind(), t)
		}
	case reflect.Float64:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}

		if err := field.Set(f); err != nil {
			return err
		}
	case reflect.Int64:
		switch t := field.Value().(type) {
		case time.Duration:
			d, err := time.ParseDuration(v)
			if err != nil {
				return err
			}

			if err := field.Set(d); err != nil {
				return err
			}
		case int64:
			p, err := strconv.ParseInt(v, 10, 0)
			if err != nil {
				return err
			}

			if err := field.Set(p); err != nil {
				return err
			}
		default:
			return fmt.Errorf("multiconfig: field '%s' of type int64 is unsupported: %s (%T)",
				field.Name(), field.Kind(), t)
		}

	default:
		return fmt.Errorf("multiconfig: field '%s' has unsupported type: %s", field.Name(), field.Kind())
	}

	return nil
}
