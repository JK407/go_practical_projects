package out

import "fmt"

// Out
// @Description: 输出结构体
// @Author Oberl-Fitzgerald 2024-07-10 15:18:24
type Out struct {
	// data @Description: channel数据
	data chan interface{}
}

// out @Description: 全局变量
var out *Out

// NewOut
// @Description 创建Out实例
// @Author Oberl-Fitzgerald 2024-07-10 15:18:49
// @Return *Out
func NewOut() *Out {
	if out == nil {
		out = &Out{
			//  初始化channel，有缓冲区
			data: make(chan interface{}, 65535),
		}
	}
	return out
}

// Println
// @Description 将i写入channel
// @Author Oberl-Fitzgerald 2024-07-10 15:17:54
// @Param  i interface{}
func Println(i interface{}) {
	out.data <- i
}

// OutPut
// @Description 输出数据
// @Author Oberl-Fitzgerald 2024-07-10 15:19:25
func (o *Out) OutPut() {
	for {
		select {
		case i := <-o.data:
			fmt.Println(i)
		}
	}
}
