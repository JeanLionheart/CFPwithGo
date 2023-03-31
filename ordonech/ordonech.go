package ordonech

/*
ch is a individual channel,
we don't know if it can be closed be a defer_close+done(when for range ch)
(if done but ch not close, the goroutine will leak)

so we make it can be closed by done.
*/
func New(done, ch <-chan any) <-chan any {
	valStream := make(chan any)
	go func() {
		defer close(valStream)
		for true {
			select {
			case <-done:
				return
			case v, ok := <-ch:
				if !ok {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}
