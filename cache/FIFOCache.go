package cache

import (
	"listCache/dlList"
	"listCache/entry"
)

/*
在FIFO Cache设计中，核⼼原则就是：如果⼀个数据最先进⼊缓存中，则应该最早淘汰掉。也就是说，当缓存满的时候，应当把最先进⼊
缓存的数据给淘汰掉。在FIFO Cache中应该⽀持以下操作;
Get(key)：如果Cache中存在该key，则返回对应的value值，否则，返回-1；
Put(key,value)：如果Cache中存在该key，则重置value值；如果不存在该key，则将该key插⼊到到Cache中，若Cache已满，则淘汰最
早进⼊Cache的数据

算法：页面cache算法之FIFO
特点：最先进入最先淘汰
实现：链表+hashmap，hashmap用来加速查找
插入/查找/删除：O(1)
*/
type FIFOCache struct {
	list     *dlList.DoubleLinkedList
	elemMap  map[interface{}]*dlList.Node
	capacity int //容量
	//缺页中断的次数
	count int
}

func NewFIFOCache(capacity int) *FIFOCache {
	return &FIFOCache{
		capacity: capacity,
		list:     dlList.NewDoubleLinkedList(),
		elemMap:  make(map[interface{}]*dlList.Node),
	}
}

func (ffCache *FIFOCache) GetCount() int {
	return ffCache.count
}

func (ffCache *FIFOCache) Get(key interface{}) interface{} {
	if node, ok := ffCache.elemMap[key]; ok {
		return node.GetData().(*entry.Entry).GetValue()
	}
	return -1
}

func (ffCache *FIFOCache) Put(key, value interface{}) {
	if ffCache.capacity == 0 {
		return
	}
	if node, ok := ffCache.elemMap[key]; ok { //存在，重置
		ffCache.list.RemoveNode(node)
		node.SetData(entry.NewEntry(key, value))
		ffCache.list.AddTailNode(node)
	} else {
		if ffCache.list.GetSize() == ffCache.capacity { //Cache已满，则淘汰最早进⼊Cache的数据
			node = ffCache.list.RemoveHeadNode()
			eval := node.GetData().(*entry.Entry)
			delete(ffCache.elemMap, eval.GetKey())
		}
		node = dlList.NewNode(entry.NewEntry(key, value))
		ffCache.elemMap[key] = node
		ffCache.list.AddTailNode(node)
	}
}

func (ffCache *FIFOCache) Display() {
	ffCache.list.Display()
}

func (ffCache *FIFOCache) GetSize() int {
	return ffCache.list.GetSize()
}
