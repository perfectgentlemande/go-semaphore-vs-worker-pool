package workerpool

import (
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/perfectgentlemande/go-semaphore-vs-worker-pool/users"
)

func generateUsers(count int) (activated, deactivated []users.User) {
	activated = make([]users.User, count)
	deactivated = make([]users.User, count)

	for i := 0; i < count; i++ {
		activated[i] = users.User{
			ID:       uuid.New(),
			Fullname: "Fullname " + strconv.Itoa(i),
			Username: "Username " + strconv.Itoa(i),
			Active:   true,
		}
		deactivated[i] = users.User{
			ID:       activated[i].ID,
			Fullname: activated[i].Fullname,
			Username: activated[i].Username,
			Active:   false,
		}
	}

	return res
}

func TestDeactivateUsers(t *testing.T) {
	activated, deactivated := generateUsers(10000)

	type TestCase struct {
		Activated   []users.User
		Deactivated []users.User
	}

	cases := []TestCase{
		{
			Activated:   activated[:1],
			Deactivated: deactivated[:1],
		},
		{
			Activated:   deactivated[:10],
			Deactivated: deactivated[:10],
		},
		{
			Activated:   deactivated[:100],
			Deactivated: deactivated[:100],
		},
		{
			Activated:   deactivated[:1000],
			Deactivated: deactivated[:1000],
		},
		{
			Activated:   deactivated[:10000],
			Deactivated: deactivated[:10000],
		},
	}

	for i := range cases {

	}
}
