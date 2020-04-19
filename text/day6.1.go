package text

/*
go语言结构体


感觉和C语言没啥差别，对于想通过传参来改变结构体内的参数值，或者其他，需要通过传递地址
*/
type Books struct {
	Title   string
	Auther  string
	subject string
	Book_id int
}

func Changebook(book Books) {
	book.Title = "change book"

}
func Changebook1(book *Books) {
	book.Title = "change book1"
}
