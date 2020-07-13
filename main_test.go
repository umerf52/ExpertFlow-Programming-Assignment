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

func TestOperations(t *testing.T) {
	name := "DefaultQueue"
	desc := "This queue is for demonstration of Priority Queue implementation"
	cap, k, cn := 1, 0, 0
	isIn := false
	pq := &PriorityQueue{
		queueName:        name,
		queueDescription: desc,
		capacity:         cap,
		key:              k,
		count:            cn,
		isInitialized:    isIn}

	cr := &CustomerRequest{
		PriorityWeight: 10,
		CustomerName:   "new",
		Description:    "important work",
		EnqueueTime:    time.Now(),
	}
	_ = insert(pq, cr, false)

	if len(pq.harr) != pq.count || pq.count != 1 {
		t.Errorf("Insertion creation failed. count and heap array length do no match")
	}

	_ = insert(pq, cr, false)

	if pq.count > 1 {
		t.Errorf("Heap array length exceeded")
	}
}
