package text

/*
函数闭包：可以看作是一个函数中再嵌套一个函数
就比如：func a() func()int{}
*/

func Add(x1, x2 int) func() (int, int) {
	i := 0
	return func() (int, int) {
		i++
		return i, x1 + x2
	}

}

//如何做到不用在主文件中调用，就可以执行函数??

//第一个函数是用来传入参数，func 是使用参数 ，（int,int,int）是Add函数返回的参数类型
//说起类型，数组指针，指针数组，函数指针要加紧看了，快忘光了
func Add1(x1, x2 int) func(x3 int, x4 int) (int, int, int) {
	i := 0
	return func(x3 int, x4 int) (int, int, int) {
		i++
		return i, x1 + x2, x3 + x4
	}
}
