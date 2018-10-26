package main

import (
	"strconv"
	"reflect"
	"fmt"
)

/**
 * reflect设计的select语法
**/
func main() {
	var (
		chs1     = make(chan int)
		chs2     = make(chan float64)
		chs3     = make(chan string)
		ch4close = make(chan int)
	)
	defer close(ch4close)

	go func(c chan int, ch4close chan int) {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
		ch4close <- 1
	}(chs1, ch4close)

	go func(c chan float64, ch4close chan int) {
		for i := 0; i < 5; i++ {
			c <- float64(i) + 0.1
		}
		close(c)
		ch4close <- 1
	}(chs2, ch4close)

	go func(c chan string, ch4close chan int) {
		for i := 0; i < 5; i++ {
			c <- "string:" + strconv.Itoa(i)
		}
		close(c)
		ch4close <- 1
	}(chs3, ch4close)

	var selectcase = make([]reflect.SelectCase, 4)
	selectcase[0].Dir = reflect.SelectRecv
	selectcase[0].Chan = reflect.ValueOf(chs1)

	selectcase[1].Dir = reflect.SelectRecv
	selectcase[1].Chan = reflect.ValueOf(chs2)

	selectcase[2].Dir = reflect.SelectRecv
	selectcase[2].Chan = reflect.ValueOf(chs3)

	selectcase[3].Dir = reflect.SelectRecv
	selectcase[3].Chan = reflect.ValueOf(ch4close)

	done := 0
	finished := 0
	for finished < len(selectcase)-1 {
		chosen, recv, recvok := reflect.Select(selectcase)
		if recvok {
			done = done + 1
			switch chosen {
			case 0:
				fmt.Println(chosen, recv.Int())
			case 1:
				fmt.Println(chosen, recv.Float())
			case 2:
				fmt.Println(chosen, recv.String())
			case 3:
				finished = finished + 1
				done = done - 1
			}
		}
	}
	fmt.Println("Done", done)
}
