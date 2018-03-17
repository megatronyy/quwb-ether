package main

import "fmt"
import "quwb-ether/learn"

func main() {


	fmt.Println(learn.Plus(123.456, 456.789))

	fmt.Println(learn.Swap("demon", "18"))

	swapPointer()

	fmt.Println(learn.GetSequence()())

	CircleDemo()

	factorialDemo()
}

//指针
func swapPointer(){
	fmt.Println("----指针引用-----")
	var x int = 10
	var y int = 20

	fmt.Println("未进入函数中的x值:", x)
	fmt.Println("未进入函数中的y值:", y)
	fmt.Println("----------------------")
	learn.SwapPointer(&x, &y)

	fmt.Println("*********************")
	fmt.Println("全局变量x的值为:", x)
	fmt.Println("全局变量y的值为:", y)
}

//Go 语言函数作为值(匿名函数的使用)
func anonFunc(){
	fmt.Println("----匿名函数-----")
	//匿名函数定义
	f := func(x int) string {
		return fmt.Sprintf("%d * %d = %d", x, x, x * x)
	}

	fmt.Println(f(2))
}

func CircleDemo()  {
	fmt.Println("----匿名函数-----")
	//给结构体赋值
	//c1 := learn.Circle{ Redius:100.00 }

	//给结构体赋值的另外一种写法
	var c1 learn.Circle
	c1.Redius = 100.00

	fmt.Printf("Area of Circle(c1) = %.2f ", c1.GetArea())
}

func factorialDemo()  {
	fmt.Println("----阶乘-----")
	fmt.Printf("15的阶乘为:%d ",learn.Factorial(15))
}

