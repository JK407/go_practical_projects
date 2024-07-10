package multi_multi

import (
	"fmt"
	"project_02/out"
	"time"
)

// Task
// @Description: Task
// @Author Oberl-Fitzgerald 2024-07-10 15:26:54
type Task struct {
	ID int64
}

// run
// @Description 执行任务
// @Author Oberl-Fitzgerald 2024-07-10 15:26:56
func (t *Task) run() {
	out.Println(t.ID)
}

// taskChan，缓冲区大小为10
var taskChan = make(chan *Task, 10)

// 关闭channel的channel
var done = make(chan struct{})

// taskNum，任务数量
const taskNum int64 = 10000

// producer
// @Description 生产者
// @Author Oberl-Fitzgerald 2024-07-10 15:27:24
// @Param  wo chan<- *Task
// @Param  done chan struct{}
func producer(wo chan<- *Task, done chan struct{}) {
	var i int64
	//  死循环
	for {
		if i >= taskNum {
			i = 0
		}
		i++
		select {
		//  将任务写入channel
		case wo <- &Task{ID: i}:
		//  如果done里面有数据，则退出
		case <-done:
			out.Println("producer done")
			return
		}
	}
}

// consumer
// @Description 消费者
// @Author Oberl-Fitzgerald 2024-07-10 15:28:06
// @Param  ro <-chan *Task
// @Param  done chan struct{}
func consumer(ro <-chan *Task, done chan struct{}) {
	for {
		select {
		//  从channel中读取数据，如果有数据则执行任务
		case t := <-ro:
			if t.ID != 0 {
				t.run()
			}
		//  如果done里面有数据，则退出
		case <-done:
			//  关闭后，因为channel是有缓存的，所以会继续执行完缓存中的数据
			for t := range ro {
				if t.ID != 0 {
					t.run()
				}
			}
			out.Println("consumer done")
			return
		}
	}
}

func Exec() {
	go producer(taskChan, done)
	go producer(taskChan, done)
	go producer(taskChan, done)
	go producer(taskChan, done)
	go producer(taskChan, done)
	go producer(taskChan, done)

	go consumer(taskChan, done)
	go consumer(taskChan, done)
	go consumer(taskChan, done)

	time.Sleep(5 * time.Second)
	//  先关闭done
	close(done)
	//  再关闭taskChan
	close(taskChan)

	time.Sleep(5 * time.Second)
	fmt.Println(len(taskChan))
	out.Println("task done")
}
