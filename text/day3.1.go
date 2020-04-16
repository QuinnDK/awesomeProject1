package text

type Circle struct {
	Radius float64
}

func (c Circle) GetArea() float64 {
	return 3.14 * c.Radius * c.Radius
}

/*
Go 没有面向对象，而我们知道常见的 Java。
C++ 等语言中，实现类的方法做法都是编译器隐式的给函数加一个 this 指针
而在 Go 里，这个 this 指针需要明确的申明出来，其实和其它 OO 语言并没有很大的区别。
*/
