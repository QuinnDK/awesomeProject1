package text

import "fmt"

const LTNES int = 10

//杨辉三角的实现
func Showyanghuisanjiao() {
	nums := []int{}

	for i := 0; i < LTNES; i++ {
		for j := 0; j < LTNES-i; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j < i+1; j++ {
			var length = len(nums)
			var value int
			if j == 0 || j == i {
				value = 1
			} else {
				value = nums[length-i] + nums[length-i-1]
			}
			nums = append(nums, value)

			fmt.Print(value, " ")
		}
		fmt.Println("")
	}
}
