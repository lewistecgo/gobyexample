package main

import (
	"fmt"
	"sync"
	"time"
)

func workerw(id int) {
	fmt.Printf("Worker %d starting\n\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n\n", id)
}
func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i

		go func() {
			defer wg.Done()
			workerw(i)
		}()
	}
	wg.Wait()
}
