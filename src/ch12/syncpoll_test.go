package ch12

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			return 100

		},
	}

	go func() {
		for {
			pool.Put(rand.Int())
		}
	}()

	go func() {
		for {
			var v = pool.Get().(int)
			t.Logf("get value is %d", v)
		}
	}()
	time.Sleep(time.Second * 3)
}
