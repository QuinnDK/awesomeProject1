package text

import "fmt"

type PPPhone interface {
	Call3(param int) string
	Takephoto()
}

type Huawei struct {
}

func (huawei Huawei) Call3(param int) string {
	fmt.Println("i am Huawei, i can call you!", param)
	return "damon"
}

func (huawei Huawei) Takephoto() {
	fmt.Println("i can take a photo for you")
}
