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

func (n *Node) Slice() *[]int {
	s := []int{}
	currentNode := n
	for currentNode != nil {
		s = append(s, currentNode.data)
		currentNode = currentNode.next
	}
	return &s
}

//2.1 Remove Dups: Write code to remove duplicates from an unsorted linked list.
//FOLLOW UP How would you solve this problem if a temporary buffer is not allowed?

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

//2.2 Return Kth to Last:
//Implement an algorithm to find the kth to last element of a singly linked list.

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

//LastRec This approach uses O(n) times and O(n) spaces. Moreover, not readable :(
func (n *Node) LastRec(k int) *Node {
	i := 0
	return n.lastRec(k, &i)
}

func (n *Node) lastRec(k int, i *int) *Node {
	if n == nil {
		return nil
	}
	nd := n.next.lastRec(k, i)
	*i += 1
	if *i == k {
		return n
	}
	return nd
}

//LastIter is the most optimal. O(n) time, O(1) space.
//We don't need a length counter.
func (n *Node) LastIter(k int) *Node {
	runner := n
	for i := 0; i < k; i++ {
		if runner == nil {
			return nil
		}
		runner = runner.next
	}
	current := n
	for runner.next != nil {
		current = current.next
		runner = runner.next
	}
	return current
}

//2.3 Delete Middle Node: Implement an algorithm to delete a node in the middle
//(i.e., any node but the first and last node, not necessarily the exact middle)
//of a singly linked list, given only access to that node.
//EXAMPLE input: the node c from the linked list a->b->c->d->e->f
//	Result: nothing is returned, but the new linked list looks like a ->b->d->e->f

//↓なんか解答はこういうことじゃなかった

func (n *Node) DeleteMiddle() {
	runner := n
	current := n
	for runner != nil && runner.next != nil {
		runner = runner.next.next
		current = current.next
	}
	if current != nil && current.next != nil {
		current.next = current.next.next
	}
}
