package syncGroup

import (
	"sync"
)

/*
func example() {
	wg := NewPool(20)
	startTime := time.Now()
	for _, num := range serialNums {
		mv, err := l.svcCtx.Model.FindOneBySerialNum(num)
		if err != nil {
			continue
		}
		if mv == nil {
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			// handler logic
		}()
	}

	wg.Wait()
	endTime := time.Now()
	useTime := endTime.Sub(startTime).Seconds()
	fmt.Printf("\n\n\n\n\n total use time ====================>>>>>  %v \n\n\n\n", useTime)
}
*/

// SimplePool
type SimplePool struct {
	workChan chan int
	wg       sync.WaitGroup
}

// NewPool 生成一个工作池, coreNum 限制
func NewPool(coreNum int) *SimplePool {
	ch := make(chan int, coreNum)
	return &SimplePool{
		workChan: ch,
		wg:       sync.WaitGroup{},
	}
}

// Add 添加
func (ap *SimplePool) Add(num int) {
	for i := 0; i < num; i++ {
		ap.workChan <- i
		ap.wg.Add(1)
	}
}

// Done 完结
func (ap *SimplePool) Done() {
LOOP:
	for {
		select {
		case <-ap.workChan:
			break LOOP
		}
	}
	ap.wg.Done()
}

// Wait 等待
func (ap *SimplePool) Wait() {
	ap.wg.Wait()
}
