package scheduler

import "crawler/concurrent/engine"

//所有worker公用一个输入
type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.workerChan <- r }() //避免worker的in,out之间循环等待
}
