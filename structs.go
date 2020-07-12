package main

import "time"

// An CustomerRequest is something we manage in a priority queue.
type CustomerRequest struct {
	ID             int       `json:"id"`
	CustomerName   string    `json:"customerName"`
	Description    string    `json:"description"`
	PriorityWeight int       `json:"priorityWeight"`
	EnqueueTime    time.Time `json:"enqueueTime"`
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
	ID int `json:"id"`
}

type Selection1Struct struct {
	QueueName        string   `json:"queueName"`
	QueueDescription string   `json:"queueDescription"`
	Size             int      `json:"size"`
	OldestTaskID     int      `json:"oldestTaskId"`
	CustomerRequests []IDJSON `json:"customerRequests"`
}

type Selection2Struct struct {
	QueueName        string             `json:"queueName"`
	QueueDescription string             `json:"queueDescription"`
	Size             int                `json:"size"`
	OldestTaskID     int                `json:"oldestTaskId"`
	CustomerRequests []*CustomerRequest `json:"customerRequests"`
}

type Selection3Struct struct {
	ID             int       `json:"id"`
	CustomerName   string    `json:"customerName"`
	Description    string    `json:"description"`
	PriorityWeight int       `json:"priorityWeight"`
	EnqueueTime    time.Time `json:"enqueueTime"`
	WaitTimeinSec  float64   `json:"waitTimeinSec"`
}

type Selection4Struct struct {
	CustomerName    string    `json:"customerName"`
	Description     string    `json:"description"`
	PriorityWeight  int       `json:"priorityWeight"`
	ID              int       `json:"id"`
	EnqueueTime     time.Time `json:"enqueueTime"`
	PositionInQueue int       `json:"positionInQueue"`
}

type Selection5Struct struct {
	CustomerName  string    `json:"customerName"`
	ID            int       `json:"id"`
	EnqueueTime   time.Time `json:"enqueueTime"`
	WaitTimeinSec float64   `json:"waitTimeinSec"`
	Message       string    `json:"message"`
}

type QueueInfo struct {
	Name                           string  `json:"name"`
	Size                           string  `json:"size"`
	OldestCustomerRequestTimeInSec float64 `json:"oldestCustomerRequestTimeInSec"`
}

type Selection6Struct struct {
	Status string    `json:"status"`
	Queue  QueueInfo `json:"queue"`
}

type ErrorStruct struct {
	Msg string `json:"message"`
}

type Selection4ErrorStruct struct {
	Error string `json:"error"`
	Msg   string `json:"message"`
}
