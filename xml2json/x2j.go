package xml2json

type Attr struct {
	Key   string
	Value string
}

type Node struct {
	Name       string
	Content    string
	Attrs      []Attr
	ChildNodes []Node
}

type StackItem struct {
	Value Node
	Next  *StackItem
}

var size = 0
var Stack = new(StackItem)

func Push(node *Node) bool {
	if Stack == nil {
		Stack = &StackItem{*node, nil}
		size = 1
		return true
	}
	temp := &StackItem{*node, nil}
	temp.Next = Stack
	Stack = temp
	size++
	return true
}

func Pop(t *StackItem) (Node, bool) {
	if size == 0 {
		return Node{}, false
	}
	if size == 1 {
		size = 0
		Stack = nil
		return t.Value, true
	}
	Stack = Stack.Next
	size--
	return t.Value, true
}
