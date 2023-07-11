package semaphore

import (
	"fmt"

	"github.com/perfectgentlemande/go-semaphore-vs-worker-pool/users"
)

type Semaphore struct {
	C chan struct{}
}

func (s *Semaphore) Acquire() {
	s.C <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.C
}

type resultWithError struct {
	User users.User
	Err  error
}

func DeactivateUsers(usrs []users.User, gCount int) ([]users.User, error) {
	sem := Semaphore{
		C: make(chan struct{}, gCount),
	}

	outputCh := make(chan resultWithError, len(usrs))
	sgnlCh := make(chan struct{})

	output := make([]users.User, 0, len(usrs))

	for _, v := range usrs {
		go func(usr users.User) {
			sem.Acquire()
			defer sem.Release()

			err := usr.Deactivate()

			select {
			case outputCh <- resultWithError{
				User: usr,
				Err:  err,
			}:
			case <-sgnlCh:
				return
			}
		}(v)
	}

	for i := len(usrs); i > 0; i-- {
		resErr := <-outputCh
		if resErr.Err != nil {
			close(sgnlCh)
			return nil, fmt.Errorf("an error occurred: %w", resErr.Err)
		}

		output = append(output, resErr.User)
	}

	return output, nil
}
