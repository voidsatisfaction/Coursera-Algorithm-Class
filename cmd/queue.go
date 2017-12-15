package cmd

import (
	"fmt"

	"../pkg/queue"
)

func QueueMain() {
	q := &queue.StringListQueue{}
	q.Enqueue("hi")
	fmt.Printf("%+v", q)
	q.Dequeue()
	fmt.Printf("%+v", q)
}
