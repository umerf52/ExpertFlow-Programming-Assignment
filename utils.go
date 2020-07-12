package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func PrintHeader() {
	fmt.Println("*************************************************************")
	fmt.Println("* Welcome to Priority Queue, please select from following   *")
	fmt.Println("*************************************************************")
}

func PrintMenu() {
	fmt.Println("1. List Customers in Queue")
	fmt.Println("2. List Customer Details in Queue")
	fmt.Println("3. Service Customer")
	fmt.Println("4. Enqueue Customer Request")
	fmt.Println("5. Renege Customer Request")
	fmt.Println("6. System Information")
	fmt.Println("9. Reprint Menu")
	fmt.Println("0. Exit")
	fmt.Println("")
}

func GetSelection() string {
	fmt.Printf("Enter selection: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return string([]byte(input)[0])
}

func GetOldestTaskID(pq *PriorityQueue) (int, error) {
	if pq.count <= 0 {
		return -1, errors.New("queue is empty")
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

func GetInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	temp := scanner.Text()
	return temp
}

func GetCrByID(pq *PriorityQueue, ID int) (*CustomerRequest, error) {
	for i := 0; i < len(pq.harr); i++ {
		if pq.harr[i].ID == ID {
			return pq.harr[i], nil
		}
	}
	logger.Printf("id %d not found", ID)
	return nil, errors.New("id not found")
}
