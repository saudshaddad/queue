package queue_test

import (
	"testing"

	"github.com/saudshaddad/queue"
)

func TestEnqueue(t *testing.T) {

	q := queue.NewQueue[*int]()
	n1 := 1
	n2 := 2

	q.Enqueue(&n1)
	q.Enqueue(&n2)

	elem, err := q.Peak()
	if err != nil {
		t.FailNow()
	}

	if *elem != n1 {
		t.FailNow()
	}

	if !q.Contains(&n1) {
		t.FailNow()
	}
}

func TestDequeue(t *testing.T) {

	q := queue.NewQueue[*int]()
	n1 := 1
	n2 := 2

	_, err := q.Dequeue()
	if err != queue.ErrorEmptyQueue {
		t.FailNow()
	}

	q.Enqueue(&n1)
	q.Enqueue(&n2)

	_, err = q.Dequeue()
	if err != nil {
		t.FailNow()
	}

	// the element n1 should be removed from the queue after dequeueing
	peak, _ := q.Peak()
	if peak == &n1 {
		t.FailNow()
	}

	// the element n1 should not exist in the queue any more
	if q.Contains(&n1) {
		t.FailNow()
	}
}

func TestPeak(t *testing.T) {

	q := queue.NewQueue[*int]()
	n1 := 1
	n2 := 2
	n3 := 3

	_, err := q.Peak()
	if err != queue.ErrorEmptyQueue {
		t.FailNow()
	}

	q.Enqueue(&n1)
	q.Enqueue(&n2)
	q.Enqueue(&n3)

	peak, err := q.Peak()
	if err != nil {
		t.FailNow()
	}

	if peak != &n1 {
		t.FailNow()
	}
}

func TestPeakBack(t *testing.T) {

	q := queue.NewQueue[*int]()
	n1 := 1
	n2 := 2
	n3 := 3

	_, err := q.PeakBack()
	if err != queue.ErrorEmptyQueue {
		t.FailNow()
	}

	q.Enqueue(&n1)
	q.Enqueue(&n2)
	q.Enqueue(&n3)

	peak, err := q.PeakBack()
	if err != nil {
		t.FailNow()
	}

	if peak != &n3 {
		t.FailNow()
	}
}

func TestIsEmpty(t *testing.T) {

	q := queue.NewQueue[*int]()
	n1 := 1
	n2 := 2
	n3 := 3

	if !q.IsEmpty() {
		t.FailNow()
	}

	q.Enqueue(&n1)
	q.Enqueue(&n2)
	q.Enqueue(&n3)

	if q.IsEmpty() {
		t.FailNow()
	}

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()

	if !q.IsEmpty() {
		t.FailNow()
	}
}

func TestContains(t *testing.T) {

	q := queue.NewQueue[*int]()
	n1 := 1
	n2 := 2
	n3 := 3

	if q.Contains(&n1) {
		t.FailNow()
	}

	q.Enqueue(&n1)
	if !q.Contains(&n1) {
		t.FailNow()
	}
	if q.Contains(&n2) {
		t.FailNow()
	}

	q.Enqueue(&n2)
	q.Enqueue(&n3)

	if !q.Contains(&n3) {
		t.FailNow()
	}
}
