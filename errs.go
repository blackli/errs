package errs

import (
	"fmt"
	"sync"
)

type errors struct {
	sync.RWMutex
	errors []error
}

func New() *errors {
	return &errors{}
}

func (e *errors) Append(errors ...error) {
	e.RWMutex.Lock()
	defer e.RWMutex.Unlock()
	e.errors = append(e.errors, errors...)
}

func (e *errors) String() string {
	return fmt.Sprintf("many errs: %v", e.errors)
}
