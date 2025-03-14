package main

import (
	"fmt"
	"time"
)

func main() {
	go fetchAPI1()
	go fetchAPI2()

	time.Sleep(2 * time.Second)
	fmt.Println("All done")

}

func fetchAPI1() {
	time.Sleep(1 * time.Second)
	fmt.Println("API-1 done")
}

func fetchAPI2() {
	time.Sleep(1 * time.Second)
	fmt.Println("API-2 done")
}
