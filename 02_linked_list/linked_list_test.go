package linkedlist

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestNewNodeFromSlice(t *testing.T) {
	tests := []struct {
		input  []int
		expect *Node
	}{
		{
			input:  []int{},
			expect: nil,
		},
		{
			input: []int{1},
			expect: &Node{
				Data: 1,
			},
		},
		{
			input: []int{1, 2},
			expect: &Node{
				Next: &Node{
					Data: 2,
				},
				Data: 1,
			},
		},
		{
			input: []int{1, 2, 3},
			expect: &Node{
				Next: &Node{
					Next: &Node{
						Next: nil,
						Data: 3,
					},
					Data: 2,
				},
				Data: 1,
			},
		},
	}
	for _, tt := range tests {
		n := NewNodeFromSlice(tt.input)
		if !cmp.Equal(n, tt.expect) {
			t.Fatalf("got %+v. expected %+v.",
				n.Slice(),
				tt.expect.Slice(),
			)
		}
	}
}

func TestNode_AppendToTail(t *testing.T) {
	tests := []struct {
		input  *Node
		data   int
		expect *Node
	}{
		{
			input:  NewNodeFromSlice([]int{}),
			data:   1,
			expect: NewNodeFromSlice([]int{1}),
		},
		{
			input:  NewNodeFromSlice([]int{1}),
			data:   2,
			expect: NewNodeFromSlice([]int{1, 2}),
		},
		{
			input:  NewNodeFromSlice([]int{1, 2, 3, 5, 3, 3, 2, 7, 9}),
			data:   5,
			expect: NewNodeFromSlice([]int{1, 2, 3, 5, 3, 3, 2, 7, 9, 5}),
		},
	}
	for _, tt := range tests {
		n := tt.input.AppendToTail(tt.data)
		if !cmp.Equal(n, tt.expect) {
			t.Fatalf("got %+v. expected %+v.",
				n.Slice(),
				tt.expect.Slice(),
			)
		}
	}
}

func TestNode_Delete(t *testing.T) {
	tests := []struct {
		input  *Node
		data   int
		expect *Node
	}{
		{
			input:  NewNodeFromSlice([]int{}),
			data:   1,
			expect: nil,
		},
		{
			input:  NewNodeFromSlice([]int{1}),
			data:   1,
			expect: NewNodeFromSlice([]int{}),
		},
		{
			input:  NewNodeFromSlice([]int{1, 2, 3, 5, 3, 3, 2, 7, 9}),
			data:   5,
			expect: NewNodeFromSlice([]int{1, 2, 3, 3, 3, 2, 7, 9}),
		},
	}
	for _, tt := range tests {
		n := tt.input
		n = n.Delete(tt.data)
		if !cmp.Equal(n, tt.expect) {
			t.Fatalf("got %+v. expected %+v.",
				n.Slice(),
				tt.expect.Slice(),
			)
		}
	}
}

func TestNode_DeleteDups(t *testing.T) {
	tests := []struct {
		input  *Node
		expect *Node
	}{
		{
			input:  NewNodeFromSlice([]int{}),
			expect: nil,
		},
		{
			input:  NewNodeFromSlice([]int{1}),
			expect: NewNodeFromSlice([]int{1}),
		},
		{
			input:  NewNodeFromSlice([]int{1, 2, 3, 5, 3, 3, 2, 7, 9}),
			expect: NewNodeFromSlice([]int{1, 2, 3, 5, 7, 9}),
		},
	}
	for _, tt := range tests {
		tt.input.DeleteDups()
		if !cmp.Equal(tt.input, tt.expect) {
			t.Fatalf("got %+v. expected %+v.",
				tt.input.Slice(),
				tt.expect.Slice(),
			)
		}
	}
}

func TestNode_DeleteDupsNoBuf(t *testing.T) {
	tests := []struct {
		input  *Node
		expect *Node
	}{
		{
			input:  NewNodeFromSlice([]int{}),
			expect: nil,
		},
		{
			input:  NewNodeFromSlice([]int{1}),
			expect: NewNodeFromSlice([]int{1}),
		},
		{
			input:  NewNodeFromSlice([]int{1, 2, 3, 5, 3, 3, 2, 7, 9}),
			expect: NewNodeFromSlice([]int{1, 2, 3, 5, 7, 9}),
		},
	}
	for _, tt := range tests {
		tt.input.DeleteDupsNoBuf()
		if !cmp.Equal(tt.input, tt.expect) {
			t.Fatalf("got %+v. expected %+v.",
				tt.input.Slice(),
				tt.expect.Slice(),
			)
		}
	}
}

