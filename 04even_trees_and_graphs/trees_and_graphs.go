package trees_and_graphs

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type TreeNode struct {
	data     int
	children []*TreeNode
}

type Tree struct {
	root *TreeNode
}

type BTNode struct {
	data   int
	parent *BTNode
	left   *BTNode
	right  *BTNode
}

func (bt *BTNode) String() string {
	if bt.left == nil && bt.right == nil {
		return strconv.Itoa(bt.data)
	}
	var left, right string
	if bt.left == nil {
		left = "nil"
	} else {
		left = bt.left.String()
	}
	if bt.right == nil {
		right = "nil"
	} else {
		right = bt.right.String()
	}
	return fmt.Sprintf("{ %d %s %s }", bt.data, left, right)
}

func (bt *BTNode) AddChild(left, right int) {
	bt.left = &BTNode{
		data:   left,
		parent: bt,
	}
	bt.right = &BTNode{
		data:   right,
		parent: bt,
	}
}

// 4.2 Minimal Tree: Given a sorted (increasing order) array with unique integer elements, write an
//algorithm to create a binary search tree with minimal height.

func NewBST(arr []int) *BTNode {
	sort.Ints(arr)
	return newBST(arr, 0, len(arr)-1)
}

func newBST(sortedArr []int, start int, end int) *BTNode {
	if end < start {
		return nil
	}
	mid := (start + end) / 2
	return &BTNode{
		data:  sortedArr[mid],
		left:  newBST(sortedArr, start, mid-1),
		right: newBST(sortedArr, mid+1, end),
	}
}

//4.4 Check Balanced: Implement a function to check if a binary tree is balanced. For the purposes of
//this question, a balanced tree is defined to be a tree such that the heights of the two subtrees of any
//node never differ by more than one

func (bt *BTNode) checkBalancedHeight() int {
	if bt == nil {
		return -1
	}
	lHeight := bt.left.checkBalancedHeight()
	if lHeight == math.MinInt32 {
		return math.MinInt32
	}
	rHeight := bt.right.checkBalancedHeight()
	if rHeight == math.MinInt32 {
		return math.MinInt32
	}
	if lHeight-rHeight > 1 || lHeight-rHeight < -1 {
		return math.MinInt32
	}

	if lHeight > rHeight {
		return lHeight + 1
	}
	return rHeight + 1
}

func (bt *BTNode) IsBalanced() bool {
	return bt.checkBalancedHeight() != math.MinInt32
}

//4.6 Successor: Write an algorithm to find the "next" node (i.e., in-order successor) of a given node in a
//binary search tree. You may assume that each node has a link to its parent.

func (bt *BTNode) leftMostChild() *BTNode {
	if bt == nil {
		return nil
	}
	n := bt
	for n.left != nil {
		n = n.left
	}
	return n
}

func (bt *BTNode) InOrderSucc() *BTNode {
	if bt == nil {
		return nil
	}

	if bt.right != nil {
		return bt.right.leftMostChild()
	}
	q := bt
	x := q.parent
	for x != nil && x.left != q {
		q = q.parent
		x = x.parent
	}
	return x
}

//4.8 First Common Ancestor: Design an algorithm and write code to find the first common ancestor
//of two nodes in a binary tree. Avoid storing additional nodes in a data structure. NOTE: This is not
//necessarily a binary search tree.

//4.1O Check Subtree: T1 and T2 are two very large binary trees, with T1 much bigger than T2. Create an
//algorithm to determine ifT2 is a subtree of T1.
//A tree T2 is a subtree of T1 if there exists a node n in Tl such that the subtree of n is identical to T2.
//That is, if you cut off the tree at node n, the two trees would be identical.

func (bt *BTNode) Contains(t *BTNode) bool {
	s1 := bytes.NewBufferString("")
	s2 := bytes.NewBufferString("")
	bt.getOrderString(s1)
	t.getOrderString(s2)
	return strings.Contains(s1.String(), s2.String())
}

func (bt *BTNode) getOrderString(s *bytes.Buffer) {
	if bt == nil {
		s.WriteString("X")
		return
	}
	//pre-order
	s.WriteString(fmt.Sprintf("%d ", bt.data))
	bt.left.getOrderString(s)
	bt.right.getOrderString(s)
}

//T1のノードを探索していって、T2のrootと同じデータを見つけたらT2と同じか確かめるという手もある
//こっちのほうが効率は良い

//4.12 Paths with Sum: You are given a binary tree in which each node contains an integer value (which
//might be positive or negative). Design an algorithm to count the number of paths that sum to a
//given value. The path does not need to start or end at the root or a leaf, but it must go downwards
//(traveling only from parent nodes to child nodes).

//runningSum: pathの累積和
//btで終わるパスとbt.left, bt.right以下で終わるパスを足す

func (bt *BTNode) CountPathWithSum(targetSum int) int {
	return bt.countPathWithSum(targetSum, 0, &map[int]int{})
}

func (bt *BTNode) countPathWithSum(targetSum int, runningSum int, pathCount *map[int]int) int {
	if bt == nil {
		return 0
	}

	runningSum += bt.data
	sum := runningSum - targetSum
	totalPaths := (*pathCount)[sum]
	if runningSum == targetSum {
		totalPaths++
	}

	(*pathCount)[runningSum] += 1
	totalPaths += bt.left.countPathWithSum(targetSum, runningSum, pathCount)
	totalPaths += bt.right.countPathWithSum(targetSum, runningSum, pathCount)
	(*pathCount)[runningSum] -= 1

	return totalPaths
}
