package main

import (
	"fmt"
	"time"
)

func main() {
	//c := make(chan bool, 1)
	//c := make(chan string, 1)
	c := make(chan int, 1)

	go func() {
		time.Sleep(time.Second)
		//c <- true
		close(c)
	}()

	fmt.Println(<-c)
}
