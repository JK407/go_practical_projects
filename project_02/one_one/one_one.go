package one_one

import (
	"project_02/out"
	"sync"
)

// Task
// @Description: Task
// @Author Oberl-Fitzgerald 2024-07-10 15:19:59
type Task struct {
	ID int64
}

// run
// @Description 执行任务
// @Author Oberl-Fitzgerald 2024-07-10 15:20:03
func (t *Task) run() {
	out.Println(t.ID)
}

// taskChan，缓冲区大小为10
var taskChan = make(chan *Task, 10)

// taskNum，任务数量
const taskNum int64 = 10000

// producer
// @Description 生产者
// @Author Oberl-Fitzgerald 2024-07-10 15:21:19
// @Param  wo chan<- *Task
func producer(wo chan<- *Task) {
	//  生成任务
	for i := int64(1); i <= taskNum; i++ {
		wo <- &Task{ID: i}
	}
	//  关闭channel
	close(wo)
}

// consumer
// @Description 消费者
// @Author Oberl-Fitzgerald 2024-07-10 15:21:35
// @Param  ro <-chan *Task
func consumer(ro <-chan *Task) {
	//  遍历channel
	for t := range ro {
		if t.ID != 0 {
			t.run()
		}
	}
}

// Exec
// @Description 执行任务
// @Author Oberl-Fitzgerald 2024-07-10 15:22:03
func Exec() {
	//  创建WaitGroup
	wg := &sync.WaitGroup{}
	//  添加两个任务
	wg.Add(2)
	//  创建一个goroutine，执行producer
	go func(wg *sync.WaitGroup) {
		//  函数执行完毕，WaitGroup计数器减1
		defer wg.Done()
		//  执行producer
		producer(taskChan)
	}(wg)
	//  创建一个goroutine，执行consumer
	go func(wg *sync.WaitGroup) {
		//  函数执行完毕，WaitGroup计数器减1
		defer wg.Done()
		//  执行consumer
		consumer(taskChan)
	}(wg)
	//  等待所有任务执行完毕
	wg.Wait()
	//  输出任务完成
	out.Println("task done")
}
