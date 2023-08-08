
This package is an implementation of a Queue data structure in Golang

### Installation

To install the package use the following command

```sh
$ go get github.com/saudshaddad/queue
```

### Usage & Features

The queue package tries to be as efficient as possible. It uses a doubly linked list as it's back bone. You can add element to the queue and dequeue element from the existing function in the struct

**Time complexity of the functions**
- ``Enqueue(element any)``: O(1)
- ``Dequeue()``: O(1)
- ``Peak()``: O(1)
- ``PeakBack()``: O(1)
- ``IsEmpty()``: O(1)
- ``Contains(element any)``: O(n)

The queue uses generics, you can pass whatever data type you want (it is recommended to pass in pointers)

### Getting Started

You only need to import the package to use the Queue and initialize a new queue.

```go
package main

import (
  "github.com/saudshaddad/queue"
)

type User struct {
	Name string
	Mobile string
	Id int
}

func main() {
	q := queue.NewQueue[*User]() // initialize the queue with your data type

	// add new elements to the queue
	user1 := &User{
		Name: "Saud Shaddad",
		Mobile: "66778844",
		Id: 1,
	}
	
	user2 := &User{
		Name: "Saud Khaled",
		Mobile: "88776655",
		Id: 2,
	}
	
	q.Enqueue(user1)
	q.Enqueue(user2)

	// for loop against the queue until it is empty
	for !q.IsEmpty() {
		// fetch the next element and remove it from the queue
		elem, err := q.Dequeue()  
		if err != nil {
			// you can check if the error is due to empty queue
			if err == queue.ErrorEmptyQueue {
				// do something like breaking
				// note: this error checking if not necessary in this case since we are checking if the queue is empty in the for loop condition but you can use this if you have different implementation
				break
			}
		}
		
		// process the element ... 
	}
}
```

### Notes

The queue structure is not concurrent safe, you should use a mutex if you are passing the queue to different go routines
