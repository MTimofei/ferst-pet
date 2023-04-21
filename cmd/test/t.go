package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	var i int
	go func() {
		for {
			i++
			time.Sleep(1 * time.Millisecond)
			fmt.Println(i)
		}
	}()
	go func() {
		for {
			i++
			time.Sleep(1 * time.Millisecond)
			log.Printf("\ti:%d", i)
		}
	}()
	<-time.After(3 * time.Second)
}
