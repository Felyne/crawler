package scheduler

import "crawler/concurrent/engine"

type SimpleScheduler struct {
	requestChan chan engine.Request
}

// 所有worker拿到同一个输入chan
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.requestChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.requestChan <- r }() //避免worker的in,out之间循环等待
}

func (s *SimpleScheduler) Run() {
	s.requestChan = make(chan engine.Request)
}
