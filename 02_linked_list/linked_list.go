package linkedlist

type Node struct {
	Next *Node
	Data int
}

func New(data int) *Node {
	return &Node{Data: data}
}

func NewFromSlice(s []int) *Node {
	if len(s) == 0 {
		return nil
	}
	n := New(s[0])
	current := n
	for _, d := range s[1:] {
		current.Next = New(d)
		current = current.Next
	}
	return n
}

func (n *Node) AppendToTail(data int) *Node {
	if n == nil {
		return New(data)
	}
	currentNode := n
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = New(data)
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

func (n *Node) Length() int {
	if n == nil {
		return 0
	}
	return 1 + n.Next.Length()
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
	tail.Next = New(n.Data)
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

//2.5 Sum Lists: You have two numbers represented by a linked list, where each node contains a single digit.
//The digits are stored in reverse order, such that the 1's digit is at the head of the list.
//Write a function that adds the two numbers and returns the sum as a linked list.
//EXAMPLE
//Input: (7-> 1 -> 6) + (5 -> 9 -> 2). That is, 617 + 295. Output:2 -> 1 -> 9. That is, 912.
//FOLLOW UP
//Suppose the digits are stored in forward order. Repeat the above problem.
//Input: (6 -> 1 -> 7) + (2 -> 9 -> 5). That is,617 + 295. Output:9 ->1 ->2. That is, 912.

func (n *Node) Add(right *Node, carry int) *Node {
	s := carry
	if n != nil {
		s += n.Data
	}
	if right != nil {
		s += right.Data
	}
	sum := New(s % 10)
	if n == nil && right == nil {
		if carry == 0 {
			return nil
		}
		return sum
	}
	if n == nil {
		sum.Next = right.Next.Add(nil, s/10)
		return sum
	}
	if right == nil {
		sum.Next = n.Next.Add(nil, s/10)
		return sum
	}
	sum.Next = n.Next.Add(right.Next, s/10)
	return sum
}

func (n *Node) AddForward(right *Node) *Node {
	left := n
	leftLen := left.Length()
	rightLen := right.Length()

	//0 padding
	current := New(0)
	head := current
	if leftLen < rightLen {
		for i := 1; i < rightLen-leftLen; i++ {
			current.Next = New(0)
			current = current.Next
		}
		current.Next = left
		left = head
	} else if leftLen > rightLen {
		for i := 1; i < leftLen-rightLen; i++ {
			current.Next = New(0)
			current = current.Next
		}
		current.Next = right
		right = head
	}

	sum, carry := addForwardSameLen(left, right)
	if carry != 0 {
		s := New(carry)
		s.Next = sum
		return s
	}
	return sum
}

//addForwardSameLen
func addForwardSameLen(left, right *Node) (currentSum *Node, carry int) {
	if left == nil && right == nil {
		return nil, 0
	}

	nextSum, carry := addForwardSameLen(left.Next, right.Next)
	sum := left.Data + right.Data + carry
	currentSum = New(sum % 10)
	currentSum.Next = nextSum
	carry = sum / 10
	return
}

//2.6 Palindrome: Implement a function to check if a linked list is a palindrome.

func (n *Node) IsPalindrome() bool {
	if n == nil {
		return true
	}

	c := &Node{Next: n}
	r := n
	//  -> 1 -> 2 -> 3
	//c
	//     r
	var left *Node
	for {
		if r.Next == nil {
			//length is odd
			break
		}
		if r.Next.Next == nil {
			//length is even
			left = left.Unshift(c.Next.Data)
			break
		}

		c = c.Next
		left = left.Unshift(c.Data)
		r = r.Next.Next
	}
	right := c.Next.Next

	//if length is odd
	//  -> 1 -> 2 -> 3 -> 4 -> 5
	//          c
	//                         r
	// left  = 2 -> 1
	// right = 4 -> 5
	//
	//if length is even
	//  -> 1 -> 2 -> 3 -> 4
	//     c
	//               r
	// left  = 2 -> 1
	// right = 3 -> 4
	for left != nil && right != nil {
		if left.Data != right.Data {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}

func (n *Node) Unshift(data int) *Node {
	return &Node{
		Next: n,
		Data: data,
	}
}

// 2.6 can be implemented using recursion

//2.7 Intersection:
//Given two (singly) linked lists, determine if the two lists intersect.
//Return the intersecting node.
//Note that the intersection is defined based on reference, not value.
//That is, if the kth node of the first linked list is the exact same node
//(by reference) as the jth node of the second linked list, then they are intersecting.

func (n *Node) Intersection(n2 *Node) *Node {
	seen := make(map[*Node]bool, n.Length())
	n1Current := n
	for n1Current != nil {
		seen[n1Current] = true
		n1Current = n1Current.Next
	}

	n2Current := n2
	for n2Current != nil {
		if seen[n2Current] {
			return n2Current
		}
		n2Current = n2Current.Next
	}
	return nil
}

func (n *Node) IntersectionNoBuf(n2 *Node) *Node {
	n1Last, n1Len := n.lastAndLen()
	n2Last, n2Len := n2.lastAndLen()
	if n1Last != n2Last {
		return nil
	}

	n1Current := n
	n2Current := n2
	if n1Len > n2Len {
		n1Current = n1Current.advance(n1Len - n2Len)
	} else {
		n2Current = n2Current.advance(n2Len - n1Len)
	}

	for n1Current != nil {
		if n1Current == n2Current {
			return n1Current
		}
		n1Current = n1Current.Next
		n2Current = n2Current.Next
	}
	return nil
}

func (n *Node) lastAndLen() (*Node, int) {
	if n == nil {
		return n, 0
	}
	c := n
	ln := 1
	for c.Next != nil {
		ln += 1
		c = c.Next
	}
	return c, ln
}

func (n *Node) advance(i int) *Node {
	if n == nil {
		return nil
	}
	c := n
	for j := 0; j < i; j++ {
		c = c.Next
	}
	return c
}
