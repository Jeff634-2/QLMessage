package scheduler

import "QLProject/engine"

/*
	request channel、work channel
*/

//type QueuedScheduler
type SimpleScheduler struct {
	workChan chan engine.Request
}

func (s *SimpleScheduler) ConfigerMasterWorkerChan(c chan engine.Request) {
	s.workChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.workChan <- r }()
}
