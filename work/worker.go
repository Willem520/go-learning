package work

import "sync"

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

/**
创建新工作池
*/
func New(maxGoroutines int) *Pool {
	p := &Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return p
}

/**
提交工作到工作池
*/
func (p *Pool) Run(w Worker) {
	p.work <- w
}

/**
等待所有goroutine停止工作
*/
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
