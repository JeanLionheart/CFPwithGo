package pipeline

import (
	"gcm/ordonech"
)

type SourceFunc func() any

type HandleFunc func(any) any

func Generate(done <-chan any, src SourceFunc) <-chan any {
	generator := make(chan any)
	go func() {
		defer close(generator)
		for true {
			select {
			case <-done:
				return
			case generator <- src():
			}
		}
	}()
	return generator
}

func Handle(done <-chan any, srcChan <-chan any, hdl HandleFunc) <-chan any {
	results := make(chan any)
	go func() {
		defer close(results)
		odc := ordonech.New(done, srcChan)
		for e := range odc {
			q := hdl(e)
			select {
			case <-done:
				return
			case results <- q:
			}
		}
	}()
	return results
}
