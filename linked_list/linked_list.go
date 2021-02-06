package linkedlist

type Node struct {
	next *Node
	data int
}

func NewNode(data int) *Node {
	return &Node{data: data}
}

func NewNodeFromSlice(s []int) *Node {
	if len(s) == 0 {
		return nil
	}
	n := NewNode(s[0])
	if len(s) == 1 {
		return n
	}
	current := n
	for d := range s[1:] {
		current.next = NewNode(d)
		current = current.next
	}
	return n
}

func (n *Node) AppendToTail(data int) {
	if n == nil {
		n = NewNode(data)
		return
	}
	currentNode := n
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = NewNode(data)
}

func (n *Node) Delete(data int) {
	if n == nil {
		return
	}
	if n.data == data {
		n = n.next
		return
	}
	currentNode := n
	for currentNode.next != nil {
		if currentNode.next.data == data {
			currentNode.next = currentNode.next.next
			return
		}
		currentNode = currentNode.next
	}
}

//DeleteDups remove duplicates
//O(n)
func (n *Node) DeleteDups() {
	if n == nil {
		return
	}
	seen := map[int]bool{}
	currentNode := n
	seen[currentNode.data] = true
	for currentNode.next != nil {
		if seen[currentNode.next.data] {
			currentNode.next = currentNode.next.next
			continue
		}
		seen[currentNode.next.data] = true
		currentNode = currentNode.next
	}
}

//DeleteDups remove duplicates (No buffer allowed)
//O(n^2)
func (n *Node) DeleteDupsNoBuf() {
	if n == nil {
		return
	}
	current := n
	for current != nil {
		runner := current
		for runner.next != nil {
			if current.data == runner.next.data {
				runner.next = runner.next.next
				continue
			}
			runner = runner.next
		}
		current = current.next
	}
}

func (n *Node) Last(k int) *Node {
	runner := n
	length := 0
	for runner != nil {
		length += 1
		runner = runner.next
	}

	current := n
	for i := 0; i < length-k; {
		if current.next == nil {
			break
		}
		current = current.next
	}

	return current
}

func (n *Node) Slice() *[]int {
	s := []int{}
	currentNode := n
	for currentNode != nil {
		s = append(s, currentNode.data)
		currentNode = currentNode.next
	}
	return &s
}
