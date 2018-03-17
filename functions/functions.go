package functions

import (
	"fmt"
)

func Run()  {
	fmt.Println(plus(123.456, 456.789))
	fmt.Println(swap("demon", "18"))
	swapPointer()
	fmt.Println(getSequence()())
	circleDemo()
	factorialDemo()
}

//自己定义的求合函数
func plus(num1, num2 float32) float32 {
	result := num1 + num2
	return result
}

//返回多个值的函数
func swap(x, y string) (string, string) {
	fmt.Println("----返回多个值的函数-----")
	return y, x
}

//指针
func swapPointer(){
	fmt.Println("----指针引用-----")
	var x int = 10
	var y int = 20

	fmt.Println("未进入函数中的x值:", x)
	fmt.Println("未进入函数中的y值:", y)
	fmt.Println("----------------------")
	swapPointerFun(&x, &y)

	fmt.Println("*********************")
	fmt.Println("全局变量x的值为:", x)
	fmt.Println("全局变量y的值为:", y)
}

func swapPointerFun(x, y *int) string{
	*x += 50 //*x是为了保持x地址上的值
	*y += 50 //*x是为了保持x地址上的值

	fmt.Println("swap函数内部的x值为: ", *x)
	fmt.Println("swap函数内部的y值为: ", *y)

	return ""
}

//Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。
//匿名函数的优越性在于可以直接使用函数内的变量，不必申明
func getSequence() func() int  {

	fmt.Println("----闭包-----")

	i := 0

	return func() int {
		i+=1
		return i
	}
}

//方法
// Go 语言中同时有函数和方法。
// 一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
// 所有给定类型的方法属于该类型的方法集。
func circleDemo()  {
	fmt.Println("----匿名函数-----")
	//给结构体赋值
	//c1 := functions.Circle{ Redius:100.00 }

	//给结构体赋值的另外一种写法
	var c1 circle
	c1.redius = 100.00

	fmt.Printf("Area of Circle(c1) = %.2f ", c1.getArea())
}

type circle struct {
	redius float64
}

func (c circle) getArea() float64  {
	return 3.14 * c.redius
}

//递归
//阶乘
func factorialDemo()  {
	fmt.Println("----阶乘-----")
	fmt.Printf("15的阶乘为:%d\n ", factorial(15))
}

func factorial(n int) (result int){
	if n == 0{
		result = 1
	}else{
		result = n * factorial(n - 1)
	}

	return result
}

