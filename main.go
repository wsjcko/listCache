package main

import (
	"fmt"
	"listCache/cache"
	"listCache/dlList"
	"listCache/entry"
	"listCache/mcache"
)

func main() {
	// TestDoubleLinkedList()
	TestFIFOCache()
	// TestLRUCache()
	// TestLFUCache()

	//标准库实现
	TestMCacheFIFOCache()
	// TestMCacheLRUCache()
	// TestMCacheLFUCache()
}

func TestDoubleLinkedList() {
	dl := dlList.NewDoubleLinkedList()
	fmt.Println("从开头添加节点")
	for i := 0; i < 5; i++ {
		node := dlList.NewNode(entry.NewEntry(i, i))
		dl.AddHeadNode(node)
	}
	dl.Display()

	fmt.Println("从末尾添加节点")
	for i := 5; i < 10; i++ {
		node := dlList.NewNode(entry.NewEntry(i, i))
		dl.AddTailNode(node)
	}
	dl.Display()

	fmt.Println("删除最后一个节点")

	dl.RemoveTailNode()
	dl.Display()

	fmt.Println("删除第3个节点")
	node := dl.GetNode(3)
	dl.RemoveNode(node)
	dl.Display()

	fmt.Println("第2节点前添加新节点")
	node = dlList.NewNode(33)
	dl.Insert(node, 2)
	dl.Display()

	fmt.Println("第2节点和第5节点区间")
	dlList := dl.Range(2, 5)
	for i := 0; i < len(dlList); i++ {
		fmt.Println("range node : ", i+2, " data : ", dlList[i].GetData())
	}

}

func TestFIFOCache() {
	ffCache := cache.NewFIFOCache(2) //容量2
	ffCache.Put(1, 1)
	ffCache.Display()
	ffCache.Put(2, 2)
	ffCache.Display()

	fmt.Println(ffCache.Get(1))
	ffCache.Display()

	ffCache.Put(3, 3)
	ffCache.Display()

	fmt.Println(ffCache.Get(1))

	ffCache.Put(4, 4)
	ffCache.Display()

	fmt.Println(ffCache.Get(2))

}

func TestMCacheFIFOCache() {
	ffCache := mcache.NewFIFOCache(2) //容量2
	ffCache.Put(1, 1)
	ffCache.Display()
	ffCache.Put(2, 2)
	ffCache.Display()

	fmt.Println(ffCache.Get(1))
	ffCache.Display()

	ffCache.Put(3, 3)
	ffCache.Display()

	fmt.Println(ffCache.Get(1))

	ffCache.Put(4, 4)
	ffCache.Display()

	fmt.Println(ffCache.Get(2))

}

func TestLRUCache() {
	lruCache := cache.NewLRUCache(4) //容量4
	lruCache.Put(1, 1)
	lruCache.Display()
	lruCache.Put(2, 2)
	lruCache.Display()

	fmt.Println(lruCache.Get(1))
	lruCache.Display()

	lruCache.Put(3, 3)
	lruCache.Put(4, 4)
	lruCache.Display()

	lruCache.Put(5, 5)
	lruCache.Display()

	fmt.Println(lruCache.Get(2))

	fmt.Println(lruCache.Get(1))
	lruCache.Display()

}

func TestMCacheLRUCache() {
	lruCache := mcache.NewLRUCache(4) //容量4
	lruCache.Put(1, 1)
	lruCache.Display()
	lruCache.Put(2, 2)
	lruCache.Display()

	fmt.Println(lruCache.Get(1))
	lruCache.Display()

	lruCache.Put(3, 3)
	lruCache.Put(4, 4)
	lruCache.Display()

	lruCache.Put(5, 5)
	lruCache.Display()

	fmt.Println(lruCache.Get(2))

	fmt.Println(lruCache.Get(1))
	lruCache.Display()

}

func TestLFUCache() {
	lfuCache := cache.NewLFUCache(4) //容量4
	lfuCache.Put(1, 1)
	lfuCache.Display()
	lfuCache.Put(2, 2)
	lfuCache.Display()

	fmt.Println(lfuCache.Get(1))
	lfuCache.Display()

	lfuCache.Put(3, 3)
	lfuCache.Put(4, 4)
	lfuCache.Display()

	lfuCache.Put(5, 5)
	lfuCache.Display()

	fmt.Println(lfuCache.Get(2))

	fmt.Println(lfuCache.Get(1))
	lfuCache.Display()

}

func TestMCacheLFUCache() {
	lfuCache := mcache.NewLFUCache(10) //容量4
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key-%d", i)
		lfuCache.Set(key, i)
	}

	for i := 0; i < 3; i++ {
		key := fmt.Sprintf("key-%d", i)
		lfuCache.Set(key, i)
	}

	count := 0
	for count < 10 {
		for i := 5; i < 8; i++ {
			key := fmt.Sprintf("key-%d", i)
			if count == 0 {
				fmt.Println("key: ", key, " value: ", lfuCache.Get(key))
			}
		}
		count++

	}

	lfuCache.Del("key-3")

	//清空输出验证
	lfuCache.Display()
}
