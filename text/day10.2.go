package text

//求解平方根
func sqrt(x float64, i float64) (float64, float64) {
	remain := (i*i - x) / (2 * i)
	i = i - remain
	if remain > 0 {
		return sqrt(x, i)
	} else {
		return i, remain
	}
}
func Get_sqrt(x float64) float64 {
	i, _ := sqrt(x, x)
	return i
}
