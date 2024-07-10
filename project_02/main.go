package main

import (
	"os"
	"os/signal"
	"project_02/multi_multi"
	"project_02/multi_one"
	"project_02/one_multi"
	"project_02/one_one"
	"project_02/out"
	"syscall"
)

func main() {
	//  创建一个异步输入输出
	o := out.NewOut()
	go o.OutPut()
	//  输入数据
	out.Println("hello world")
	out.Println("hello world")
	out.Println("hello world")

	//  一个生产者，一个消费者
	one_one.Exec()
	//  一个生产者，多个消费者
	one_multi.Exec()
	//  多个生产者，一个消费者
	multi_one.Exec()
	//  多个生产者，多个消费者
	multi_multi.Exec()

	//  监听信号
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
