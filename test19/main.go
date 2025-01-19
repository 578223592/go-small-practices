package main

import (
	"fmt"
	"sync"
	"time"
)

// TimeWindow 代表每个小窗口的结构体
type TimeWindow struct {
	count int       // 记录请求数量
	start time.Time // 记录窗口的开始时间
}

// LeapArray 是一个具有时间窗口的数组
type LeapArray struct {
	windows  []TimeWindow // 循环数组，每个元素是一个时间窗口
	size     int64        // 数组的大小（窗口数量）
	interval int64        // 每个窗口的时长
	mu       sync.Mutex   // 保证并发安全
}

// NewLeapArray 初始化一个 Leap Array
func NewLeapArray(size int64, intervalMs int64) *LeapArray {
	if size <= 0 {
		panic("size<=0")
	}
	windows := make([]TimeWindow, size)
	return &LeapArray{
		windows:  windows,
		size:     size,
		interval: intervalMs,
	}
}

// getCurrentWindow 获取当前时间对应的窗口
/**
下标 = (当前时间 % 时间跨度) / 一个时间窗口的时间,而 时间跨度= 一个时间窗口的时间 * 窗口数量（整个array能代表的时间长度）
*/
func (la *LeapArray) getCurrentWindow() *TimeWindow {
	la.mu.Lock()
	defer la.mu.Unlock()

	now := time.Now()
	index := (now.UnixMilli() % (la.size * la.interval)) / la.interval // 计算当前时间对应的窗口索引

	curWindow := &la.windows[index]
	if now.Sub(curWindow.start).Milliseconds() >= la.interval {
		// 如果窗口已过期，重置它
		curWindow.count = 0
		curWindow.start = now.Add(-1 * time.Duration(now.UnixMilli()%la.interval) * time.Millisecond) //计算当前窗口的开始时间
	}
	return curWindow
}

// AddRequest 向当前窗口添加一个请求计数
func (la *LeapArray) AddRequest() {
	window := la.getCurrentWindow()
	la.mu.Lock()
	window.count++
	la.mu.Unlock()
}

// GetRequestCount 获取最近一分钟内的请求总数
func (la *LeapArray) GetRequestCount() int {
	la.mu.Lock()
	defer la.mu.Unlock()

	total := 0
	now := time.Now()
	for _, window := range la.windows {
		// 仅统计还在有效期内的窗口
		if now.Sub(window.start).Milliseconds() <= la.size*la.interval {
			total += window.count
		}
	}
	return total
}

func main() {
	leapArray := NewLeapArray(6, 30) // 窗口个数，每个窗口长度（Ms）

	// 模拟请求
	for i := 0; i < 200; i++ {
		//log.Println(i)
		leapArray.AddRequest()
		//精确暂停1ms
		time.Sleep(1 * time.Millisecond) //1ms一个请求
	}
	count := leapArray.GetRequestCount() //既然是将时间划分成了窗口，统计结果会波动一个窗口
	// 虽然说是波动，实际上基本都是少统计。具体少统计的原因是：
	// 当前时间会落在某一个窗口中间，那么这个窗口还没经历的部分就没办法统计进去了。而上一个周期经历的会因为leap array的特性被清空掉呢
	//或者说leap array的粒度就是其中bucket的粒度，所以能统计到的时间片段的精度是有限的。

	fmt.Printf("最近周期内的请求数: %d\n", count) //周期有效值为：窗口数 * 窗口长度
}

//2024-11-11 10:05:35 正在研究滑动窗口：跳跃数组 leap array
// 2024年11月14日10:42:22 待研究sentinel源码中的leap array实现
