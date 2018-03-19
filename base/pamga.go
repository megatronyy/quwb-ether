package base

import (
	"fmt"
)

func Run(){
	pointer()
	any(333)
	any("666")
}

func pointer()  {
	x := 1
	p := &x

	fmt.Printf("x变量的地址是：%s\n", p)
	fmt.Printf("x变量的值是：%s\n", *p)

	*p = 2

	fmt.Printf("x变量的地址是：%s\n", *p)
}

func any(v interface{}){
	fmt.Println("interface{}参数验证")
	if v1,ok := v.(string); ok{
		fmt.Println(v1)
	}else if v1, ok := v.(int); ok{
		fmt.Println(v1)
	}
}

