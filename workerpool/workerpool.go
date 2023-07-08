package workerpool

import (
	"sync"

	"github.com/perfectgentlemande/go-semaphore-vs-worker-pool/users"
)

type resultWithError struct {
	User users.User
	Err  error
}

func deactivateUser(wg *sync.WaitGroup, inCh <-chan users.User, outCh chan<- resultWithError) {
	defer wg.Done()

	for usr := range inCh {
		err := usr.Deactivate()
		outCh <- resultWithError{
			User: usr,
			Err:  err,
		}
	}
}

func DeactivateUsers(usrs []users.User, wgCount int) ([]users.User, error) {
	inCh := make(chan users.User)
	outCh := make(chan resultWithError)
	wg := &sync.WaitGroup{}

	output := make([]users.User, 0, len(usrs))

	go func() {
		defer close(inCh)

		for i := range usrs {
			inCh <- usrs[i]
		}
	}()

	go func() {
		for i := 0; i < wgCount; i++ {
			wg.Add(1)

			go deactivateUser(wg, inCh, outCh)
		}
		wg.Wait()
		close(outCh)
	}()

	for res := range outCh {
		if res.Err != nil {
			return nil, res.Err
		}

		output = append(output, res.User)
	}

	return output, nil
}
