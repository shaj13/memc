package memc_test

import (
	"fmt"
	"time"

	"github.com/shaj13/memc"
	"github.com/shaj13/memc/container/fifo"
	_ "github.com/shaj13/memc/container/idle"
	"github.com/shaj13/memc/container/lfu"
	"github.com/shaj13/memc/container/lru"
)

func Example_idle() {
	//  it can be unsafe, no any race conditions
	c := memc.IDLE.NewUnsafe()
	c.Store(1, 0)
	fmt.Println(c.Contains(1))
	// Output:
	// false
}

func Example_fifo() {
	cap := fifo.Capacity(2)
	c := memc.FIFO.New(cap)
	c.Store(1, 0)
	c.Store(2, 0)
	c.Store(3, 0)
	fmt.Println(c.Contains(1))
	// Output:
	// false
}

func Example_lru() {
	cap := lru.Capacity(2)
	c := memc.LRU.New(cap)
	c.Store(1, 0)
	c.Store(2, 0)
	c.Store(3, 0)
	fmt.Println(c.Contains(1))
	// Output:
	// false
}

func Example_lfu() {
	cap := lfu.Capacity(2)
	c := memc.LFU.New(cap)
	c.Store(1, 0)
	c.Store(2, 0)
	c.Load(1)
	c.Store(3, 0)
	fmt.Println(c.Contains(2))
	// Output:
	// false
}

func Example_onexpired() {
	// c must be thread safe
	var c memc.Cache

	exp := lru.RegisterOnExpired(func(key interface{}) {
		// use Peek/Load over delete, perhaps a new entry added with the same key during expiration,
		// or entry refreshed from other thread.
		c.Peek(key)
	})

	c = memc.LRU.New(exp)
	c.SetTTL(time.Millisecond)
	c.Store(1, 0)

	time.Sleep(time.Millisecond * 5)
	fmt.Println(c.Contains(1))
	// Output:
	// false
}
