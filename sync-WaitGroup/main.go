package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	fetchAPI1(&wg)
	fetchAPI2(&wg)

	wg.Wait()
	fmt.Println("All done")
}

func fetchAPI1(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("API-1 done")

}

func fetchAPI2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("API-2 done")

}