func TestNode_Last(t *testing.T) {
	tests := []struct {
		input  *Node
		k      int
		expect *Node
	}{
		{
			input:  NewNodeFromSlice([]int{}),
			k:      1,
			expect: nil,
		},
		{
			input:  NewNodeFromSlice([]int{1}),
			k:      0,
			expect: NewNodeFromSlice([]int{}),
		},
		{
			input:  NewNodeFromSlice([]int{1, 2, 3, 5, 3, 3, 2, 7, 9}),
			k:      3,
			expect: NewNodeFromSlice([]int{2, 7, 9}),
		},
	}
	for _, tt := range tests {
		n := tt.input.Last(tt.k)
		if !cmp.Equal(n, tt.expect) {
			t.Fatalf("got %+v. expected %+v.",
				n.Slice(),
				tt.expect.Slice(),
			)
		}
	}
	for _, tt := range tests {
		n := tt.input.LastRec(tt.k)
		if !cmp.Equal(n, tt.expect) {
			t.Fatalf("got %+v. expected %+v.",
				n.Slice(),
				tt.expect.Slice(),
			)
		}
	}
	for _, tt := range tests {
		n := tt.input.LastIter(tt.k)
		if !cmp.Equal(n, tt.expect) {
			t.Fatalf("got %+v. expected %+v.",
				n.Slice(),
				tt.expect.Slice(),
			)
		}
	}
}

func TestNode_DeleteMiddle(t *testing.T) {
	tests := []struct {
		input  *Node
		expect *Node
	}{
		{
			input:  NewNodeFromSlice([]int{}),
			expect: nil,
		},
		{
			input:  NewNodeFromSlice([]int{1}),
			expect: NewNodeFromSlice([]int{1}),
		},
		{
			input:  NewNodeFromSlice([]int{1, 2}),
			expect: NewNodeFromSlice([]int{1}),
		},
		{
			input:  NewNodeFromSlice([]int{1, 2, 3}),
			expect: NewNodeFromSlice([]int{1, 3}),
		},
		{
			input:  NewNodeFromSlice([]int{1, 2, 3, 5, 3, 1, 2, 7, 9}),
			expect: NewNodeFromSlice([]int{1, 2, 3, 5, 1, 2, 7, 9}),
		},
		{
			input:  NewNodeFromSlice([]int{1, 2, 3, 5, 3, 3, 2, 7}),
			expect: NewNodeFromSlice([]int{1, 2, 3, 5, 3, 2, 7}),
		},
	}
	for _, tt := range tests {
		tt.input.DeleteMiddle()
		if !cmp.Equal(tt.input, tt.expect) {
			t.Fatalf("got %+v. expected %+v.",
				tt.input.Slice(),
				tt.expect.Slice(),
			)
		}
	}
}

func TestNode_Partition(t *testing.T) {
	tests := []struct {
		input  *Node
		x      int
		expect *Node
	}{
		{
			input:  NewNodeFromSlice([]int{}),
			x:      3,
			expect: nil,
		},
		{
			input:  NewNodeFromSlice([]int{1}),
			x:      3,
			expect: NewNodeFromSlice([]int{1}),
		},
		{
			input:  NewNodeFromSlice([]int{2, 1}),
			x:      3,
			expect: NewNodeFromSlice([]int{2, 1}),
		},
		{
			input:  NewNodeFromSlice([]int{3, 2, 1}),
			x:      2,
			expect: NewNodeFromSlice([]int{1, 2, 3}),
		},
		{
			input:  NewNodeFromSlice([]int{1, 2, 3, 5, 3, 1, 2, 7, 9}),
			x:      3,
			expect: NewNodeFromSlice([]int{1, 2, 1, 2, 9, 7, 3, 5, 3}),
		},
		{
			input:  NewNodeFromSlice([]int{10, 2, 9, 9, 9, 5, 3, 3, -1, 6, 2, 7}),
			x:      5,
			expect: NewNodeFromSlice([]int{2, 3, 3, -1, 2, 7, 6, 5, 9, 9, 9, 10}),
		},
	}
	for _, tt := range tests {
		n, _ := tt.input.Partition(tt.x)
		if !cmp.Equal(n, tt.expect) {
			t.Fatalf("got %+v. expected %+v.",
				n.Slice(),
				tt.expect.Slice(),
			)
		}
	}
}

func TestNode_PartitionStable(t *testing.T) {
	tests := []struct {
		input  *Node
		x      int
		expect *Node
	}{
		{
			input:  NewNodeFromSlice([]int{}),
			x:      3,
			expect: nil,
		},
		{
			input:  NewNodeFromSlice([]int{1}),
			x:      3,
			expect: NewNodeFromSlice([]int{1}),
		},
		{
			input:  NewNodeFromSlice([]int{2, 1}),
			x:      3,
			expect: NewNodeFromSlice([]int{2, 1}),
		},
		{
			input:  NewNodeFromSlice([]int{3, 2, 1}),
			x:      2,
			expect: NewNodeFromSlice([]int{1, 3, 2}),
		},
		{
			input:  NewNodeFromSlice([]int{1, 2, 3, 5, 3, 1, 2, 7, 9}),
			x:      3,
			expect: NewNodeFromSlice([]int{1, 2, 1, 2, 3, 5, 3, 7, 9}),
		},
		{
			input:  NewNodeFromSlice([]int{10, 2, 9, 9, 9, 5, 3, 3, -1, 6, 2, 7}),
			x:      5,
			expect: NewNodeFromSlice([]int{2, 3, 3, -1, 2, 10, 9, 9, 9, 5, 6, 7}),
		},
	}
	for _, tt := range tests {
		n := tt.input.PartitionStable(tt.x)
		if !cmp.Equal(n, tt.expect) {
			t.Fatalf("got %+v. expected %+v.",
				n.Slice(),
				tt.expect.Slice(),
			)
		}
	}
}
