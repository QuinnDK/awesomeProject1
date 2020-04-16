package main

import (
	"awesomeProject1/text"
	"fmt"
)

//供day.2.1使用的函数类型申明
type cb func(int) int

func main() {
	////day1.1.go
	//var a string = "Dongkun"
	//fmt.Println(a)
	//
	//var c, b int = 1, 3
	//fmt.Println(c, b)
	//
	//var _, number, str = text.Numbers()
	//fmt.Println(number, str)
	//
	////day1.2.go
	//text.Text()
	//
	////day1.3.go使用
	//text.FindSingle()
	//
	////day2.1使用
	//text.TestCallBack(1, text.CallBack)
	//text.TestCallBack(2, func(x int) int {
	//	fmt.Printf("我是回调，x：%d\n", x)
	//	return x
	//})
	//
	////day2.2使用
	//text.Showyanghuisanjiao()
	//
	////day2.3使用
	//add_func := text.Add(1, 2)
	//fmt.Println(add_func())
	//fmt.Println(add_func())
	//fmt.Println(add_func())
	//
	//add_func1 := text.Add1(1, 2)
	//fmt.Println(add_func1(1, 1))
	//fmt.Println(add_func1(0, 0))
	//fmt.Println(add_func1(2, 2))

	/*
		由于前面的练习有点小多了，全部注释掉了
	*/

	//day3.1.go 使用

	var C1 text.Circle
	C1.Radius = 10.00
	fmt.Println("圆的面积：", C1.GetArea())

}
