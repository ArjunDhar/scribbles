package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[main] panic occurred:", err)
		}
	}()

	watchDog := func() {
		fmt.Println("[watchDog] start")
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("[watchdog] panic occurred:", err)
			}
		}()
		// Watchdog

		{
			// Your code here that checks for End condition
			// Enxample waiting for INDEXING to finish in the DB for all indexes etc.
			time.Sleep(10 * time.Second) // only for test to run
		}

		fmt.Println("[watchDog] end")
		wg.Done()
	}

	go watchDog()
	time.Sleep(1 * time.Second) // ensures the Go scheduler executes the watchdog first

	wg.Add(1)
	{
		go func() { fmt.Println("Asynch Task 1 fired") }()
		go func() { fmt.Println("Asynch Task 2 fired") }()
		go func() { fmt.Println("Asynch Task 3 fired") }()
		// ...
	}

	wg.Wait()
	fmt.Println("~~~ End program ~~~")
}
