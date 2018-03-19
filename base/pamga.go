package base

import (
	"fmt"
)

func Run(){
	pointer()
}

func pointer()  {
	x := 1
	p := &x

	fmt.Printf("x变量的地址是：%s\n", p)
	fmt.Printf("x变量的值是：%s\n", *p)

	*p = 2

	fmt.Printf("x变量的地址是：%s\n", *p)
}

