package main

import "sync"

type AtomicCounter struct {
	counter int
	mutex   sync.Mutex
}

func (ac *AtomicCounter) Inc() {
	ac.mutex.Lock()
	defer ac.mutex.Unlock()

	ac.counter += 1
}

func (ac *AtomicCounter) IncAndGet() int {
	ac.mutex.Lock()
	defer ac.mutex.Unlock()

	ac.counter += 1
	return ac.counter
}

func (ac *AtomicCounter) DecAndGet() int {
	ac.mutex.Lock()
	defer ac.mutex.Unlock()

	ac.counter -= 1
	return ac.counter
}

func (ac *AtomicCounter) Dec() {
	ac.mutex.Lock()
	defer ac.mutex.Unlock()

	ac.counter -= 1
}

func (ac *AtomicCounter) GetCurrentCount() int {
	ac.mutex.Lock()
	defer ac.mutex.Unlock()

	return ac.counter
}
