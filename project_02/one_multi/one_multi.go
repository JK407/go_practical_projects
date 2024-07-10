package one_multi

import (
	"project_02/out"
	"sync"
)

// Task
// @Description: Task
// @Author Oberl-Fitzgerald 2024-07-10 15:23:21
type Task struct {
	ID int64
}

// run
// @Description 执行任务
// @Author Oberl-Fitzgerald 2024-07-10 15:23:24
func (t *Task) run() {
	out.Println(t.ID)
}

// taskChan，缓冲区大小为10
var taskChan = make(chan *Task, 10)

// taskNum，任务数量
const taskNum int64 = 10000

// producer
// @Description 生产者
// @Author Oberl-Fitzgerald 2024-07-10 15:23:41
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
// @Author Oberl-Fitzgerald 2024-07-10 15:23:53
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
// @Author Oberl-Fitzgerald 2024-07-10 15:24:02
func Exec() {
	//  创建WaitGroup
	wg := &sync.WaitGroup{}
	//  添加一个任务
	wg.Add(1)
	//  创建一个goroutine，执行producer
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		producer(taskChan)
	}(wg)
	//  创建多个goroutine，执行consumer
	for i := int64(0); i < taskNum; i++ {
		if i%100 == 0 {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				consumer(taskChan)
			}(wg)
		}
	}
	//  等待所有任务执行完毕
	wg.Wait()
	out.Println("task done")
}
