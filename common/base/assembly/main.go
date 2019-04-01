package main

import (
	"fmt"
	"os"
)

func main() {
	//s := make([]int, 3, 5)

	//_ = f(s)

	// defer_demo()

	start()
}

func f(s []int) int {
	return s[1]
}

func open_file(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}

	if f != nil {

		defer f.Close()
	}
}

func defer_demo() {
	var whatever [3]struct{}

	for i := range whatever {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}

type number int

func start() {
	var n number

	defer n.print()
	defer n.pprint()
	defer func() {
		n.print()
	}()
	defer func() {
		n.pprint()
	}()

	n = 3
}

func (n number) print() {
	fmt.Println(n)
}

func (n *number) pprint() {
	fmt.Println(*n)
}
