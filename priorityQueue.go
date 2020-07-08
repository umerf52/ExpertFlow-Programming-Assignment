// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
	"math/rand"
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
	ids, priorities := make([]int, 0), make([]int, 0)
	names, descs := make([]string, 0), make([]string, 0)

	for i := 0; i < 10; i++ {
		ids = append(ids, i)
		priorities = append(priorities, rand.Intn(10)+1)
		names = append(names, "one")
		descs = append(descs, "")
	}

	pq := PriorityQueue{
		queueName:        "DefaultQueue",
		queueDescription: "This queue is for demonstration of Priority Queue implementation"}

	// Create a priority queue, put the customerRequests in it, and
	// establish the priority queue (heap) invariants.
	pq.harr = make(Queue, 10)

	for i := 0; i < 10; i++ {
		pq.harr[i] = &CustomerRequest{
			id:             ids[i],
			priorityWeight: priorities[i],
			customerName:   names[i],
			description:    descs[i],
			enqueueTime:    time.Now(),
			index:          i,
		}
	}

	heap.Init(&pq.harr)

	fmt.Println(pq.queueName, pq.queueDescription)

	// Take the customerRequests out; they arrive in decreasing priorityWeight order.
	for pq.harr.Len() > 0 {
		cr := heap.Pop(&pq.harr).(*CustomerRequest)
		fmt.Println(cr.priorityWeight, cr.customerName, cr.description, cr.enqueueTime)
	}
}
