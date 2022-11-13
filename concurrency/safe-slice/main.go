package main

import (
	"fmt"
	"sync"
)

func done() {
	type answer struct {
		mu   sync.Mutex
		data []int
	}

	var odd answer
	var even answer
	var wg = sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		if i%2 == 0 {
			go func(i int) {
				defer wg.Done()
				even.mu.Lock()
				even.data = append(even.data, i)
				even.mu.Unlock()
			}(i)
		} else {
			go func(i int) {
				defer wg.Done()
				odd.mu.Lock()
				odd.data = append(odd.data, i)
				odd.mu.Unlock()
			}(i)
		}
	}
	wg.Wait()
	fmt.Println(odd.data)
	fmt.Println(even.data)
}

func main() {
	for i := 0; i < 9; i++ {
		fmt.Println("=======")
		done()
	}
}
