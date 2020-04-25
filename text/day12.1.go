package text

// 自定义错误信息结构
type DIV_ERR struct {
	etype int // 错误类型
	v1    int // 记录下出错时的除数、被除数
	v2    int
}

// 实现接口方法 error.Error()
func (div_err DIV_ERR) Error() string {
	if 0 == div_err.etype {
		return "除零错误"
	} else {
		return "其他未知错误"
	}
}

// 除法
func Div(a int, b int) (int, string) {
	if b == 0 {
		// 返回错误信息
		return 0, string(0)
	} else {
		// 返回正确的商
		var nill = '0'
		return a / b, string(nill)
	}
}
