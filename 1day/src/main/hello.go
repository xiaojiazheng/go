package main // 声明 main 包，表明当前是一个可执行程序

import (
	"fmt"
) // 导入包

const (
	n1 = 6
	n2
	n3
	n4 = 7
	n5
)

//定义常量
var (
	globalA = 3
	globalB = 4
)
var globalC = 5
var globalD, globalE = 6, 7

//定义全局变量
func init() { //init函数，在main函数之前执行，可以做一些准备工作

}
func main() { // main函数，是程序执行的入口
	fmt.Println("预定义的常量:", n1, n2, n3, n4, n5)
	//后续常量值默认为上一个常量的值，第一个为0
	a, b := 3, 4
	//:=只能声明局部变量(定义必须使用）
	c := a + b
	fmt.Println("a + b:", c)
	c = a - b
	fmt.Println("a - b:", c)
	c = a * b
	fmt.Println("a * b:", c)
	c = a / b
	fmt.Println("a / b:", c)
	id, _, sex := getThree() //使用下滑杠代替不想获得的返回值
	fmt.Println("ID:", id, "SEX:", sex)
}
func getThree() (id int, name, sex string) { //在后面附带返回类型
	return 3, "李华", "男"
}
