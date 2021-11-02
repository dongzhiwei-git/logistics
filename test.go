package main

import (
	"fmt"
)

var Ch = make(chan int)

func main() {
	go func() {
		one()
	}()
	two()
}
func two() {
	for i := 0; i < 100; i = i + 2 {
		Ch <- 1

		fmt.Println("偶数", i)

	}
}

func one() {
	for i := 1; i <= 100; i = i + 2 {
		<-Ch
		fmt.Println("奇数", i)

	}

}
