package text

import "fmt"

/*
今日学习党课的一天，比较担心理论测试，但感觉又没什么，毕竟是考研过来的
对于go语言的学习，即使时间再紧张，多少也要坚持写代码，学完基本语法，还有许多事要做，任重道远
*/

//利用数组实现杨辉三角
func GetYangHuiTriangleNextLine(inArr []int) []int {
	var out []int
	var i int
	arrLen := len(inArr)
	out = append(out, 1)
	if 0 == arrLen {
		return out
	}
	for i = 0; i < arrLen-1; i++ {
		out = append(out, inArr[i]+inArr[i+1])
	}
	out = append(out, 1)
	return out
}

func Triangle(n int) {
	var item []int
	for i := 1; i <= n; i++ {
		item_len := len(item)
		if item_len == 0 {
			item = append(item, 1)
		} else {
			temp_s := []int{1}
			for j := 0; j < item_len-1; j++ {
				temp_s = append(temp_s, item[j]+item[j+1])
			}
			temp_s = append(temp_s, 1)
			item = temp_s
		}
		fmt.Println(item)
	}
}
