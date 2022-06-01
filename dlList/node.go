package dlList

type Node struct {
	data interface{}
	freq int
	prev *Node
	next *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		data: data,
		prev: nil,
		next: nil,
		freq: 0,
	}
}

func (nd *Node) GetPrevNode() *Node {
	return nd.prev
}

func (nd *Node) SetPrevNode(node *Node) {
	nd.prev = node
}

func (nd *Node) GetNextNode() *Node {
	return nd.next
}

func (nd *Node) SetNextNode(node *Node) {
	nd.next = node
}

func (nd *Node) GetData() interface{} {
	return nd.data
}

func (nd *Node) SetData(data interface{}) {
	nd.data = data
}

func (nd *Node) IncFreq() {
	nd.freq++
}
