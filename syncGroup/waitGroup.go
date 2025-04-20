package syncGroup

import "sync"

type WaitGroup struct {
	wg sync.WaitGroup
}

func (w WaitGroup) Go(fn func()) {
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		fn()
	}()
}
