package text

import "fmt"

//画个菱形
const X int = 9
const Y int = 9

func Lingxing() {
	row := 1
	for row <= X {
		//计算星星数
		count := 0
		if row <= X/2+1 {
			count = 2*row - 1
		} else {
			count = 2*(Y-row) + 1
		}
		row++
		text := " "

		//填充的星星范围
		start_min := ((X - count) / 2) + 1
		start_max := ((X - count) / 2) + count
		for indedx := 1; indedx <= X; indedx++ {
			if indedx >= start_min && indedx <= start_max {
				text += "*"
			} else {
				text += " "
			}
		}
		fmt.Println(text)
	}
}
