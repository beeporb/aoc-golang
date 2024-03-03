package main

import (
	"fmt"
	"sync"
)

func sendInt(i int, ints chan int, wg *sync.WaitGroup) {
	ints <- i
	return
}

func totalInt(ints chan int, wg *sync.WaitGroup) int {

	total := 0

	go func() {
		for i := range ints {
			total += i
			wg.Done()
		}
	}()

	wg.Wait()

	return total
}

func main() {

	ints := make(chan int)

	wg := new(sync.WaitGroup)

	for i := 1; i <= 2500; i++ {
		wg.Add(1)
		go sendInt(i, ints, wg)

	}

	total := totalInt(ints, wg)

	fmt.Println(total)
}
