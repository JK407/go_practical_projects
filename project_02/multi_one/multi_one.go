package multi_one

import (
	"project_02/out"
	"sync"
)

// Task
// @Description: Task
// @Author Oberl-Fitzgerald 2024-07-10 15:24:59
type Task struct {
	ID int64
}

// run
// @Description 执行任务
// @Author Oberl-Fitzgerald 2024-07-10 15:25:01
func (t *Task) run() {
	out.Println(t.ID)
}

// taskChan，缓冲区大小为10
var taskChan = make(chan *Task, 10)

// taskNum，任务数量
const taskNum int64 = 10000

// nums，每次生成任务数量
const nums int64 = 100

// producer
// @Description 生产者
// @Author Oberl-Fitzgerald 2024-07-10 15:25:16
// @Param  wo chan<- *Task
// @Param  startNum int64
// @Param  nums int64
func producer(wo chan<- *Task, startNum int64, nums int64) {
	var i int64
	for i = startNum; i < startNum+nums; i++ {
		wo <- &Task{ID: i}
	}
}

// consumer
// @Description 消费者
// @Author Oberl-Fitzgerald 2024-07-10 15:26:01
// @Param  ro <-chan *Task
func consumer(ro <-chan *Task) {
	for t := range ro {
		if t.ID != 0 {
			t.run()
		}
	}
}

func Exec() {
	//  创建WaitGroup
	wg := &sync.WaitGroup{}
	//  创建一个关闭channel的WaitGroup
	pwg := &sync.WaitGroup{}
	var i int64
	for i = 0; i < taskNum; i += nums {
		if i >= taskNum {
			break
		}
		wg.Add(1)
		pwg.Add(1)
		go func(i int64) {
			defer wg.Done()
			defer pwg.Done()
			producer(taskChan, i, nums)
		}(i)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		consumer(taskChan)
	}()
	pwg.Wait()
	//  创建一个关闭channel的goroutine
	go close(taskChan)
	wg.Wait()
	out.Println("task done")
}
