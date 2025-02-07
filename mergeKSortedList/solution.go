package mergeksortedlist

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeHeap []*ListNode

// sort Interface implementations

func (h ListNodeHeap) Len() int {
	return len(h)
}

func (h ListNodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h ListNodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {

	lheap := ListNodeHeap{}
	heap.Init(&lheap)
	for _, l := range lists {
		if l == nil {
			continue
		}
		heap.Push(&lheap, l)
	}

	head := &ListNode{}
	current := head

	for lheap.Len() > 0 {
		x := heap.Pop(&lheap)
		listNode := x.(*ListNode)
		current.Next = &ListNode{Val: listNode.Val}
		current = current.Next
		listNode = listNode.Next

		if listNode != nil {
			heap.Push(&lheap, listNode)
		}
	}

	return head.Next

}
