package scheduler

import "imooc.com/sb/learngo/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(req engine.Request) {
	// send request down to worker chan
	go func() { s.workerChan <- req }()
}

func (s *SimpleScheduler) ConfigMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}
