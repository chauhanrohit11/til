package main

import (
	"fmt"
	"sync"
	"time"
)

func process(process_id int) {
	fmt.Printf("starting process with id %d \n", process_id)
	time.Sleep(2000)
	fmt.Printf("process %d is completed \n", process_id)
}

/*
as per golang docs:
A WaitGroup waits for a collection of goroutines to finish.
The main goroutine calls Add to set the number of goroutines to wait for.
Then each of the goroutines runs and calls Done when finished.
At the same time, Wait can be used to block until all goroutines have finished.
*/
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			process(id)
		}(i)
	}
	wg.Wait()
}
