package queue

import (
	"fmt"
	"sync"
)

// Queue Queue
type Queue struct {
	tasks sync.Map
}

// AddTask add task to queue
func (q *Queue) AddTask(task string, seconds int64) {
	if val, ok := q.tasks.Load(seconds); ok {
		val = append(val.([]string), task)
	} else {
		q.tasks.Store(seconds, []string{task})
	}
}

// GetTask get task from queue
func (q *Queue) GetTask(seconds int64) []string {
	val, ok := q.tasks.Load(seconds)
	if ok {
		q.tasks.Delete(seconds)
		return val.([]string)
	}
	return nil
}

// DeleteTask deletes task from queue
func (q *Queue) DeleteTask(task string, seconds int64) error {
	val, ok := q.tasks.Load(seconds)
	if !ok {
		return fmt.Errorf("no task found base on %d", seconds)
	}
	for index, task := range val.([]string) {
		if task == val.([]string)[index] {
			val = append(val.([]string)[:index], val.([]string)[index+1:]...)
		}
	}
	q.tasks.Store(seconds, val)
	return nil
}
