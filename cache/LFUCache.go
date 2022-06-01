package cache

import (
	"listCache/dlList"
	"listCache/entry"
)

//最不经常使用算法
//freq 根据频率增删, 就是增加比较麻烦，需要for循环比较freq,O(n)效率低

type LFUCache struct {
	capacity int
	list     *dlList.DoubleLinkedList
	elemMap  map[interface{}]*dlList.Node
}

func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		capacity: capacity,
		list:     dlList.NewDoubleLinkedList(),
		elemMap:  make(map[interface{}]*dlList.Node),
	}
}

func (lfuCache *LFUCache) Get(key interface{}) interface{} {
	if node, ok := lfuCache.elemMap[key]; ok {
		node.IncFreq()
		lfuCache.list.RemoveNode(node)
		lfuCache.list.AddFreqNode(node)
		return node.GetData().(*entry.Entry).GetValue()
	}
	return -1
}

func (lfuCache *LFUCache) Put(key, value interface{}) {
	if lfuCache.capacity == 0 {
		return
	}
	if node, ok := lfuCache.elemMap[key]; ok {
		lfuCache.list.RemoveNode(node)
		node.IncFreq()
		node.SetData(entry.NewEntry(key, value))
		lfuCache.list.AddFreqNode(node)
	} else {
		if lfuCache.list.GetSize() == lfuCache.capacity {
			node = lfuCache.list.RemoveTailNode()
			eval := node.GetData().(*entry.Entry)
			delete(lfuCache.elemMap, eval.GetKey())
		}
		node = dlList.NewNode(entry.NewEntry(key, value))
		lfuCache.list.AddTailNode(node)
		lfuCache.elemMap[key] = node
	}
}

func (lfuCache *LFUCache) Display() {
	lfuCache.list.Display()
}

func (lfuCache *LFUCache) GetSize() int {
	return lfuCache.list.GetSize()
}
