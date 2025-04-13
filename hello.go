package main

import (
	"fmt"
	"time"
)

func greet(name string, channel chan string) {
	time.Sleep(1 * time.Second)
	channel <- fmt.Sprintf("Hello, %s!", name)
}

func greetSlow(name string, channel chan string) {
	time.Sleep(5 * time.Second)
	channel <- fmt.Sprintf("Hello, %s!", name)
}

func main() {
	channel := make(chan string)
	go greet("Old World", channel)
	go greetSlow("New World", channel)
	for i := 0; i < 2; i++ {
		select {
		case message := <-channel:
			fmt.Println(message)
		case <-time.After(20 * time.Second):
			fmt.Println("Timeout!")
		}
	}
}
