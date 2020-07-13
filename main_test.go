package main

import (
	"testing"
	"time"
)

func TestCreation(t *testing.T) {
	name := "DefaultQueue"
	desc := "This queue is for demonstration of Priority Queue implementation"
	cap, k, cn := 10, 0, 0
	isIn := false
	pq := &PriorityQueue{
		queueName:        name,
		queueDescription: desc,
		capacity:         cap,
		key:              k,
		count:            cn,
		isInitialized:    isIn}

	if pq == nil {
		t.Errorf("Priority Queue creation failed")
	}

	if pq.queueName != name {
		t.Errorf("Priority Queue creation failed. queueName does not match")
	}

	if pq.queueDescription != desc {
		t.Errorf("Priority Queue creation failed. queueDescription does not match")
	}

	if pq.capacity != cap {
		t.Errorf("Priority Queue creation failed. capacity does not match")
	}

	if pq.key != k {
		t.Errorf("Priority Queue creation failed. key does not match")
	}

	if pq.count != cn {
		t.Errorf("Priority Queue creation failed. count does not match")
	}

	if pq.isInitialized != isIn {
		t.Errorf("Priority Queue creation failed. isIntialized does not match")
	}
}

func TestInsert(t *testing.T) {
	pq := &PriorityQueue{
		queueName:        "DefaultQueue",
		queueDescription: "This queue is for demonstration of Priority Queue implementation",
		capacity:         1,
		key:              0,
		count:            0,
		isInitialized:    false}

	cr := &CustomerRequest{
		PriorityWeight: 10,
		CustomerName:   "new",
		Description:    "important work",
		EnqueueTime:    time.Now(),
	}
	result := insert(pq, cr, false)

	if len(pq.harr) != pq.count || pq.count != 1 || !result {
		t.Errorf("Insertion failed. count and heap array length do no match")
	}

	result = insert(pq, cr, false)

	if pq.count > 1 || result {
		t.Errorf("Heap array length exceeded")
	}
}

func TestRemoval(t *testing.T) {
	pq := &PriorityQueue{
		queueName:        "DefaultQueue",
		queueDescription: "This queue is for demonstration of Priority Queue implementation",
		capacity:         2,
		key:              0,
		count:            0,
		isInitialized:    false}

	firstName := "first"
	firstDesc := "this is first customer"
	firstPriority := 8

	cr1 := &CustomerRequest{
		PriorityWeight: firstPriority,
		CustomerName:   firstName,
		Description:    firstDesc,
		EnqueueTime:    time.Now(),
	}

	_ = insert(pq, cr1, false)

	secondName := "second"
	secondDesc := "this is second customer"
	secondPriority := 10

	cr2 := &CustomerRequest{
		PriorityWeight: secondPriority,
		CustomerName:   secondName,
		Description:    secondDesc,
		EnqueueTime:    time.Now(),
	}

	_ = insert(pq, cr2, false)

	cr3 := &CustomerRequest{
		PriorityWeight: 9,
		CustomerName:   "third",
		Description:    "this is third customer",
		EnqueueTime:    time.Now(),
	}

	_ = insert(pq, cr3, false)

	extractedMax := extractMax(pq)
	if extractedMax.CustomerName != secondName || extractedMax.Description != secondDesc || extractedMax.PriorityWeight != secondPriority {
		t.Errorf("extractMax() failed. Received unxexpected value")
	}

	deletedCr, _ := deleteByID(pq, 0, false)

	if deletedCr.CustomerName != firstName || deletedCr.Description != firstDesc {
		t.Errorf("deleteCrById() failed. Received unxexpected value")
	}
}
