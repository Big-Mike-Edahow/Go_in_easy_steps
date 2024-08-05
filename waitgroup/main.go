/* main.go */

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go report(i, &wg)
	}
	
	wg.Wait()
}

func report(i int, wg *sync.WaitGroup) {
	fmt.Println("Goroutine", i, "started.")
	time.Sleep(1 * time.Second)
	fmt.Println("Goroutine", i, "ended.")
	wg.Done()
}
