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
	harr                        Queue // harr is a Queue that implements heap interface
	queueName, queueDescription string
	capacity, count, key        int  // key is used to uniquely identify CustomerRequests
	isInitialized               bool // it is used to check if the at least one item has been inserted in harr or not
}

// IDJSON is used to in Selection1Struct
type IDJSON struct {
	ID int `json:"id"`
}

// Selection1Struct is the struct to represent selection 1
type Selection1Struct struct {
	QueueName        string   `json:"queueName"`
	QueueDescription string   `json:"queueDescription"`
	Size             int      `json:"size"`
	OldestTaskID     int      `json:"oldestTaskId"`
	CustomerRequests []IDJSON `json:"customerRequests"`
}

// Selection2Struct is the struct to represent selection 2
type Selection2Struct struct {
	QueueName        string             `json:"queueName"`
	QueueDescription string             `json:"queueDescription"`
	Size             int                `json:"size"`
	OldestTaskID     int                `json:"oldestTaskId"`
	CustomerRequests []*CustomerRequest `json:"customerRequests"`
}

// Selection3Struct is the struct to represent selection 3
type Selection3Struct struct {
	ID             int       `json:"id"`
	CustomerName   string    `json:"customerName"`
	Description    string    `json:"description"`
	PriorityWeight int       `json:"priorityWeight"`
	EnqueueTime    time.Time `json:"enqueueTime"`
	WaitTimeinSec  float64   `json:"waitTimeinSec"`
}

// Selection4Struct is the struct to represent selection 4
type Selection4Struct struct {
	CustomerName    string    `json:"customerName"`
	Description     string    `json:"description"`
	PriorityWeight  int       `json:"priorityWeight"`
	ID              int       `json:"id"`
	EnqueueTime     time.Time `json:"enqueueTime"`
	PositionInQueue int       `json:"positionInQueue"`
}

// Selection5Struct is the struct to represent selection 5
type Selection5Struct struct {
	CustomerName  string    `json:"customerName"`
	ID            int       `json:"id"`
	EnqueueTime   time.Time `json:"enqueueTime"`
	WaitTimeinSec float64   `json:"waitTimeinSec"`
	Message       string    `json:"message"`
}

// QueueInfo is used in Selection6Struct
type QueueInfo struct {
	Name                           string  `json:"name"`
	Size                           string  `json:"size"`
	OldestCustomerRequestTimeInSec float64 `json:"oldestCustomerRequestTimeInSec"`
}

// Selection6Struct is the struct to represent selection 6
type Selection6Struct struct {
	Status string    `json:"status"`
	Queue  QueueInfo `json:"queue"`
}

// ErrorStruct is used to show error messages
type ErrorStruct struct {
	Msg string `json:"message"`
}

// Selection4ErrorStruct is also used to show error messages
type Selection4ErrorStruct struct {
	Error string `json:"error"`
	Msg   string `json:"message"`
}
