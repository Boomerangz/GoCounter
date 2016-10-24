package main

import (
	"fmt"
	"time"
)

type TaskExecutor struct {
	dataChannel  chan string

	currentlyWorking AtomicCounter
	task             func(url string) error
	maximumWorking   int
}

func InitTaskExecutor(workersMaximumCount int, channelSize int, task func(data string) error) *TaskExecutor {
	return &TaskExecutor{
		dataChannel: make(chan string, channelSize),
		task:task,
		maximumWorking: workersMaximumCount, // <= 0 means unlimited count of goroutines
	}
}

func (t *TaskExecutor) AppendData(data string) {
	t.dataChannel <- data
	if t.maximumWorking <= 0 || t.currentlyWorking.GetCurrentCount() < t.maximumWorking {
		n := t.currentlyWorking.IncAndGet()
		go t.runExecutor(n)
	}
}

func (t *TaskExecutor) Close() {
	for {
		if t.currentlyWorking.GetCurrentCount() == 0 {
			break
		} else {
			time.Sleep(50)
		}
	}
}

func (t *TaskExecutor) runExecutor(n int) {
	for {
		select {
		case data := <-t.dataChannel:
			{
				err := t.task(data)
				if err != nil {
					fmt.Printf("Error: executing task failed with data %s and error %v\n", data, err.Error())
				}
			}
		default:
			{
				t.currentlyWorking.Dec()
				return
			}
		}
	}
}
