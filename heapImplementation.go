package main

import (
	"container/heap"
	"fmt"
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

func insert(pq *PriorityQueue, cr *CustomerRequest, isConsole bool) bool {
	logger.Printf("inserting Customer Request")
	if pq.count >= pq.capacity {
		errorMsg := "Capacity reached. Could not insert.\n\n"
		if isConsole {
			fmt.Printf(errorMsg)
		}
		logger.Printf("ERROR: inserting Customer Request. %s", errorMsg)
		return false
	}
	cr.ID = pq.key
	cr.index = len(pq.harr)
	pq.key++
	if !pq.isInitialized {
		pq.harr = make(Queue, 1)
		pq.harr[0] = cr
		heap.Init(&pq.harr)
		pq.isInitialized = true
	} else {
		heap.Push(&pq.harr, cr)
	}
	if isConsole {
		fmt.Println(cr.ID, cr.PriorityWeight, cr.CustomerName, cr.Description, cr.EnqueueTime)
	}
	pq.count++
	logger.Printf("successfully inserted")
	return true
}
