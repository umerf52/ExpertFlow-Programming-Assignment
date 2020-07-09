// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"bufio"
	"container/heap"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// An CustomerRequest is something we manage in a priority queue.
type CustomerRequest struct {
	ID                        int
	CustomerName, Description string // The value of the customerRequest; arbitrary.
	PriorityWeight            int    // The priority of the customerRequest in the queue.
	EnqueueTime               time.Time
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the customerRequest in the heap.
}

// A Queue implements heap.Interface and holds CustomerRequests.
type Queue []*CustomerRequest

// PriorityQueue wraps the actual priority queue and provIDes additional functionality
type PriorityQueue struct {
	harr                        Queue
	queueName, queueDescription string
}

type IDJSON struct {
	ID int
}

type Selection1Struct struct {
	QueueName, QueueDescription string
	Size, OldestTaskID          int
	CustomerRequests            []IDJSON
}

type Selection2Struct struct {
	QueueName, QueueDescription string
	Size, OldestTaskID          int
	CustomerRequests            []*CustomerRequest
}

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

func printMenu() {
	fmt.Println("*************************************************************")
	fmt.Println("* Welcome to Priority Queue, please select from following   *")
	fmt.Println("*************************************************************")
	fmt.Println("1. List Customers in Queue")
	fmt.Println("2. List Customer Details in Queue")
	fmt.Println("3. Service Customer")
	fmt.Println("4. Enqueue Customer Request")
	fmt.Println("5. Renege Customer Request")
	fmt.Println("6. System Information")
	fmt.Println("7. System Memory Dump")
	fmt.Println("8. Reprint Menu")
	fmt.Println("0. Exit")
	fmt.Println("")
}

func getInput() string {
	fmt.Printf("Enter selection: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return string([]byte(input)[0])
}

func getOldestTaskID(pq *PriorityQueue) int {
	oldestID := pq.harr[0].ID
	oldestTime := pq.harr[0].EnqueueTime
	for i := 1; i < len(pq.harr); i++ {
		cr := pq.harr[i]
		if cr.EnqueueTime.Before(oldestTime) {
			oldestTime = cr.EnqueueTime
			oldestID = cr.ID
		}
	}
	return oldestID
}

func selection1(pq *PriorityQueue) {
	tempArray := make([]IDJSON, 0)
	for i := 0; i < len(pq.harr); i++ {
		tempArray = append(tempArray, IDJSON{ID: pq.harr[i].ID})
	}
	s1Struct := Selection1Struct{QueueName: pq.queueName,
		QueueDescription: pq.queueDescription,
		Size:             len(pq.harr),
		OldestTaskID:     getOldestTaskID(pq),
		CustomerRequests: tempArray}

	jsonData, _ := json.MarshalIndent(s1Struct, "", "    ")
	fmt.Println(string(jsonData))
}

func selection2(pq *PriorityQueue) {
	tempArray := make([]*CustomerRequest, 0)
	for i := 0; i < len(pq.harr); i++ {
		cr := &CustomerRequest{
			ID:             pq.harr[i].ID,
			PriorityWeight: pq.harr[i].PriorityWeight,
			CustomerName:   pq.harr[i].CustomerName,
			Description:    pq.harr[i].Description,
			EnqueueTime:    pq.harr[i].EnqueueTime,
			index:          pq.harr[i].index,
		}
		tempArray = append(tempArray, cr)
	}
	s2Struct := Selection2Struct{QueueName: pq.queueName,
		QueueDescription: pq.queueDescription,
		Size:             len(pq.harr),
		OldestTaskID:     getOldestTaskID(pq),
		CustomerRequests: tempArray}

	jsonData, _ := json.MarshalIndent(s2Struct, "", "    ")
	fmt.Println(string(jsonData))
}

// This example creates a Queue with some customerRequests, adds and manipulates an customerRequest,
// and then removes the customerRequests in PriorityWeight order.
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

	for i := 9; i >= 0; i-- {
		cr := &CustomerRequest{
			ID:             ids[i],
			PriorityWeight: priorities[i],
			CustomerName:   names[i],
			Description:    descs[i],
			EnqueueTime:    time.Now(),
			index:          i,
		}
		fmt.Println(cr.ID, cr.PriorityWeight, cr.CustomerName, cr.Description, cr.EnqueueTime)
		pq.harr[i] = cr
		time.Sleep(500000000)
	}
	fmt.Println("")

	heap.Init(&pq.harr)
	printMenu()
	for true {
		c := getInput()

		switch c {
		case "1":
			selection1(&pq)
		case "2":
			selection2(&pq)
		case "8":
			printMenu()
		case "0":
			return
		default:
			fmt.Println("Invalid selection")
		}
	}

}
