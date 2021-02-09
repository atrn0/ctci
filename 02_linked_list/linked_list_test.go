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
		n := NewFromSlice(tt.input)
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
			input:  NewFromSlice([]int{}),
			data:   1,
			expect: NewFromSlice([]int{1}),
		},
		{
			input:  NewFromSlice([]int{1}),
			data:   2,
			expect: NewFromSlice([]int{1, 2}),
		},
		{
			input:  NewFromSlice([]int{1, 2, 3, 5, 3, 3, 2, 7, 9}),
			data:   5,
			expect: NewFromSlice([]int{1, 2, 3, 5, 3, 3, 2, 7, 9, 5}),
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
			input:  NewFromSlice([]int{}),
			data:   1,
			expect: nil,
		},
		{
			input:  NewFromSlice([]int{1}),
			data:   1,
			expect: NewFromSlice([]int{}),
		},
		{
			input:  NewFromSlice([]int{1, 2, 3, 5, 3, 3, 2, 7, 9}),
			data:   5,
			expect: NewFromSlice([]int{1, 2, 3, 3, 3, 2, 7, 9}),
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

func TestNode_Length(t *testing.T) {
	tests := []struct {
		input *Node
		want  int
	}{
		{
			input: NewFromSlice([]int{}),
			want:  0,
		},
		{
			input: NewFromSlice([]int{1}),
			want:  1,
		},
		{
			input: NewFromSlice([]int{1, 2, 3, 5, 3, 3, 2, 7, 9}),
			want:  9,
		},
	}
	for _, tt := range tests {
		length := tt.input.Length()
		if !cmp.Equal(length, tt.want) {
			t.Fatalf("got %+v. expected %+v.",
				length,
				tt.want,
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
			input:  NewFromSlice([]int{}),
			expect: nil,
		},
		{
			input:  NewFromSlice([]int{1}),
			expect: NewFromSlice([]int{1}),
		},
		{
			input:  NewFromSlice([]int{1, 2, 3, 5, 3, 3, 2, 7, 9}),
			expect: NewFromSlice([]int{1, 2, 3, 5, 7, 9}),
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
			input:  NewFromSlice([]int{}),
			expect: nil,
		},
		{
			input:  NewFromSlice([]int{1}),
			expect: NewFromSlice([]int{1}),
		},
		{
			input:  NewFromSlice([]int{1, 2, 3, 5, 3, 3, 2, 7, 9}),
			expect: NewFromSlice([]int{1, 2, 3, 5, 7, 9}),
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
			input:  NewFromSlice([]int{}),
			k:      1,
			expect: nil,
		},
		{
			input:  NewFromSlice([]int{1}),
			k:      0,
			expect: NewFromSlice([]int{}),
		},
		{
			input:  NewFromSlice([]int{1, 2, 3, 5, 3, 3, 2, 7, 9}),
			k:      3,
			expect: NewFromSlice([]int{2, 7, 9}),
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
			input:  NewFromSlice([]int{}),
			expect: nil,
		},
		{
			input:  NewFromSlice([]int{1}),
			expect: NewFromSlice([]int{1}),
		},
		{
			input:  NewFromSlice([]int{1, 2}),
			expect: NewFromSlice([]int{1}),
		},
		{
			input:  NewFromSlice([]int{1, 2, 3}),
			expect: NewFromSlice([]int{1, 3}),
		},
		{
			input:  NewFromSlice([]int{1, 2, 3, 5, 3, 1, 2, 7, 9}),
			expect: NewFromSlice([]int{1, 2, 3, 5, 1, 2, 7, 9}),
		},
		{
			input:  NewFromSlice([]int{1, 2, 3, 5, 3, 3, 2, 7}),
			expect: NewFromSlice([]int{1, 2, 3, 5, 3, 2, 7}),
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
			input:  NewFromSlice([]int{}),
			x:      3,
			expect: nil,
		},
		{
			input:  NewFromSlice([]int{1}),
			x:      3,
			expect: NewFromSlice([]int{1}),
		},
		{
			input:  NewFromSlice([]int{2, 1}),
			x:      3,
			expect: NewFromSlice([]int{2, 1}),
		},
		{
			input:  NewFromSlice([]int{3, 2, 1}),
			x:      2,
			expect: NewFromSlice([]int{1, 2, 3}),
		},
		{
			input:  NewFromSlice([]int{1, 2, 3, 5, 3, 1, 2, 7, 9}),
			x:      3,
			expect: NewFromSlice([]int{1, 2, 1, 2, 9, 7, 3, 5, 3}),
		},
		{
			input:  NewFromSlice([]int{10, 2, 9, 9, 9, 5, 3, 3, -1, 6, 2, 7}),
			x:      5,
			expect: NewFromSlice([]int{2, 3, 3, -1, 2, 7, 6, 5, 9, 9, 9, 10}),
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
			input:  NewFromSlice([]int{}),
			x:      3,
			expect: nil,
		},
		{
			input:  NewFromSlice([]int{1}),
			x:      3,
			expect: NewFromSlice([]int{1}),
		},
		{
			input:  NewFromSlice([]int{2, 1}),
			x:      3,
			expect: NewFromSlice([]int{2, 1}),
		},
		{
			input:  NewFromSlice([]int{3, 2, 1}),
			x:      2,
			expect: NewFromSlice([]int{1, 3, 2}),
		},
		{
			input:  NewFromSlice([]int{1, 2, 3, 5, 3, 1, 2, 7, 9}),
			x:      3,
			expect: NewFromSlice([]int{1, 2, 1, 2, 3, 5, 3, 7, 9}),
		},
		{
			input:  NewFromSlice([]int{10, 2, 9, 9, 9, 5, 3, 3, -1, 6, 2, 7}),
			x:      5,
			expect: NewFromSlice([]int{2, 3, 3, -1, 2, 10, 9, 9, 9, 5, 6, 7}),
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

func TestNode_Sum(t *testing.T) {
	tests := []struct {
		left   *Node
		right  *Node
		expect *Node
	}{
		{
			left:   NewFromSlice([]int{0}),
			right:  NewFromSlice([]int{0}),
			expect: NewFromSlice([]int{0}),
		},
		{
			left:   NewFromSlice([]int{1}),
			right:  NewFromSlice([]int{}),
			expect: NewFromSlice([]int{1}),
		},
		{
			left:   NewFromSlice([]int{}),
			right:  NewFromSlice([]int{1, 2}),
			expect: NewFromSlice([]int{1, 2}),
		},
		{
			left:   NewFromSlice([]int{3, 2, 1}),
			right:  NewFromSlice([]int{4, 5, 6}),
			expect: NewFromSlice([]int{7, 7, 7}),
		},
		{
			left:   NewFromSlice([]int{9, 9, 9}),
			right:  NewFromSlice([]int{9, 9, 9}),
			expect: NewFromSlice([]int{8, 9, 9, 1}),
		},
		{
			left:   NewFromSlice([]int{5, 6, 4, 6, 2, 4, 5, 1, 9}),
			right:  NewFromSlice([]int{3, 4, 5, 6, 6, 9, 7, 7, 1, 2, 2, 2, 1, 1, 1}),
			expect: NewFromSlice([]int{8, 0, 0, 3, 9, 3, 3, 9, 0, 3, 2, 2, 1, 1, 1}),
		},
	}
	for _, tt := range tests {
		n := tt.left.Add(tt.right, 0)
		if !cmp.Equal(n, tt.expect) {
			t.Fatalf("got %+v. expected %+v.",
				n.Slice(),
				tt.expect.Slice(),
			)
		}
	}
}

func TestNode_AddForward(t *testing.T) {
	tests := []struct {
		left   *Node
		right  *Node
		expect *Node
	}{
		{
			left:   NewFromSlice([]int{0}),
			right:  NewFromSlice([]int{0}),
			expect: NewFromSlice([]int{0}),
		},
		{
			left:   NewFromSlice([]int{1}),
			right:  NewFromSlice([]int{}),
			expect: NewFromSlice([]int{1}),
		},
		{
			left:   NewFromSlice([]int{}),
			right:  NewFromSlice([]int{1, 2}),
			expect: NewFromSlice([]int{1, 2}),
		},
		{
			left:   NewFromSlice([]int{3, 2, 1}),
			right:  NewFromSlice([]int{4, 5, 6}),
			expect: NewFromSlice([]int{7, 7, 7}),
		},
		{
			left:   NewFromSlice([]int{9, 9, 9}),
			right:  NewFromSlice([]int{9, 9, 9}),
			expect: NewFromSlice([]int{1, 9, 9, 8}),
		},
		{
			left:   NewFromSlice([]int{5, 6, 4, 6, 2, 4, 5, 1, 9}),
			right:  NewFromSlice([]int{3, 4, 5, 6, 6, 9, 7, 7, 1, 2, 2, 2, 1, 1, 1}),
			expect: NewFromSlice([]int{3, 4, 5, 6, 7, 0, 3, 3, 5, 8, 4, 6, 6, 3, 0}),
		},
	}
	for _, tt := range tests {
		n := tt.left.AddForward(tt.right)
		if !cmp.Equal(n, tt.expect) {
			t.Fatalf("got %+v. expected %+v.",
				n.Slice(),
				tt.expect.Slice(),
			)
		}
	}
}

func TestNode_IsPalindrome(t *testing.T) {
	tests := []struct {
		input *Node
		want  bool
	}{
		{
			input: NewFromSlice([]int{}),
			want:  true,
		},
		{
			input: NewFromSlice([]int{1}),
			want:  true,
		},
		{
			input: NewFromSlice([]int{1, 2}),
			want:  false,
		},
		{
			input: NewFromSlice([]int{2, 2}),
			want:  true,
		},
		{
			input: NewFromSlice([]int{1, 2, 1}),
			want:  true,
		},
		{
			input: NewFromSlice([]int{2, 6, 3, 2}),
			want:  false,
		},
		{
			input: NewFromSlice([]int{1, 2, 3}),
			want:  false,
		},
		{
			input: NewFromSlice([]int{2, 2, 4, 4}),
			want:  false,
		},
		{
			input: NewFromSlice([]int{2, 3, 3, 2}),
			want:  true,
		},
		{
			input: NewFromSlice([]int{2, 3, 1, 3, 2}),
			want:  true,
		},
		{
			input: NewFromSlice([]int{5, 4, 4, 3, 4, 3, 4, 4, 5}),
			want:  true,
		},
	}
	for _, tt := range tests {
		b := tt.input.IsPalindrome()
		if b != tt.want {
			t.Fatalf("got %+v for %+v. expected %+v.",
				b,
				tt.input.Slice(),
				tt.want,
			)
		}
	}
}

func TestNode_Intersection(t *testing.T) {
	intersection := NewFromSlice([]int{1, 2, 3})
	tests := []struct {
		input1 *Node
		input2 *Node
		want   *Node
	}{
		{
			input1: nil,
			input2: nil,
			want:   nil,
		},
		{
			input1: New(1),
			input2: New(1),
			want:   nil,
		},
		{
			input1: &Node{
				Next: intersection,
				Data: 9,
			},
			input2: &Node{
				Next: &Node{
					Next: intersection,
					Data: 3,
				},
				Data: 5,
			},
			want: intersection,
		},
	}
	for _, tt := range tests {
		i := tt.input1.Intersection(tt.input2)
		if !cmp.Equal(tt.want, i) {
			t.Fatalf("got %+v. want %+v.",
				i,
				tt.want,
			)
		}
	}
	for _, tt := range tests {
		i := tt.input1.IntersectionNoBuf(tt.input2)
		if !cmp.Equal(tt.want, i) {
			t.Fatalf("got %+v. want %+v.",
				i,
				tt.want,
			)
		}
	}
}
