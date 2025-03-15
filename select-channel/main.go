package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	channel3 := make(chan string)

	go func() {
		channel1 <- fetchAPI1()
	}()

	go func() {
		channel2 <- fetchAPI2()
	}()

	go func() {
		channel3 <- fetchAPI3()
	}()

	for i := 0; i < 3; i++ {
		select {
		case data1 := <-channel1:
			fmt.Println(data1)
		case data2 := <-channel2:
			fmt.Println(data2)
		case data3 := <-channel3:
			fmt.Println(data3)

		}
	}
	fmt.Println("Alldata received")
}

func fetchAPI1() string {
	time.Sleep(1 * time.Second)
	return "Data of API-1"
}

func fetchAPI2() string {
	time.Sleep(1 * time.Second)
	return "Data of API-2"

}

func fetchAPI3() string {
	time.Sleep(1 * time.Second)
	return "Data of API-3"

}
