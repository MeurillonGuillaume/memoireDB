package shared

import (
	"context"
	"testing"
	"time"
)

func TestCombineChans(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	intStream := make(chan interface{})
	floatStream := make(chan interface{})

	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				floatStream <- i
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				floatStream <- float64(i)
			}
		}
	}()

	outStream := CombineChans(ctx, intStream, floatStream)
	for item := range outStream {
		switch item.(type) {
		case int:
			t.Logf("%d (%T)", item, item)
		case float64:
			t.Logf("%f (%T)", item, item)
		default:
			t.Fatalf("Got unexpected item %v of type %T", item, item)
		}

	}
}
