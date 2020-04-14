package main

import (
	"awesomeProject1/text"
	"fmt"
)

func main() {
	//day1.1.go
	var a string = "Dongkun"
	fmt.Println(a)

	var c, b int = 1, 3
	fmt.Println(c, b)

	var _, number, str = text.Numbers()
	fmt.Println(number, str)

	//day1.2.go
	text.Text()

	//day1.3.go使用
	text.FindSingle()

}
