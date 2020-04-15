package text

import "fmt"

type cb func(int) int

//函数作为参数回调

func TestCallBack(x int, f cb) {
	f(x)
}

func CallBack(x int) int {
	fmt.Printf("我是回调，x：%d\n", x)
	return x
}
