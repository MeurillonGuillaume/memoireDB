package shared

import "context"

// CombineChans will combine any number of input channels to a single output channel of interfaces.
func CombineChans(ctx context.Context, chs ...<-chan interface{}) <-chan interface{} {
	outStream := make(chan interface{})

	for _, inStream := range chs {
		go forwardStream(ctx, inStream, outStream)
	}

	go func() {
		<-ctx.Done()
		close(outStream)
	}()

	return outStream
}

// forwardStream will loop over an input channel as long as its Context is alive and forward its values to an output channel.
func forwardStream(ctx context.Context, in <-chan interface{}, out chan<- interface{}) {
	for {
		select {
		case <-ctx.Done():
			return
		case item := <-in:
			select {
			case <-ctx.Done():
				return
			case out <- item:
			}
		}
	}
}
