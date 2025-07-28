package util

import (
	"fmt"
	"sync"
)

// Uniquify is a type of advanced mutex. It allows to create named resource locks.
type Uniquify interface {
	// Call executes only one callable with same id at a time.
	// Multilpe asynchronous calls with same id will be executed sequentally.
	Call(id string, callable func() error) error
}

// NewUniquify returns a new thread-safe uniquify object.
func NewUniquify() Uniquify {
	return &uniquify{
		tasks: make(map[string]*sync.WaitGroup),
	}
}

type uniquify struct {
	lock  sync.Mutex
	tasks map[string]*sync.WaitGroup
}

func (u *uniquify) Call(id string, callable func() error) error {
	errC := make(chan error)

	// Check for existing task
	u.lock.Lock()
	oldWg := u.tasks[id]
	if oldWg != nil {
		u.lock.Unlock()

		// Wait for existing task and retry
		go func() {
			oldWg.Wait()
			result := u.Call(id, callable)
			errC <- result
		}()

		return <-errC
	}

	// Set up new task
	wg := new(sync.WaitGroup)
	wg.Add(1)
	u.tasks[id] = wg
	u.lock.Unlock()

	// Execute task
	go u.executeCallable(id, callable, wg, errC)

	return <-errC
}

func (u *uniquify) executeCallable(id string, callable func() error, wg *sync.WaitGroup, errC chan error) {
	var err error

	// Clean up
	defer func() {
		errC <- err

		u.lock.Lock()
		defer u.lock.Unlock()
		delete(u.tasks, id)
	}()
	defer wg.Done()

	// Handle panics
	defer func(err *error) {
		if panicData := recover(); panicData != nil {
			if e, ok := panicData.(error); ok {
				*err = e
				return
			}
			*err = fmt.Errorf("%+v", panicData)
		}
	}(&err)

	// Execute the callable
	err = callable()
}
