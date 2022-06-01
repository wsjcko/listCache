package cache

import (
	"listCache/dlList"
	"listCache/entry"
)

//最近最少使用算法，如果存在，就重置到头节点，不存在，头节点加入
type LRUCache struct {
	list     *dlList.DoubleLinkedList
	elemMap  map[interface{}]*dlList.Node
	capacity int //容量
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		list:     dlList.NewDoubleLinkedList(),
		elemMap:  make(map[interface{}]*dlList.Node),
	}
}

func (lruCache *LRUCache) Get(key interface{}) interface{} {
	if node, ok := lruCache.elemMap[key]; ok {
		lruCache.list.RemoveNode(node)
		lruCache.list.AddHeadNode(node)
		return node.GetData().(*entry.Entry).GetValue()
	}
	return -1
}

func (lruCache *LRUCache) Put(key, value interface{}) {
	if lruCache.capacity == 0 {
		return
	}
	if node, ok := lruCache.elemMap[key]; ok {
		lruCache.list.RemoveNode(node)
		node.SetData(entry.NewEntry(key, value))
		lruCache.list.AddHeadNode(node)
	} else {
		if lruCache.list.GetSize() == lruCache.capacity {
			node = lruCache.list.RemoveTailNode()
			eval := node.GetData().(*entry.Entry)
			delete(lruCache.elemMap, eval.GetKey())
		}
		node = dlList.NewNode(entry.NewEntry(key, value))
		lruCache.list.AddHeadNode(node)
		lruCache.elemMap[key] = node
	}
}

func (lruCache *LRUCache) Display() {
	lruCache.list.Display()
}

func (lruCache *LRUCache) GetSize() int {
	return lruCache.list.GetSize()
}
