package fibs



type LinkedList struct {
	rootNode *LinkedListNode
	endNode *LinkedListNode
	Length int64
}

type LinkedListNode struct {
	nextNode *LinkedListNode
	value    int64
}

func(l *LinkedList) Append(value int64){
	node := &LinkedListNode{
		nextNode: nil,
		value:    value,
	}
	if l.rootNode == nil {
		l.rootNode = node
		l.endNode = node
	}else{
		l.endNode.nextNode = node
		l.endNode = node
	}


	l.Length += 1
}

func(l *LinkedList) PopFirst(){
	if l.rootNode == nil {
		return
	}

	currentRoot := l.rootNode
	l.rootNode = currentRoot.nextNode
	currentRoot.nextNode = nil

	if l.rootNode == nil && l.endNode != nil {
		l.endNode = nil
	}

	l.Length -= 1
}

func CreateLinkedList() *LinkedList{
	return &LinkedList{
		rootNode: nil,
		endNode: nil,
		Length: 0,
	}
}
