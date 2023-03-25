package utilities_test

import (
	"testing"

	i "github.com/tsanton/goflake-client/goflake/models/assets/interface"
	u "github.com/tsanton/goflake-client/goflake/utilities"
)

func Test_queue_is_empty(t *testing.T) {
	/* Arrange */
	q := u.Queue[i.ISnowflakeAsset]{}

	/* Act and Assert */
	if !q.IsEmpty() {
		t.FailNow()
	}
}

func Test_queue_put(t *testing.T) {
	/* Arrange */
	q := u.Queue[int]{}

	/* Act */
	q.Put(1)

	/* Act and Assert */
	if q.IsEmpty() {
		t.FailNow()
	}
}

func Test_queue_order(t *testing.T) {
	/* Arrange */
	s := u.Queue[int]{}
	s.Put(1)
	s.Put(2)
	s.Put(3)

	/* Act */
	g1 := s.Get()
	g2 := s.Get()
	g3 := s.Get()

	/* Assert */
	if g1 != 1 {
		t.FailNow()
	}

	if g2 != 2 {
		t.FailNow()
	}

	if g3 != 3 {
		t.FailNow()
	}

	if !s.IsEmpty() {
		t.FailNow()
	}
}
