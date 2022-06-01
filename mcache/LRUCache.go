package mcache

import (
	"container/list"
	"fmt"
)

/*
*LRU:Least Recently Used,优先淘汰最少使用的算法
*实现：链表+hashmap。链表用来存放数据，hashmap用来加速查找。
插入时如果链表满则先淘汰再插入头部。链表满时进行淘汰，优先淘汰尾部元素。
查找时用hashmap查找，然后将元素移动到表头
*性能：插入/删除/查找:O(1)
*/

type LRUCache struct {
	capacity int
	list     *list.List
	elemMap  map[interface{}]*list.Element
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		list:     list.New(),
		elemMap:  make(map[interface{}]*list.Element),
	}
}

func (lruCache *LRUCache) Get(key interface{}) interface{} {
	if element, ok := lruCache.elemMap[key]; ok {
		lruCache.list.MoveToFront(element)
		return element.Value.(*ElemItem).value
	}
	return -1
}

func (lruCache *LRUCache) Put(key, value interface{}) {
	if lruCache.capacity == 0 {
		return
	}
	if element, ok := lruCache.elemMap[key]; ok {
		element.Value = &ElemItem{key: key, value: value}
		lruCache.list.MoveToFront(element)
	} else {
		if lruCache.list.Len() == lruCache.capacity {
			element = lruCache.list.Back()
			eval, ok := element.Value.(*ElemItem)
			if ok {
				delete(lruCache.elemMap, eval.key)
				lruCache.list.Remove(element)
			} else {
				fmt.Println("LRUCache.Put Error: element.Value.(*ElemItem) ")
			}

		}
		element = lruCache.list.PushFront(&ElemItem{key: key, value: value})
		lruCache.elemMap[key] = element
	}
}

func (lruCache *LRUCache) Display() {
	for e := lruCache.list.Front(); e != nil; e = e.Next() {
		elemItem, ok := e.Value.(*ElemItem)
		if ok {
			fmt.Printf("%v:%v->", elemItem.key, elemItem.value)
		} else {
			fmt.Println("LRUCache.Display Error: element.Value.(*ElemItem) ")
		}
	}
	fmt.Println()
}

func (lruCache *LRUCache) GetSize() int {
	return lruCache.list.Len()
}
