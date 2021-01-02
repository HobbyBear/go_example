package main

import (
	"sync"
	"time"
)

func main() {


	time.Sleep(2 * time.Second)
}

func test() {

	var wg sync.WaitGroup
	for i := 0; i< 100; i++{
		wg.Add(1)
		go func() {
			wg.Done()

		}()
	}
	wg.Wait()

}
