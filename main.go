// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Globals

// SIZE is the capacity of the priority queue
var SIZE = 50000

// PQ is the priority queue
var PQ = PriorityQueue{
	queueName:        "DefaultQueue",
	queueDescription: "This queue is for demonstration of Priority Queue implementation",
	capacity:         SIZE,
	key:              0,
	count:            0,
	isInitialized:    false}
var logger = initLogger()

// This example creates a Queue with some customerRequests, adds and manipulates an customerRequest,
// and then removes the customerRequests in PriorityWeight order.
func main() {
	logger.Println("logger started")
	logger.Println("making database with dummy data")
	priorities := make([]int, 0)

	for i := 0; i < SIZE; i++ {
		priorities = append(priorities, rand.Intn(10)+1)
	}

	for i := SIZE - 1; i >= 0; i-- {
		cr := &CustomerRequest{
			PriorityWeight: priorities[i],
			CustomerName:   "name" + strconv.Itoa(i),
			Description:    "desc" + strconv.Itoa(i),
			EnqueueTime:    time.Now(),
		}
		_ = insert(&PQ, cr, true)
	}
	go handleRequests()
	printHeader()

	logger.Println("entering selection mode")
	for true {
		printMenu()
		c := getSelection()

		switch c {
		case "1":
			_ = selection1(&PQ, true)
		case "2":
			_ = selection2(&PQ, true)
		case "3":
			selection3(&PQ, true)
		case "4":
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
			}
			_, _ = selection4(&PQ, cr, true)
		case "5":
			fmt.Printf("Please enter customer ID: ")
			tempStr := getInput()
			delID, _ := strconv.Atoi(tempStr)
			_, _ = selection5(&PQ, delID, true)
		case "6":
			selection6(&PQ, true)
		case "9":
			printMenu()
		case "0":
			return
		default:
			fmt.Printf("Invalid selection/n")
		}
		fmt.Println("")
	}

}
