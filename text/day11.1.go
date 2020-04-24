package text

///* 定义接口 */
//type interface_name interface {
//	method_name1 [return_type]
//method_name2 [return_type]
//method_name3 [return_type]
//...
//method_namen [return_type]
//}
//
///* 定义结构体 */
//type struct_name struct {
//	/* variables */
//}
//
///* 实现接口方法 */
//func (struct_name_variable struct_name) method_name1() [return_type] {
//	/* 方法实现 */
//}
//...
//func (struct_name_variable struct_name) method_namen() [return_type] {
//	/* 方法实现*/
//}

import "fmt"

type Phone interface {
	Call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) Call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) Call() {
	fmt.Println("I am iPhone, I can call you!")
}

type Int struct {
}

func (int Int) Call() {
	x := 1
	for x < 10 {
		x = x + 1
	}
	fmt.Println(x)
}
