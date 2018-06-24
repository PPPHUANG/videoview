package taskrunner

import "time"

//你要这样拿 这样 我说你这不行啊
type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker:time.NewTicker(interval *time.Second),
		runner:r,
	}
}

func (w *Worker) startWorker() {
	for {
		select {
		case <- w.ticker.C:
			go w.runner.StartAll()
		}
	}
}

func Start() {
	r := NewRunner(3,true,VideoClearDispatcher,VideoClearExecutor)
	w := NewWorker(3, r)
	go w.startWorker()
}