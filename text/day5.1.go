package text

//向函数传递数组

func GetAverage(arr [5]float32, size int) float32 {
	var i int
	var avg, sum float32

	for i = 0; i < size; i++ {
		sum = sum + arr[i]
	}

	avg = (sum) / float32(size)

	return avg
}
