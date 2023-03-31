package orchannel

import (
	"fmt"
	"testing"
	"time"
)

const CHNUM = 10

func TestWait(t *testing.T) {
	chs := make([]chan any, CHNUM)
	for i, _ := range chs {
		chs[i] = make(chan any)
	}
	produce := func(i int) {
		// time.Sleep(5 * time.Second)
		defer close(chs[i])
		for k := 0; k < (i+1)*CHNUM; k++ {
			time.Sleep(time.Second * 1)
		}
	}
	for i, _ := range chs {
		go produce(i)
	}

	chs_ := make([]<-chan any, CHNUM)
	for i, _ := range chs_ {
		chs_[i] = chs[i]
	}

	a := time.Now()
	<-Wait(chs_...)
	if time.Since(a) < time.Second*CHNUM {
		fmt.Println(time.Since(a))
		t.Error()
	}
}
