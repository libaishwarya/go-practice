package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		channel1 <- fetchAPI1()
	}()

	go func() {
		channel2 <- fetchAPI2()
	}()

	data1 := <-channel1
	data2 := <-channel2

	fmt.Println(data1)
	fmt.Println(data2)
	fmt.Println("All data received")

}

func fetchAPI1() string {
	time.Sleep(2 * time.Second)
	return "Data from API-1"
}

func fetchAPI2() string {
	time.Sleep(2 * time.Second)
	return "Data from API-2"
}
