package main

import (
	"container/heap"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

func selection1(pq *PriorityQueue, isConsole bool) Selection1Struct {
	logger.Printf("getting selection 1, isConsole: %t", isConsole)
	tempArray := make([]IDJSON, 0)
	for i := 0; i < len(pq.harr); i++ {
		tempArray = append(tempArray, IDJSON{ID: pq.harr[i].ID})
	}
	oldest, _ := GetOldestTaskID(pq)
	s1Struct := Selection1Struct{QueueName: pq.queueName,
		QueueDescription: pq.queueDescription,
		Size:             len(pq.harr),
		OldestTaskID:     oldest,
		CustomerRequests: tempArray}

	if isConsole {
		jsonData, _ := json.MarshalIndent(s1Struct, "", "    ")
		fmt.Println(string(jsonData))
	}
	logger.Printf("returning selection 1, isConsole: %t", isConsole)
	return s1Struct
}

func selection2(pq *PriorityQueue, isConsole bool) Selection2Struct {
	logger.Printf("getting selection 2, isConsole: %t", isConsole)
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
	oldest, _ := GetOldestTaskID(pq)
	s2Struct := Selection2Struct{QueueName: pq.queueName,
		QueueDescription: pq.queueDescription,
		Size:             len(pq.harr),
		OldestTaskID:     oldest,
		CustomerRequests: tempArray}

	if isConsole {
		fmt.Println("Printing List of Customer Details in Queue")
		jsonData, _ := json.MarshalIndent(s2Struct, "", "    ")
		fmt.Println(string(jsonData))
	}
	logger.Printf("returning selection 2, isConsole: %t", isConsole)
	return s2Struct
}

func selection3(pq *PriorityQueue, isConsole bool) (Selection3Struct, ErrorStruct, error) {
	logger.Printf("getting selection 3, isConsole: %t", isConsole)
	if pq.count <= 0 {
		errorMsg := "Queue is empty."
		if isConsole {
			fmt.Println(errorMsg)
		}
		logger.Printf("error getting selection 3. %s isConsole: %t", errorMsg, isConsole)

		return Selection3Struct{}, ErrorStruct{Msg: errorMsg}, errors.New(errorMsg)
	}
	cr := heap.Pop(&pq.harr).(*CustomerRequest)
	pq.count--
	s3Struct := Selection3Struct{ID: cr.ID,
		PriorityWeight: cr.PriorityWeight,
		CustomerName:   cr.CustomerName,
		Description:    cr.Description,
		EnqueueTime:    cr.EnqueueTime,
		WaitTimeinSec:  time.Since(cr.EnqueueTime).Seconds()}

	if isConsole {
		fmt.Println("Dequeuing Customer Request")
		jsonData, _ := json.MarshalIndent(s3Struct, "", "    ")
		fmt.Println(string(jsonData))
	}
	logger.Printf("returning selection 3, isConsole: %t", isConsole)
	return s3Struct, ErrorStruct{}, nil
}

func selection4(pq *PriorityQueue, cr *CustomerRequest, isConsole bool) (Selection4Struct, error) {
	logger.Printf("getting selection 4, isConsole: %t", isConsole)
	if !insert(pq, cr, isConsole) {
		errorMsg := "The system is working at its peak capacity, please try again later."
		logger.Printf("error getting selection 4. %s isConsole: %t", errorMsg, isConsole)
		return Selection4Struct{}, errors.New(errorMsg)
	}

	s4Struct := Selection4Struct{ID: cr.ID,
		PriorityWeight:  cr.PriorityWeight,
		CustomerName:    cr.CustomerName,
		Description:     cr.Description,
		EnqueueTime:     cr.EnqueueTime,
		PositionInQueue: len(pq.harr) - 1}

	if isConsole {
		fmt.Printf("\nCustomer Request is enqueued with following information:\n")
		jsonData, _ := json.MarshalIndent(s4Struct, "", "    ")
		fmt.Println(string(jsonData))
	}
	logger.Printf("returning selection 4, isConsole: %t", isConsole)
	return s4Struct, nil
}

func selection5(pq *PriorityQueue, delID int, isConsole bool) (Selection5Struct, error) {
	logger.Printf("getting selection 5, isConsole: %t", isConsole)
	cr, err := GetCrByID(pq, delID)
	if err != nil {
		if isConsole {
			fmt.Println(err)
		}
		logger.Printf("error getting selection 5. %s, isConsole: %t", err.Error(), isConsole)
		return Selection5Struct{}, errors.New(err.Error())
	}
	maxInt := int(^uint(0) >> 1) // Make a MAX_INT
	cr.PriorityWeight = maxInt
	heap.Fix(&pq.harr, cr.index)
	_ = heap.Pop(&pq.harr).(*CustomerRequest)
	pq.count--

	s5Struct := Selection5Struct{
		CustomerName:  cr.CustomerName,
		ID:            cr.ID,
		EnqueueTime:   cr.EnqueueTime,
		WaitTimeinSec: time.Since(cr.EnqueueTime).Seconds(),
		Message:       "Request reneged successfully"}

	if isConsole {
		fmt.Println("Reneged following customer request sucessfully:")
		jsonData, _ := json.MarshalIndent(s5Struct, "", "    ")
		fmt.Println(string(jsonData))
	}
	logger.Printf("returning selection 5, could not delete %d, isConsole: %t", delID, isConsole)
	return s5Struct, nil
}

func selection6(pq *PriorityQueue, isConsole bool) (Selection6Struct, ErrorStruct, error) {
	logger.Printf("getting selection 6, isConsole: %t", isConsole)
	status := "IN_SERVICE"
	if len(pq.harr) >= pq.capacity {
		status = "MAX_CAPACITY_REACHED"
	}
	oldest, err := GetOldestTaskID(pq)
	if err != nil {
		if isConsole {
			fmt.Println(err)
		}
		logger.Printf("error getting selection 6. %s, isConsole: %t", err.Error(), isConsole)
		return Selection6Struct{}, ErrorStruct{Msg: err.Error()}, errors.New(err.Error())
	}
	oldestCr, err := GetCrByID(pq, oldest)
	if err != nil {
		if isConsole {
			fmt.Println(err)
		}
		logger.Printf("error getting selection 6. %s, isConsole: %t", err.Error(), isConsole)
		return Selection6Struct{}, ErrorStruct{Msg: err.Error()}, errors.New(err.Error())
	}
	queueInfo := QueueInfo{
		Name:                           pq.queueName,
		Size:                           strconv.Itoa(pq.count),
		OldestCustomerRequestTimeInSec: time.Since(oldestCr.EnqueueTime).Seconds()}
	s6Struct := Selection6Struct{
		Status: status,
		Queue:  queueInfo}

	if isConsole {
		jsonData, _ := json.MarshalIndent(s6Struct, "", "    ")
		fmt.Println(string(jsonData))
	}
	logger.Printf("returning selection 6, isConsole: %t", isConsole)
	return s6Struct, ErrorStruct{}, nil
}
