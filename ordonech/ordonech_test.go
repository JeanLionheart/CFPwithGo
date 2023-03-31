package ordonech

import (
	"testing"
	"time"
)

func TestOD(t *testing.T) {
	ch := make(chan any)
	done := make(chan any)
	go func() {
		defer close(done)
		time.Sleep(time.Second * 3)
	}()

	a := time.Now()
	<-New(done, ch)
	if time.Since(a) < time.Second*3 {
		t.Error()
	}

}
