package semaphore

import (
	"fmt"
	"reflect"
	"sort"
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

	return activated, deactivated
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
			Activated:   activated[:10],
			Deactivated: deactivated[:10],
		},
		{
			Activated:   activated[:100],
			Deactivated: deactivated[:100],
		},
		{
			Activated:   activated[:1000],
			Deactivated: deactivated[:1000],
		},
		{
			Activated:   activated[:10000],
			Deactivated: deactivated[:10000],
		},
	}

	for i := range cases {
		got, _ := DeactivateUsers(cases[i].Activated, 10)
		sort.Slice(cases[i].Deactivated, func(j, k int) bool {
			return cases[i].Deactivated[j].Fullname < cases[i].Deactivated[k].Fullname
		})
		sort.Slice(got, func(j, k int) bool {
			return got[j].Fullname < got[k].Fullname
		})

		if !reflect.DeepEqual(got, cases[i].Deactivated) {
			t.Errorf("expected: %v got: %v", cases[i].Deactivated, got)
		}
	}
}

func BenchmarkDeactivateUsers(b *testing.B) {
	activated, deactivated := generateUsers(1000000)

	type TestCase struct {
		Activated   []users.User
		Deactivated []users.User
		GCount      int
	}

	cases := []TestCase{
		{
			Activated:   activated[:1],
			Deactivated: deactivated[:1],
			GCount:      10,
		},
		{
			Activated:   activated[:10],
			Deactivated: deactivated[:10],
			GCount:      10,
		},
		{
			Activated:   activated[:100],
			Deactivated: deactivated[:100],
			GCount:      10,
		},
		{
			Activated:   activated[:1000],
			Deactivated: deactivated[:1000],
			GCount:      10,
		},
		{
			Activated:   activated[:10000],
			Deactivated: deactivated[:10000],
			GCount:      10,
		},
		{
			Activated:   activated[:100000],
			Deactivated: deactivated[:100000],
			GCount:      10,
		},
		{
			Activated:   activated[:1],
			Deactivated: deactivated[:1],
			GCount:      100,
		},
		{
			Activated:   activated[:10],
			Deactivated: deactivated[:10],
			GCount:      100,
		},
		{
			Activated:   activated[:100],
			Deactivated: deactivated[:100],
			GCount:      100,
		},
		{
			Activated:   activated[:1000],
			Deactivated: deactivated[:1000],
			GCount:      100,
		},
		{
			Activated:   activated[:10000],
			Deactivated: deactivated[:10000],
			GCount:      100,
		},
		{
			Activated:   activated[:100000],
			Deactivated: deactivated[:100000],
			GCount:      100,
		},
		{
			Activated:   activated[:1],
			Deactivated: deactivated[:1],
			GCount:      1000,
		},
		{
			Activated:   activated[:10],
			Deactivated: deactivated[:10],
			GCount:      1000,
		},
		{
			Activated:   activated[:100],
			Deactivated: deactivated[:100],
			GCount:      1000,
		},
		{
			Activated:   activated[:1000],
			Deactivated: deactivated[:1000],
			GCount:      1000,
		},
		{
			Activated:   activated[:10000],
			Deactivated: deactivated[:10000],
			GCount:      1000,
		},
		{
			Activated:   activated[:100000],
			Deactivated: deactivated[:100000],
			GCount:      1000,
		},
	}

	for i := range cases {
		b.Run(fmt.Sprintf("input_size_%d_goroutines_count_%d", len(cases[i].Activated), cases[i].GCount), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				DeactivateUsers(cases[i].Activated, cases[i].GCount)
			}
		})
	}
}
