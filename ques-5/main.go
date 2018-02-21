package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello!!!")
	fmt.Println("Test Started...")
	ch := make(chan string)
	go infinityWarfare(ch)
	for {
		tmp := <-ch
		fmt.Println("From Receiver:-----------------------")
		fmt.Println(tmp)
	}
}

func infinityWarfare(ch chan string) {
	var str string
	for {
		time.Sleep(5)
		fmt.Scanf("%s", &str)
		ch <- str
	}
}
