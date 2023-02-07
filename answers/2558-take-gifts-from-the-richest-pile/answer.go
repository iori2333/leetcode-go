package answer2558

import (
	"container/heap"
	"math"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func pickGifts(gifts []int, k int) (res int64) {
	var h IntHeap
	heap.Init(&h)
	for _, gift := range gifts {
		heap.Push(&h, gift)
	}

	for i := 0; i < k; i++ {
		r := heap.Pop(&h).(int)
		heap.Push(&h, int(math.Sqrt(float64(r))))
	}

	for _, r := range h {
		res += int64(r)
	}
	return
}
