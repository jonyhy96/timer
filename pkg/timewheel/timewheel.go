package timewheel

import (
	"fmt"
	"time"

	"github.com/jonyhy96/timer/pkg/queue"
	"github.com/jonyhy96/timer/pkg/util"

	"github.com/golang/glog"
)

// Handler handler tasks
type Handler func([]string)

// TimeWheel TimeWheel
type TimeWheel struct {
	Size int
	Q    []queue.Queue
}

// NewTimeWheel creates a new TimeWheel
func NewTimeWheel(size int) *TimeWheel {
	return &TimeWheel{
		Size: size,
		Q:    make([]queue.Queue, size),
	}
}

func (t *TimeWheel) run(handler Handler) {
	for {
		time.Sleep(1 * time.Second)
		unixNow := time.Now().Unix()
		if tasks := t.GetTask(unixNow); len(tasks) > 0 {
			go handler(tasks)
		} else {
			glog.V(4).Infof("task not found %d\n", unixNow)
		}
	}
}

// AddTask add task into queue
func (t *TimeWheel) AddTask(taskID string, seconds int64) error {
	if expired := util.CheckDate(seconds); expired {
		return fmt.Errorf("task expired")
	}
	qid := util.GetPosition(seconds, t.Size)
	t.Q[qid].AddTask(taskID, seconds)
	return nil
}

// GetTask get task from queue
func (t *TimeWheel) GetTask(seconds int64) []string {
	qid := util.GetPosition(seconds, t.Size)
	return t.Q[qid].GetTask(seconds)
}

// GetQueue gets the queue
func (t *TimeWheel) GetQueue(seconds int64) *queue.Queue {
	qid := util.GetPosition(seconds, t.Size)
	return &t.Q[qid]
}

// Run run
func (t *TimeWheel) Run(handler Handler) {
	t.run(handler)
}
