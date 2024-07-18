package limiter

import (
	"sync"
	"time"
)

type Limiter struct {
	tb *TokenBucket
}

type TokenBucket struct {
	mu              sync.Mutex    // mu @Description: 互斥锁
	size            int           // size @Description: 桶的大小
	count           int           // count @Description: 当前桶内的令牌数量
	rateLimit       time.Duration // rateLimit @Description: 令牌产生速率
	lastRequestTime time.Time     // lastRequestTime @Description: 最后成功请求的时间
}

// NewLimiter
// @Description 初始化限流器
// @Author Oberl-Fitzgerald 2024-07-18 14:33:53
// @Param  r time.Duration
// @Param  size int
// @Return *Limiter
func NewLimiter(r time.Duration, size int) *Limiter {
	return &Limiter{
		tb: &TokenBucket{
			rateLimit: r,
			size:      size,
			count:     size,
		},
	}
}

func (l *Limiter) Allow() bool {
	l.tb.mu.Lock()
	defer l.tb.mu.Unlock()

	return l.tb.allow()
}

func (tb *TokenBucket) allow() bool {
	//  填冲令牌
	tb.fill()
	if tb.count > 0 {
		tb.count--
		// 更新最后请求时间
		tb.lastRequestTime = time.Now()
		return true
	}
	return false
}

func (tb *TokenBucket) fill() {
	tb.count += tb.getFillTokenCount()
}

func (tb *TokenBucket) getFillTokenCount() int {
	if tb.count >= tb.size {
		return 0
	}
	if !tb.lastRequestTime.IsZero() {
		interval := time.Now().Sub(tb.lastRequestTime)
		// 计算产生的令牌数量
		generatedTokens := int(interval / tb.rateLimit)
		if tb.size-tb.count < generatedTokens {
			return tb.size - tb.count
		} else {
			return generatedTokens
		}
	}
	return 0
}
