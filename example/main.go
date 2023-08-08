package main

import (
	"fmt"

	"github.com/saudshaddad/queue"
)

type User struct {
	Id     int
	name   string
	mobile string
}

func main() {

	// make the queue with your chosen type
	que := queue.NewQueue[*User]()

	// fill the queue
	index := 0
	for {
		if index == 10000000 {
			break
		}

		user1 := &User{
			name:   "saud shaddad",
			mobile: "66766541",
			Id:     index,
		}

		que.Enqueue(user1)

		index = index + 1
	}

	// dequeue an element the process it
	element, err := que.Dequeue()
	if err != nil {
		// handle the error
		if err == queue.ErrorEmptyQueue {
			// handle the error if the queue was empty, possibly break from a for loop
			return
		}
	}
	// process the dequeued element
	fmt.Println(element)

	// Peak into the next element without removing it from the queue
	elem, err := que.Peak()
	if err != nil {
		if err == queue.ErrorEmptyQueue {
			fmt.Println("empty queue")
			return
		}

		fmt.Println(err)
		return
	}
	fmt.Println(elem)
}
