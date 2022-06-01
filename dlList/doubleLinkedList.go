package dlList

import "fmt"

type DoubleLinkedList struct {
	head *Node
	tail *Node
	size int
}

//初始化链表
func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

//获取链表长度
func (dl *DoubleLinkedList) GetSize() int {
	return dl.size
}

//获取头节点
func (dl *DoubleLinkedList) GetHead() *Node {
	return dl.head
}

//获取尾节点
func (dl *DoubleLinkedList) GetTail() *Node {
	return dl.tail
}

//头部追加节点
func (dl *DoubleLinkedList) AddHeadNode(node *Node) {
	if dl.GetSize() == 0 { //空链表
		dl.head = node
		dl.tail = node
		node.prev = nil
		node.next = nil
	} else {
		dl.head.SetPrevNode(node)
		node.SetPrevNode(nil)
		node.SetNextNode(dl.head)
		dl.head = node
	}
	dl.size++
}

//从尾部追加节点
func (dl *DoubleLinkedList) AddTailNode(node *Node) {
	if dl.GetSize() == 0 { //空链表
		dl.head = node
		dl.tail = node
		node.prev = nil
		node.next = nil
	} else {
		dl.tail.SetNextNode(node)
		node.SetPrevNode(dl.tail)
		node.SetNextNode(nil)
		dl.tail = node
	}
	dl.size++
}

//根据索引插入节点
func (dl *DoubleLinkedList) Insert(node *Node, index int) {
	if dl.GetSize() == 0 {
		dl.AddHeadNode(node)
	} else {
		//索引节点
		indexNode := dl.GetNode(index)
		node.SetNextNode(indexNode)
		node.SetPrevNode(indexNode.GetPrevNode())
		if indexNode.GetPrevNode() != nil {
			indexNode.GetPrevNode().SetNextNode(node)
		}
		indexNode.SetPrevNode(node)
	}
	dl.size++
}

//删除任意节点
func (dl *DoubleLinkedList) RemoveNode(node *Node) {
	//默认删除尾部节点
	if node == nil || dl.tail == node {
		node = dl.tail
		dl.tail = node.GetPrevNode()
		dl.tail.SetNextNode(nil)
	} else if dl.head == node { //头节点
		dl.head = node.GetNextNode()
		dl.head.SetPrevNode(nil)
	} else {
		node.GetNextNode().SetPrevNode(node.GetPrevNode())
		node.GetPrevNode().SetNextNode(node.GetNextNode())
	}
	dl.size--
}

//弹出头部节点
func (dl *DoubleLinkedList) RemoveHeadNode() *Node {
	if dl.GetSize() == 0 {
		return nil
	}

	node := dl.head
	if dl.head.GetNextNode() != nil { //有下一个节点
		dl.head = dl.head.GetNextNode()
		dl.head.SetPrevNode(nil)
	} else {
		dl.head = nil
		dl.tail = nil
	}
	dl.size--
	return node
}

//弹出尾部节点
func (dl *DoubleLinkedList) RemoveTailNode() *Node {
	if dl.GetSize() == 0 {
		return nil
	}

	node := dl.tail
	if dl.tail.GetPrevNode() != nil { //有上一个节点
		dl.tail = dl.tail.GetPrevNode()
		dl.tail.SetNextNode(nil)
	} else {
		dl.head = nil
		dl.tail = nil
	}
	dl.size--
	return node
}

//打印链表
func (dl *DoubleLinkedList) Display() {
	fmt.Println("DoubleLinkedList size: ", dl.GetSize())
	if dl.GetSize() == 0 {
		return
	}
	hNode := dl.head
	for hNode != nil {
		fmt.Println("data : ", hNode.GetData())
		hNode = hNode.GetNextNode()
	}
}

//通过索引查找节点，如果索引是负数，从尾部查询
func (dl *DoubleLinkedList) GetNode(index int) *Node {
	if index < 0 {
		index = dl.GetSize() + index
		if dl.GetSize() == 0 || index < 0 || index > dl.GetSize() {
			return nil
		}
		node := dl.tail
		for ; index > 0 && node != nil; index-- {
			node = node.GetPrevNode()
		}
		return node
	}
	if dl.GetSize() == 0 || index > dl.GetSize() {
		return nil
	}
	node := dl.head
	for i := 0; i < index && node != nil; i++ {
		node = node.GetNextNode()
	}
	return node
}

//返回指定区间元素
func (dl *DoubleLinkedList) Range(startIndex, stopIndex int) []*Node {
	nodes := make([]*Node, 0)
	if startIndex < 0 {
		startIndex = dl.GetSize() + startIndex
		if startIndex < 0 {
			startIndex = 0
		}
	}

	if stopIndex < 0 {
		stopIndex = dl.GetSize() + stopIndex
		if stopIndex < 0 {
			stopIndex = 0
		}
	}

	//区间个数
	nodesNum := stopIndex - startIndex + 1
	if nodesNum < 0 {
		return nodes
	}

	startNode := dl.GetNode(startIndex)
	for i := 0; i < nodesNum && startNode != nil; i++ {
		nodes = append(nodes, startNode)
		startNode = startNode.GetNextNode()
	}
	return nodes
}

func (dl *DoubleLinkedList) AddFreqNode(node *Node) {
	if dl.GetSize() == 0 {
		dl.AddHeadNode(node)
	} else {
		hNode := dl.head
		if hNode.freq <= node.freq {
			hNode.prev = node
			node.next = hNode
			node.prev = nil
			dl.head = node
		} else {
			for hNode.next != nil && hNode.freq > node.freq {
				hNode = hNode.next
			}
			oldNode := hNode.next
			hNode.next = node
			node.prev = hNode
			node.next = oldNode
			if oldNode == nil {
				dl.tail = node
			}
		}
	}
	dl.size++
}
