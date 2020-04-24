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

	//day3.2.go使用
	text.Lingxing()

	//day4.1.go使用
	nums := []int{}
	var i int
	for i = 0; i < 10; i++ {
		nums = text.GetYangHuiTriangleNextLine(nums)
		fmt.Println(nums)
	}
	fmt.Println()
	text.Triangle(13)

	//day5.1.go使用
	/* 数组长度为 5 */
	fmt.Println()
	var balance = [5]float32{1000, 2, 3, 17, 50}
	var avg float32

	/* 数组作为参数传递给函数 */
	avg = text.GetAverage(balance, 5)

	/* 输出返回的平均值 */
	fmt.Printf("平均值为: %f ", avg)
	fmt.Println()

	//day6.1使用
	var book1 text.Books
	book1.Title = "book1"
	book1.Auther = "Dongkun"
	book1.Book_id = 1
	text.Changebook(book1) //传递实参并不能改变其参数值
	fmt.Println(book1)
	text.Changebook1(&book1)
	fmt.Println(book1)

	//day7.1.go 使用
	fmt.Println()
	//var numbers []int

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	text.PrintSlice(numbers)
	text.PrintSlice(numbers[0:4]) //len为实际长度，cap为左标到末端的容量
	text.PrintSlice(numbers[1:4])

	fmt.Println()
	//day8.1.go使用
	text.SliceTest()
	text.TwoDimensionArray()

	//day8.2.go
	text.ChangeSliceTest()

	//day9.1.go
	text.Makemap()
	//day9.2.go使用
	text.GetInstance()
	text.Put("a", "a_put")
	text.Put("b", "b_put")
	fmt.Println(text.Get("a"))
	fmt.Println(text.Get("b"))
	text.Put("p", "p_put")
	fmt.Println(text.Get("p"))

	fmt.Println("-----------------------")

	//day10.1.go 使用
	var x int = 15
	fmt.Printf("%d 的阶乘是 %d\n", x, text.Factorial(uint64(x)))
	fmt.Println()
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", text.Fibonacci(i))
	}
	fmt.Println()
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", text.Fibonacci1(i))
	}

	//day10.2.go 使用
	fmt.Println()
	fmt.Println(text.Get_sqrt(2))
	fmt.Println(text.Get_sqrt(3))

	//day11.1.go使用
	fmt.Println()
	var phone text.Phone //接口定义

	phone = new(text.NokiaPhone)
	phone.Call()

	phone = new(text.IPhone)
	phone.Call()

	phone = new(text.Int)
	phone.Call()

	fmt.Println()
	//day11.2.go
	var phone11 text.PPhone

	//先实例化第一个接口
	phone11 = new(text.Phone1)
	phone11.Call1()
	phone11.Call2()

	//实例化第二个接口
	phone11 = new(text.Phone2)
	phone11.Call1()
	phone11.Call2()

	fmt.Println()
	//day11.3.go
	var dageda text.PPPhone
	dageda = new(text.Huawei)
	dageda.Takephoto()
	r := dageda.Call3(20)
	fmt.Println(r)

}
