package queue

import (
	"container/list"
	"errors"
)

var (
	ErrorEmptyQueue = errors.New("QueueErrorEmpty")
)

// represent a queue data structure (backed by a doubley linked list)
//
// The queue is not a concurrent safe struct, don't use it multiple go routine without a mutex
type Queue[T comparable] struct {
	doubleList *list.List
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{
		doubleList: list.New(),
	}
}

// Adds new element to the queue
//
// The time complexity is O(1)
func (q *Queue[T]) Enqueue(element T) {
	q.doubleList.PushFront(element)
}

// Returns the element that needs to be served in the queue and removes it from the queue.
// The function will return an error QueueErrorEmpty if the queue is empty
//
// The time complexity is O(1)
func (q *Queue[T]) Dequeue() (v T, e error) {
	rmElem := q.doubleList.Back()
	if rmElem == nil {
		return v, ErrorEmptyQueue
	}
	q.doubleList.Remove(rmElem)
	return rmElem.Value.(T), nil
}

// Returns the first element of the queue
// i.e. the first element that was added to the queue
//
// The time complexity is O(1)
func (q *Queue[T]) Peak() (v T, er error) {
	elem := q.doubleList.Back()
	if elem == nil {
		return v, ErrorEmptyQueue
	}
	return elem.Value.(T), nil
}

// Returns the last element of the queue
// i.e. the last element that was added to the queue
//
// The time complexity is O(1)
func (q *Queue[T]) PeakBack() (v T, er error) {
	elem := q.doubleList.Front()
	if elem == nil {
		return v, ErrorEmptyQueue
	}
	return elem.Value.(T), nil
}

// Returns a boolean if the queue is empty
//
// The time complexity is O(1)
func (q *Queue[T]) IsEmpty() bool {
	return q.doubleList.Len() == 0
}

// Determines if the queue contains the element, if queue is empty it will returns false
//
// The time complexity is O(n)
func (q *Queue[T]) Contains(element T) bool {
	first := q.doubleList.Front()
	if first == nil {
		return false
	}

	value := first.Value.(T)
	if value == element {
		return true
	}

	for {
		first = first.Next()
		if first == nil {
			return false
		}

		value := first.Value.(T)
		if value == element {
			return true
		}

	}
}
