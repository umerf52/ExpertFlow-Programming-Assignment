// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
	"time"
)

// An CustomerRequest is something we manage in a priority queue.
type CustomerRequest struct {
	customerName, description string // The value of the customerRequest; arbitrary.
	priorityWeight, id        int    // The priority of the customerRequest in the queue.
	enqueueTime               time.Time
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the customerRequest in the heap.
}

// A Queue implements heap.Interface and holds CustomerRequests.
type Queue []*CustomerRequest

// PriorityQueue wraps the actual priority queue and provides additional functionality
type PriorityQueue struct {
	harr                        Queue
	queueName, queueDescription string
}

func (q Queue) Len() int { return len(q) }

func (q Queue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priorityWeight so we use greater than here.
	return q[i].priorityWeight > q[j].priorityWeight
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

// Push : Implementation of Heap's Push()
func (q *Queue) Push(x interface{}) {
	n := len(*q)
	customerRequest := x.(*CustomerRequest)
	customerRequest.index = n
	*q = append(*q, customerRequest)
}

// Pop : Implementation of Heap's Pop()
func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	customerRequest := old[n-1]
	old[n-1] = nil             // avoid memory leak
	customerRequest.index = -1 // for safety
	*q = old[0 : n-1]
	return customerRequest
}

// update modifies the priorityWeight and value of an CustomerRequest in the queue.
func (q *Queue) update(customerRequest *CustomerRequest, description string, priorityWeight int) {
	customerRequest.description = description
	customerRequest.priorityWeight = priorityWeight
	heap.Fix(q, customerRequest.index)
}

// This example creates a Queue with some customerRequests, adds and manipulates an customerRequest,
// and then removes the customerRequests in priorityWeight order.
func main() {
	// Some customerRequests and their priorities.
	customerRequests := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the customerRequests in it, and
	// establish the priority queue (heap) invariants.
	q := make(Queue, len(customerRequests))
	i := 0
	for value, priority := range customerRequests {
		q[i] = &CustomerRequest{
			description:    value,
			priorityWeight: priority,
			index:          i,
			enqueueTime:    time.Now(),
		}
		i++
	}
	heap.Init(&q)

	priorityQueue := PriorityQueue{harr: q,
		queueName:        "DefaultQueue",
		queueDescription: "This queue is for demonstration of Priority Queue Implementation"}

	fmt.Println(priorityQueue.queueName, priorityQueue.queueDescription)
	// Insert a new customerRequest and then modify its priorityWeight.
	customerRequest := &CustomerRequest{
		description:    "orange",
		priorityWeight: 1,
	}
	heap.Push(&q, customerRequest)
	q.update(customerRequest, customerRequest.description, 5)

	// Take the customerRequests out; they arrive in decreasing priorityWeight order.
	for q.Len() > 0 {
		customerRequest := heap.Pop(&q).(*CustomerRequest)
		// fmt.Printf("%.2d:%s ", customerRequest.priorityWeight, customerRequest.description)
		fmt.Println(customerRequest.priorityWeight, customerRequest.description, customerRequest.enqueueTime)
	}
}
