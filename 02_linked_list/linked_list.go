package linkedlist

type Node struct {
	Next *Node
	Data int
}

func NewNode(data int) *Node {
	return &Node{Data: data}
}

func NewNodeFromSlice(s []int) *Node {
	if len(s) == 0 {
		return nil
	}
	n := NewNode(s[0])
	current := n
	for _, d := range s[1:] {
		current.Next = NewNode(d)
		current = current.Next
	}
	return n
}

func (n *Node) AppendToTail(data int) *Node {
	if n == nil {
		return NewNode(data)
	}
	currentNode := n
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = NewNode(data)
	return n
}

func (n *Node) Delete(data int) *Node {
	if n == nil {
		return n
	}
	if n.Data == data {
		return n.Next
	}
	currentNode := n
	for currentNode.Next != nil {
		if currentNode.Next.Data == data {
			currentNode.Next = currentNode.Next.Next
			return n
		}
		currentNode = currentNode.Next
	}
	return n
}

func (n *Node) Slice() *[]int {
	s := []int{}
	currentNode := n
	for currentNode != nil {
		s = append(s, currentNode.Data)
		currentNode = currentNode.Next
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
	seen[currentNode.Data] = true
	for currentNode.Next != nil {
		if seen[currentNode.Next.Data] {
			currentNode.Next = currentNode.Next.Next
			continue
		}
		seen[currentNode.Next.Data] = true
		currentNode = currentNode.Next
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
		for runner.Next != nil {
			if current.Data == runner.Next.Data {
				runner.Next = runner.Next.Next
				continue
			}
			runner = runner.Next
		}
		current = current.Next
	}
}

//2.2 Return Kth to Last:
//Implement an algorithm to find the kth to last element of a singly linked list.

func (n *Node) Last(k int) *Node {
	runner := n
	length := 0
	for runner != nil {
		length += 1
		runner = runner.Next
	}

	current := n
	for i := 0; i < length-k; i++ {
		current = current.Next
	}

	return current
}

//LastRec This approach uses O(n) time and O(n) space. Moreover, not readable :(
func (n *Node) LastRec(k int) *Node {
	i := 0
	return n.lastRec(k, &i)
}

func (n *Node) lastRec(k int, i *int) *Node {
	if n == nil {
		return nil
	}
	nd := n.Next.lastRec(k, i)
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
		runner = runner.Next
	}
	current := n
	for runner.Next != nil {
		current = current.Next
		runner = runner.Next
	}
	return current.Next
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
	if runner != nil && runner.Next != nil {
		runner = runner.Next.Next
	}
	for runner != nil && runner.Next != nil {
		runner = runner.Next.Next
		current = current.Next
	}
	if current != nil && current.Next != nil {
		current.Next = current.Next.Next
	}
}

//2.4 Partition: Write code to partition a linked list around a value x,
//such that all nodes less than x come before all nodes greater than or equal to x.
//If x is contained within the list the values of x only need to be after
//the elements less than x (see below).
//The partition element x can appear anywhere in the "right partition";
//it does not need to appear between the left and right partitions.
//Input: 3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1 [partition=5] Output: 3 -> 1 -> 2 -> 10 -> 5 -> 5 -> 8

func (n *Node) Partition(x int) (*Node, *Node) {
	if n == nil {
		return nil, nil
	}
	if n.Next == nil {
		return n, n
	}
	head, tail := n.Next.Partition(x)
	if n.Data < x {
		n.Next = head
		return n, tail
	}
	tail.Next = NewNode(n.Data)
	return head, tail.Next
}

func (n *Node) PartitionStable(x int) *Node {
	if n == nil {
		return nil
	}
	var beforeHead, beforeTail, afterHead, afterTail *Node
	current := n
	for current != nil {
		next := current.Next
		current.Next = nil
		if current.Data < x {
			if beforeHead == nil {
				beforeHead = current
				beforeTail = beforeHead
			} else {
				beforeTail.Next = current
				beforeTail = current
			}
		} else {
			if afterHead == nil {
				afterHead = current
				afterTail = afterHead
			} else {
				afterTail.Next = current
				afterTail = current
			}
		}
		current = next
	}

	beforeTail.Next = afterHead
	return beforeHead
}
