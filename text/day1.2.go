package text

import "fmt"

func Text()  {
	var s int
	fmt.Scan(&s)
	if s%2 == 0 {
		fmt.Printf(" %d +是偶数\n",s)
	}else {
		fmt.Printf("%d 不是偶数\n",s)
	}
	fmt.Println("输入的"+"数字是",s)
}
