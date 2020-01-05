package priorityqueue

import (
	"container/heap"
)

// Copied from https://golang.org/pkg/container/heap/
// and then modified slightly

// An Item is something we manage in a priority queue.
type Item struct {
	Value    interface{} // The value of the item; arbitrary.
	Priority int         // The priority of the item in the queue.
	// The Index is needed by update and is maintained by the heap.Interface methods.
	Index int // The index of the item in the heap.
}

// A MinPriorityQueue implements heap.Interface and holds Items.
type MinPriorityQueue []*Item

func (pq MinPriorityQueue) Len() int { return len(pq) }

func (pq MinPriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq MinPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *MinPriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *MinPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *MinPriorityQueue) update(item *Item, value int, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
