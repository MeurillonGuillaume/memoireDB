package datastore

import (
	"fmt"
	"testing"

	"github.com/MeurillonGuillaume/memoireDB/external/communication/model"
	"github.com/stretchr/testify/assert"
)

func TestLoadKeyValue(t *testing.T) {
	ds := newMemoryDatastore()
	defer func() {
		assert.NoError(t, ds.Close())
	}()

	testCases := []struct {
		setKey, getKey             string
		newValue, expectValue      interface{}
		expectSetErr, expectGetErr error
	}{
		{
			setKey:      "int",
			getKey:      "int",
			newValue:    123,
			expectValue: 123,
		},
		{
			setKey:      "int64",
			getKey:      "int64",
			newValue:    int64(123),
			expectValue: int64(123),
		},
		{
			setKey:      "string",
			getKey:      "string",
			newValue:    "a string",
			expectValue: "a string",
		},
		{
			setKey:      "float64",
			getKey:      "float64",
			newValue:    10.2,
			expectValue: 10.2,
		}, {
			setKey:      "float32",
			getKey:      "float32",
			newValue:    float32(10.2),
			expectValue: float32(10.2),
		},
		{
			setKey: "structs",
			getKey: "structs",
			newValue: struct {
				someKey string
				someInt int
			}{"key", 1},
			expectValue: struct {
				someKey string
				someInt int
			}{"key", 1},
		},
		{
			setKey:       "hello, world",
			getKey:       "other key",
			newValue:     0,
			expectValue:  1,
			expectSetErr: nil,
			expectGetErr: ErrNoSuchKey,
		},
	}

	for _, testCase := range testCases {
		var (
			err     error
			current interface{}
		)
		_, err = ds.StoreKeyValue(model.InsertModel{
			Key:   testCase.setKey,
			Value: testCase.newValue,
		})
		assert.Equal(t, testCase.expectSetErr, err)

		if testCase.expectSetErr == nil && testCase.expectGetErr == nil {
			current, err = ds.LoadKeyValue(model.RetrieveModel{
				Key: testCase.getKey,
			})
			assert.Equal(t, testCase.expectGetErr, err)
			assert.Equal(t, testCase.expectValue, current)
		}
	}
}

func TestListKeys(t *testing.T) {
	ds := newMemoryDatastore()
	defer func() {
		assert.NoError(t, ds.Close())
	}()

	currentKeys := make([]string, 10)
	for i := 0; i < len(currentKeys); i++ {
		currentKeys[i] = fmt.Sprintf("key-%d", i)
		_, err := ds.StoreKeyValue(model.InsertModel{
			Key:   currentKeys[i],
			Value: i,
		})
		assert.NoError(t, err)
	}

	testCases := []struct {
		queryPrefix  string
		expectedErr  error
		expectedKeys []string
	}{
		{
			queryPrefix: "   ",
			expectedErr: ErrPrefixWhitespace,
		},
		{
			queryPrefix: "invalid",
			expectedErr: ErrNoSuchKey,
		},
		{
			queryPrefix:  "key-",
			expectedKeys: currentKeys,
		},
		{
			expectedKeys: currentKeys,
		},
		{
			queryPrefix: "key-1",
			expectedKeys: []string{
				"key-1",
			},
		},
	}

	for _, testCase := range testCases {
		keys, err := ds.ListKeys(model.ListKeysModel{
			Prefix: testCase.queryPrefix,
		})

		assert.Equal(t, testCase.expectedErr, err)
		assert.Equal(t, testCase.expectedKeys, keys)
	}
}
