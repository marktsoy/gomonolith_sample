package sender

import (
	"container/heap"
	"sync"
)

// An Item is something we manage in a priority queue
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

var (
	mu sync.Mutex
)

func (pq PriorityQueue) Len() int {
	mu.Lock()
	defer mu.Unlock()
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {

	mu.Lock()
	defer mu.Unlock()
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	mu.Lock()
	defer mu.Unlock()
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	mu.Lock()
	defer mu.Unlock()
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	mu.Lock()
	defer mu.Unlock()
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	mu.Lock()
	defer mu.Unlock()
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}