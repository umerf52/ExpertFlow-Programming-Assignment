// This is Priority Queue implementation of heap
// The skeleton code was taken from Go's official documentation:
// https://golang.org/pkg/container/heap/

package main

import (
	"container/heap"
)

func (q Queue) Len() int { return len(q) }

func (q Queue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, PriorityWeight so we use greater than here.
	return q[i].PriorityWeight > q[j].PriorityWeight
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
	if q.Len() <= 0 {
		return CustomerRequest{}
	}
	old := *q
	n := len(old)
	customerRequest := old[n-1]
	old[n-1] = nil             // avoid memory leak
	customerRequest.index = -1 // for safety
	*q = old[0 : n-1]
	return customerRequest
}

// update modifies the PriorityWeight and value of an CustomerRequest in the queue.
func (q *Queue) update(customerRequest *CustomerRequest, description string, priorityWeight int) {
	customerRequest.Description = description
	customerRequest.PriorityWeight = priorityWeight
	heap.Fix(q, customerRequest.index)
}
