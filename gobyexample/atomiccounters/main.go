package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops atomic.Uint64
	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)

		ops.Add(uint64(i))
		wg.Done()

		//go func() {
		//	for c := 0; c < 1000; c++ {
		//		ops.Add(1)
		//	}
		//
		//	wg.Done()
		//}()
	}

	wg.Wait()
	fmt.Println("ops: ", ops.Load())
}
