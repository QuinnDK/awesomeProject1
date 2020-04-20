package text

import "fmt"

//定义切片
/*
声明一个未指定大小的数组来定义切片：var identifier []type   切片不需要说明长度。
使用make()函数来创建切片: var slice1 []type = make([]type, len)
也可以简写为slice1 := make([]type, len)


切片是可索引的，并且可以由 len() 方法获取长度。
切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少
*/

func PrintSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
	if x == nil {
		fmt.Printf("切片是空的")
	}
}
