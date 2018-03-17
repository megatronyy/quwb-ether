package learn

import (
	"fmt"
)

var (
	//声明全局变量num1, num2
	num1 float32 = 123.456
	num2 float32 = 456.789
)

//自己定义的求合函数
func Plus(num1, num2 float32) float32 {
	result := num1 + num2
	return result
}

//返回多个值的函数
func Swap(x, y string) (string, string) {
	fmt.Println("----返回多个值的函数-----")
	return y, x
}

func SwapPointer(x, y *int) string{
	*x += 50 //*x是为了保持x地址上的值
	*y += 50 //*x是为了保持x地址上的值

	fmt.Println("swap函数内部的x值为: ", *x)
	fmt.Println("swap函数内部的y值为: ", *y)

	return ""
}

//Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。
//匿名函数的优越性在于可以直接使用函数内的变量，不必申明
func GetSequence() func() int  {

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
type Circle struct {
	Redius float64
}

func (c Circle) GetArea() float64  {
	return 3.14 * c.Redius
}

//递归
//阶乘
func Factorial(n int) (result int){
	if n == 0{
		result = 1
	}else{
		result = n * Factorial(n - 1)
	}

	return result
}