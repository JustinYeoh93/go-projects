package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	tasks := []string{"task1", "task2", "task3", "task4", "task5"}

	results := make(chan string)
	var wg sync.WaitGroup

	// Fan-out: Distribute tasks
	for _, task := range tasks {
		wg.Add(1)
		go func(task string) {
			defer wg.Done()
			result := processTask(task)
			results <- result
		}(task)
	}

	// Fan-in: Collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print results
	for result := range results {
		fmt.Println("Result:", result)
	}
}

func processTask(task string) string {
	sleepTime := rand.Intn(5)
	fmt.Println("Processing:", task, "will sleep for", sleepTime, "seconds")
	time.Sleep(time.Duration(sleepTime) * time.Second)
	return task + " completed"
}
