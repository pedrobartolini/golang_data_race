package main

import (
	"time"
	"fmt"
)

var counter int

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}
	}
  
	time.Sleep(time.Second * 2)

	fmt.Println(counter)
}
