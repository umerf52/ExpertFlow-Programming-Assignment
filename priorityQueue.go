// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"bufio"
	"container/heap"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// An CustomerRequest is something we manage in a priority queue.
type CustomerRequest struct {
	ID                        int
	CustomerName, Description string
	PriorityWeight            int // The priority of the customerRequest in the queue.
	EnqueueTime               time.Time
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the customerRequest in the heap.
}

// A Queue implements heap.Interface and holds CustomerRequests.
type Queue []*CustomerRequest

// PriorityQueue wraps the actual priority queue and provides additional functionality
type PriorityQueue struct {
	harr                        Queue
	queueName, queueDescription string
	capacity, count, key        int
	isInitialized               bool
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

type Selection3Struct struct {
	ID                        int
	CustomerName, Description string
	PriorityWeight            int // The priority of the customerRequest in the queue.
	EnqueueTime               time.Time
	WaitTimeinSec             float64
}

type Selection4Struct struct {
	CustomerName, Description string
	PriorityWeight, ID        int // The priority of the customerRequest in the queue.
	EnqueueTime               time.Time
	PositionInQueue           int
}

type Selection5Struct struct {
	CustomerName  string
	ID            int
	EnqueueTime   time.Time
	WaitTimeinSec float64
	Message       string
}

type QueueInfo struct {
	Name, Size                     string
	OldestCustomerRequestTimeInSec float64
}

type Selection6Struct struct {
	Status string
	Queue  QueueInfo
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

func printHeader() {
	fmt.Println("*************************************************************")
	fmt.Println("* Welcome to Priority Queue, please select from following   *")
	fmt.Println("*************************************************************")
}

func printMenu() {
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

func getSelection() string {
	fmt.Printf("Enter selection: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return string([]byte(input)[0])
}

func getOldestTaskID(pq *PriorityQueue) (int, error) {
	if pq.count <= 0 {
		return -1, errors.New("Queue is empty.")
	}
	oldestID := pq.harr[0].ID
	oldestTime := pq.harr[0].EnqueueTime
	for i := 1; i < len(pq.harr); i++ {
		cr := pq.harr[i]
		if cr.EnqueueTime.Before(oldestTime) {
			oldestTime = cr.EnqueueTime
			oldestID = cr.ID
		}
	}
	return oldestID, nil
}

func selection1(pq *PriorityQueue) {
	tempArray := make([]IDJSON, 0)
	for i := 0; i < len(pq.harr); i++ {
		tempArray = append(tempArray, IDJSON{ID: pq.harr[i].ID})
	}
	oldest, _ := getOldestTaskID(pq)
	s1Struct := Selection1Struct{QueueName: pq.queueName,
		QueueDescription: pq.queueDescription,
		Size:             len(pq.harr),
		OldestTaskID:     oldest,
		CustomerRequests: tempArray}

	jsonData, _ := json.MarshalIndent(s1Struct, "", "    ")
	fmt.Println(string(jsonData))
}

func selection2(pq *PriorityQueue) {
	fmt.Println("Printing List of Customer Details in Queue")
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
	oldest, _ := getOldestTaskID(pq)
	s2Struct := Selection2Struct{QueueName: pq.queueName,
		QueueDescription: pq.queueDescription,
		Size:             len(pq.harr),
		OldestTaskID:     oldest,
		CustomerRequests: tempArray}

	jsonData, _ := json.MarshalIndent(s2Struct, "", "    ")
	fmt.Println(string(jsonData))
}

func selection3(pq *PriorityQueue) {
	if pq.count <= 0 {
		fmt.Println("Queue is empty.")
		return
	}
	fmt.Println("Dequeuing Customer Request")
	cr := heap.Pop(&pq.harr).(*CustomerRequest)
	pq.count--
	s3Struct := Selection3Struct{ID: cr.ID,
		PriorityWeight: cr.PriorityWeight,
		CustomerName:   cr.CustomerName,
		Description:    cr.Description,
		EnqueueTime:    cr.EnqueueTime,
		WaitTimeinSec:  time.Since(cr.EnqueueTime).Seconds()}

	jsonData, _ := json.MarshalIndent(s3Struct, "", "    ")
	fmt.Println(string(jsonData))
}

func getInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	temp := scanner.Text()
	return temp
}

func selection4(pq *PriorityQueue) {
	fmt.Println("Please enter following information: ")
	fmt.Printf("Customer Name: ")
	name := getInput()
	fmt.Printf("Description: ")
	desc := getInput()
	fmt.Printf("Priority Weight: ")
	priorityStr := getInput()
	priorityInt, _ := strconv.Atoi(priorityStr)
	cr := &CustomerRequest{
		PriorityWeight: priorityInt,
		CustomerName:   name,
		Description:    desc,
		EnqueueTime:    time.Now(),
		index:          len(pq.harr),
	}
	if !insert(pq, cr) {
		return
	}

	fmt.Printf("\nCustomer Request is enqueued with following information:\n")
	s4Struct := Selection4Struct{ID: cr.ID,
		PriorityWeight:  cr.PriorityWeight,
		CustomerName:    cr.CustomerName,
		Description:     cr.Description,
		EnqueueTime:     cr.EnqueueTime,
		PositionInQueue: len(pq.harr) - 1}

	jsonData, _ := json.MarshalIndent(s4Struct, "", "    ")
	fmt.Println(string(jsonData))
}

func getCrByID(pq *PriorityQueue, ID int) (*CustomerRequest, error) {
	for i := 0; i < len(pq.harr); i++ {
		if pq.harr[i].ID == ID {
			return pq.harr[i], nil
		}
	}
	return nil, errors.New("ID not found.")
}

func selection5(pq *PriorityQueue) {
	fmt.Printf("Please enter customer ID: ")
	tempStr := getInput()
	delID, _ := strconv.Atoi(tempStr)
	cr, err := getCrByID(pq, delID)
	if err != nil {
		fmt.Println(err)
		return
	}
	maxInt := int(^uint(0) >> 1)
	cr.PriorityWeight = maxInt
	heap.Fix(&pq.harr, cr.index)
	_ = heap.Pop(&pq.harr).(*CustomerRequest)
	pq.count--
	fmt.Println("Reneged following customer request sucessfully:")
	s5Struct := Selection5Struct{
		CustomerName:  cr.CustomerName,
		ID:            cr.ID,
		EnqueueTime:   cr.EnqueueTime,
		WaitTimeinSec: time.Since(cr.EnqueueTime).Seconds(),
		Message:       "Request reneged successfully"}

	jsonData, _ := json.MarshalIndent(s5Struct, "", "    ")
	fmt.Println(string(jsonData))
}

func selection6(pq *PriorityQueue) {
	status := "IN_SERVICE"
	if len(pq.harr) >= pq.capacity {
		status = "MAX_CAPACITY_REACHED"
	}
	oldest, err := getOldestTaskID(pq)
	if err != nil {
		fmt.Println(err)
		return
	}
	oldestCr, err := getCrByID(pq, oldest)
	if err != nil {
		fmt.Println(err)
		return
	}
	queueInfo := QueueInfo{
		Name:                           pq.queueName,
		Size:                           strconv.Itoa(pq.count),
		OldestCustomerRequestTimeInSec: time.Since(oldestCr.EnqueueTime).Seconds()}
	s6Struct := Selection6Struct{
		Status: status,
		Queue:  queueInfo}

	jsonData, _ := json.MarshalIndent(s6Struct, "", "    ")
	fmt.Println(string(jsonData))
}

func insert(pq *PriorityQueue, cr *CustomerRequest) bool {
	if pq.count >= pq.capacity {
		fmt.Printf("Capacity reached. Could not insert.\n\n")
		return false
	}
	cr.ID = pq.key
	pq.key++
	if !pq.isInitialized {
		pq.harr = make(Queue, 1)
		pq.harr[0] = cr
		heap.Init(&pq.harr)
		pq.isInitialized = true
	} else {
		heap.Push(&pq.harr, cr)
	}
	fmt.Println(cr.ID, cr.PriorityWeight, cr.CustomerName, cr.Description, cr.EnqueueTime)
	pq.count++
	return true
}

// This example creates a Queue with some customerRequests, adds and manipulates an customerRequest,
// and then removes the customerRequests in PriorityWeight order.
func main() {
	size := 10
	priorities := make([]int, 0)
	names, descs := make([]string, 0), make([]string, 0)

	for i := 0; i < size; i++ {
		priorities = append(priorities, rand.Intn(10)+1)
		names = append(names, "one")
		descs = append(descs, "")
	}

	pq := PriorityQueue{
		queueName:        "DefaultQueue",
		queueDescription: "This queue is for demonstration of Priority Queue implementation",
		capacity:         size,
		key:              0,
		count:            0,
		isInitialized:    false}

	// Create a priority queue, put the customerRequests in it, and
	// establish the priority queue (heap) invariants.

	for i := size - 1; i >= 0; i-- {
		cr := &CustomerRequest{
			PriorityWeight: priorities[i],
			CustomerName:   names[i],
			Description:    descs[i],
			EnqueueTime:    time.Now(),
			index:          i,
		}
		_ = insert(&pq, cr)
		time.Sleep(500000000)
	}
	fmt.Println("")
	printHeader()

	for true {
		printMenu()
		c := getSelection()

		switch c {
		case "1":
			selection1(&pq)
		case "2":
			selection2(&pq)
		case "3":
			selection3(&pq)
		case "4":
			selection4(&pq)
		case "5":
			selection5(&pq)
		case "6":
			selection6(&pq)
		case "8":
			printMenu()
		case "0":
			return
		default:
			fmt.Printf("Invalid selection")
		}
		fmt.Println("")
	}

}
