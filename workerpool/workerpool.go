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
	inputCh := make(chan users.User)
	outputCh := make(chan resultWithError)
	wg := &sync.WaitGroup{}

	output := make([]users.User, 0, len(usrs))

	go func() {
		defer close(inputCh)

		for i := range usrs {
			inputCh <- usrs[i]
		}
	}()

	go func() {
		for i := 0; i < wgCount; i++ {
			wg.Add(1)

			go deactivateUser(wg, inputCh, outputCh)
		}
		wg.Wait()
		close(outputCh)
	}()

	for res := range outputCh {
		if res.Err != nil {
			return nil, res.Err
		}

		output = append(output, res.User)
	}

	return output, nil
}
