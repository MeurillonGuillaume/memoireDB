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
		var i int
		for {
			select {
			case <-ctx.Done():
				return
			default:
				intStream <- i
			}
			i++
		}
	}()

	go func() {
		var i float64
		for {
			select {
			case <-ctx.Done():
				return
			default:
				floatStream <- i
			}
			i++
		}
	}()

	outStream := CombineChans(ctx, intStream, floatStream)
	for item := range outStream {
		t.Log(item)
	}
}
