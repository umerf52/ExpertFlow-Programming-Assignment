// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A Queue implements heap.Interface and holds Items.
type Queue []*Item

// PriorityQueue wraps the actual priority queue and provides additional functionality
type PriorityQueue struct {
	harr                        Queue
	queueName, queueDescription string
}

func (q Queue) Len() int { return len(q) }

func (q Queue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return q[i].priority > q[j].priority
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

// Push : Implementation of Heap's Push()
func (q *Queue) Push(x interface{}) {
	n := len(*q)
	item := x.(*Item)
	item.index = n
	*q = append(*q, item)
}

// Pop : Implementation of Heap's Pop()
func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*q = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (q *Queue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(q, item.index)
}

// This example creates a Queue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func main() {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	q := make(Queue, len(items))
	i := 0
	for value, priority := range items {
		q[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&q)

	priorityQueue := PriorityQueue{harr: q,
		queueName:        "DefaultQueue",
		queueDescription: "This queue is for demonstration of Priority Queue Implementation"}

	// Insert a new item and then modify its priority.
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&q, item)
	q.update(item, item.value, 5)

	// Take the items out; they arrive in decreasing priority order.
	for q.Len() > 0 {
		item := heap.Pop(&q).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
}
