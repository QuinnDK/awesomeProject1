package main

import (
	"awesomeProject1/text"
	"fmt"
)

func main()  {
	var a string = "Dongkun"
	fmt.Println(a)

	var c, b int = 1, 3
	fmt.Println(c,b)

	var _, number, str = text.Numbers()
	fmt.Println(number,str)

	text.Text()

}
