package main

import (
	"fmt"
	"time"
)

var tarns1 = make(chan int)
var tarns2 = make(chan int)

type t struct {
	A int
	B int
}

func (n *t) f(a1, b1 int) {
	n.A = a1
	n.B = b1
}

func main() {
	var n t
	go func() {
		var i int = 1
		for {
			i++
			tarns1 <- i
			time.Sleep(time.Second)
		}
	}()
	go func() {
		var x int = 1
		for {
			x = x + 9
			tarns2 <- x
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for {
			select {
			case a := <-tarns1:
				n.A = a
			}
		}
	}()
	go func() {
		for {
			select {
			case b := <-tarns2:
				n.B = b
			}
		}
	}()
	for {
		fmt.Println(n.A, n.B)
	}

}
