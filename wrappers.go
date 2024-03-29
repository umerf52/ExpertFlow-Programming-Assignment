package main

import (
	"container/heap"
	"errors"
	"fmt"
)

// Wrapper function to insert into Priority Queue
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
	pq.count++
	logger.Printf("successfully inserted following:")
	logger.Println(cr.ID, cr.PriorityWeight, cr.CustomerName, cr.Description, cr.EnqueueTime)
	return true
}

// This function returns CustomerRequest with highest PriorityWeight
func extractMax(pq *PriorityQueue) *CustomerRequest {
	cr := heap.Pop(&pq.harr).(*CustomerRequest) // Remove the CustomerRequest with highest PriorityWeight
	pq.count--
	return cr
}

// This function deleted the CustomerRequest with id=delID
func deleteByID(pq *PriorityQueue, delID int, isConsole bool) (*CustomerRequest, error) {
	cr, err := getCrByID(pq, delID)
	if err != nil {
		logger.Printf("error in deleteById. %s, isConsole: %t", err.Error(), isConsole)
		return &CustomerRequest{}, errors.New(err.Error())
	}
	maxInt := int(^uint(0) >> 1) // Make a MAX_INT
	cr.PriorityWeight = maxInt
	heap.Fix(&pq.harr, cr.index) // Move the CustomerRequest with highest PriorityWeight to the top
	_ = heap.Pop(&pq.harr).(*CustomerRequest)
	pq.count--

	return cr, nil
}
