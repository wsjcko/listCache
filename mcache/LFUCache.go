package mcache

import (
	"container/heap"
	"fmt"
)

/*
算法：页面cache算法之LFU(Least Frequently Used),优先淘汰最少使用的数据
基于：如果一个数据最近被使用的次数很少，将来被使用的可能性也很小
特点：优先淘汰最少使用的数据
实现：最小堆+hashmap。最小堆按使用次数来排队，hashmap用来加速查找
插入/删除：O(logN)，查找：O(1)
*/

type FreqQueue []*ElemItem

func (fq FreqQueue) Len() int {
	return len(fq)
}

func (fq FreqQueue) Less(i, j int) bool {
	return fq[i].freq < fq[j].freq
}

func (fq FreqQueue) Swap(i, j int) {
	fq[i], fq[j] = fq[j], fq[i]
	fq[i].index = i
	fq[j].index = j
}

func (fq *FreqQueue) Push(x interface{}) {
	length := len(*fq)
	elemItem := x.(*ElemItem)
	elemItem.index = length
	*fq = append(*fq, elemItem)
}

func (fq *FreqQueue) Pop() interface{} {
	fqQueue := *fq
	length := len(fqQueue)
	elemItem := fqQueue[length-1]
	elemItem.index = -1 // for safety
	*fq = fqQueue[0 : length-1]
	return elemItem
}

//freq更新后，要重新排序，时间复杂度为 O(logN)
func (fq *FreqQueue) update(elemItem *ElemItem, value interface{}, freq int) {
	elemItem.value = value
	elemItem.freq = freq
	// (*fq)[elemItem.index] = elemItem
	// 重新排序
	// 分析思路是把 堆(大D) 的树状图画出来，看成一个一个小的堆(小D)，看改变其中一个值，对 大D 有什么影响
	// 可以得出结论，下沉操作和上沉操作分别执行一次能将 queue 排列为堆
	heap.Fix(fq, elemItem.index)
}

type LFUCache struct {
	elemHeap FreqQueue
	elemMap  map[interface{}]*ElemItem
}

func NewLFUCache(capacity int) *LFUCache {
	lfuCache := &LFUCache{
		elemHeap: make([]*ElemItem, 0, capacity),
		elemMap:  make(map[interface{}]*ElemItem),
	}
	heap.Init(&(lfuCache.elemHeap))
	return lfuCache
}

func (lfuCache *LFUCache) Set(key, value interface{}) {
	if cap(lfuCache.elemHeap) == 0 {
		return
	}
	if elemItem, ok := lfuCache.elemMap[key]; ok {
		lfuCache.elemHeap.update(elemItem, value, elemItem.freq+1)
	} else {
		if len(lfuCache.elemHeap) == cap(lfuCache.elemHeap) {
			elemItem := heap.Pop(&(lfuCache.elemHeap)).(*ElemItem)
			delete(lfuCache.elemMap, elemItem.key)
			fmt.Printf("pop %s\n", elemItem.value)
		}

		elemItem = &ElemItem{key: key, value: value, freq: 0}
		heap.Push(&(lfuCache.elemHeap), elemItem)
		heap.Fix(&(lfuCache.elemHeap), elemItem.index)
		lfuCache.elemMap[key] = elemItem
	}
}

func (lfuCache *LFUCache) Get(key interface{}) interface{} {
	elemItem, ok := lfuCache.elemMap[key]
	if ok {
		elemItem.freq++
		heap.Fix(&(lfuCache.elemHeap), elemItem.index)
		return elemItem.value
	}
	return -1
}

func (lfuCache *LFUCache) Del(key interface{}) {
	if elemItem, ok := lfuCache.elemMap[key]; ok {
		heap.Remove(&(lfuCache.elemHeap), elemItem.index)
		delete(lfuCache.elemMap, elemItem.key)
	}
}

func (lfuCache *LFUCache) Display() {
	for lfuCache.elemHeap.Len() > 0 {
		elemItem := heap.Pop(&(lfuCache.elemHeap)).(*ElemItem)
		fmt.Printf("%d:%v\n", elemItem.freq, elemItem.value)
	}
	fmt.Println("End Clear ", len(lfuCache.elemHeap))
}
