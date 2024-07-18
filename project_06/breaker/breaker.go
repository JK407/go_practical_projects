package breaker

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	SATE_CLOSE      = iota //  SATE_CLOSE @Description: 熔断器关闭状态
	STATE_OPEN             //  STATE_OPEN @Description: 熔断器打开状态
	STATE_HALF_OPEN        //  STATE_HALF_OPEN @Description: 熔断器半开状态
)

// Breaker
// @Description: 熔断器
// @Author Oberl-Fitzgerald 2024-07-18 15:42:15
type Breaker struct {
	mu                sync.Mutex    // mu @Description: 互斥锁
	state             int           // state @Description: 熔断器状态
	failureThreshold  int           // failureThreshold @Description: 连续失败阈值
	successThreshold  int           // successThreshold @Description: 连续成功阈值
	halfMaxRequests   int           // halfMaxRequests @Description: 半开半闭状态最大请求数
	halfCycleReqCount int           // halfCycleReqCount @Description: 半开半闭状态周期内请求数
	timeout           time.Duration // timeout @Description: 周期时间，断开状态下的超时时间
	failureCount      int           // failureCount @Description: 失败计数
	successCount      int           // successCount @Description: 成功计数
	cycleStartTime    time.Time     // cycleStartTime @Description: 周期开始时间
}

// NewBreaker
// @Description 初始化熔断器
// @Author Oberl-Fitzgerald 2024-07-18 15:42:20
// @Param  failureThreshold int
// @Param  successThreshold int
// @Param  halfMaxRequests int
// @Param  timeout time.Duration
// @Return *Breaker
func NewBreaker(failureThreshold, successThreshold, halfMaxRequests int, timeout time.Duration) *Breaker {
	return &Breaker{
		state:            SATE_CLOSE, //  SATE_CLOSE @Description: 熔断器关闭状态
		failureThreshold: failureThreshold,
		successThreshold: successThreshold,
		halfMaxRequests:  halfMaxRequests,
		timeout:          timeout,
	}
}

// Exec
// @Description 执行函数
// @Author Oberl-Fitzgerald 2024-07-18 15:42:28
// @Param  f func() error
// @Return error
func (b *Breaker) Exec(f func() error) error {
	b.before()
	fmt.Printf("%+v\n", b)
	//  如果熔断器打开，则返回错误
	if b.state == STATE_OPEN {
		return errors.New("breaker is open")
	}
	//  如果熔断器关闭，则执行函数
	if b.state == SATE_CLOSE {
		err := f()
		b.after(err)
		return err
	}
	//  如果熔断器半开，则判断是否达到最大请求数
	if b.state == STATE_HALF_OPEN {
		//  如果未达到最大请求数，则执行函数；否则返回错误
		if b.halfCycleReqCount < b.halfMaxRequests {
			err := f()
			b.after(err)
			return err
		} else {
			return errors.New("breaker is half open")
		}
	}
	return nil
}

// before
// @Description 执行函数前的操作
// @Author Oberl-Fitzgerald 2024-07-18 15:43:49
func (b *Breaker) before() {
	b.mu.Lock()
	defer b.mu.Unlock()
	//  根据熔断器状态进行操作
	switch b.state {
	//  如果熔断器打开，则判断是否超时，如果超时则进入半开状态
	case STATE_OPEN:
		if b.cycleStartTime.Add(b.timeout).Before(time.Now()) {
			//  进入半开状态
			b.state = STATE_HALF_OPEN
			//  重置计数
			b.reset()
			return
		}
	//  如果熔断器半开，则判断是否达到成功阈值，如果达到则进入关闭状态
	case STATE_HALF_OPEN:
		if b.successCount >= b.successThreshold {
			b.state = SATE_CLOSE
			b.reset()
			return
		}
		//  如果超时，则重置计数
		if b.cycleStartTime.Add(b.timeout).Before(time.Now()) {
			b.cycleStartTime = time.Now()
			b.halfCycleReqCount = 0
			return
		}
	//  如果熔断器关闭，则判断是否超时，如果超时则重置计数
	case SATE_CLOSE:
		if b.cycleStartTime.Add(b.timeout).Before(time.Now()) {
			b.reset()
			return
		}
	}
}

// after
// @Description 执行函数后的操作
// @Author Oberl-Fitzgerald 2024-07-18 15:49:39
// @Param  err error
func (b *Breaker) after(err error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	//  根据错误情况进行操作
	if err == nil {
		//  如果成功，则执行成功操作
		b.onSuccess()
	} else {
		//  如果失败，则执行失败操作
		b.onFailure()

	}
}

// reset
// @Description 重置计数
// @Author Oberl-Fitzgerald 2024-07-18 15:49:56
func (b *Breaker) reset() {
	b.failureCount = 0
	b.successCount = 0
	b.halfCycleReqCount = 0
	b.cycleStartTime = time.Now()
}

// onSuccess
// @Description 成功操作
// @Author Oberl-Fitzgerald 2024-07-18 15:50:02
func (b *Breaker) onSuccess() {
	//  重置失败计数
	b.failureCount = 0
	//  如果熔断器处于半开状态，则增加成功计数和半开状态请求数
	if b.state == STATE_HALF_OPEN {
		b.successCount++
		b.halfCycleReqCount++
		//  如果成功计数达到成功阈值，则进入关闭状态
		if b.successCount >= b.successThreshold {
			b.state = SATE_CLOSE
			b.reset()
		}
	}
}

// onFailure
// @Description 失败操作
// @Author Oberl-Fitzgerald 2024-07-18 15:50:36
func (b *Breaker) onFailure() {
	b.successCount = 0
	b.failureCount++
	//  如果熔断器处于半开状态或者关闭状态且失败计数达到失败阈值，则进入打开状态
	if b.state == STATE_HALF_OPEN || (b.state == SATE_CLOSE && b.failureCount >= b.failureThreshold) {
		b.state = STATE_OPEN
		b.reset()
		return
	}
}
