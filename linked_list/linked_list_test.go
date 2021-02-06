package linkedlist

import "testing"

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
		tt.input.AppendToTail(tt.data)
		current := tt.input
		currentExpect := tt.expect
		for current == currentExpect && current != nil {
			if current != currentExpect {
				t.Fatalf("got %+v. expected %+v.",
					tt.input.Slice(),
					tt.expect.Slice(),
				)
			}
			current = current.next
			currentExpect = currentExpect.next
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
		tt.input.Delete(tt.data)
		current := tt.input
		currentExpect := tt.expect
		for current == currentExpect && current != nil {
			if current != currentExpect {
				t.Fatalf("got %+v. expected %+v.",
					tt.input.Slice(),
					tt.expect.Slice(),
				)
			}
			current = current.next
			currentExpect = currentExpect.next
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
		current := tt.input
		currentExpect := tt.expect
		for current == currentExpect && current != nil {
			if current != currentExpect {
				t.Fatalf("got %+v. expected %+v.",
					tt.input.Slice(),
					tt.expect.Slice(),
				)
			}
			current = current.next
			currentExpect = currentExpect.next
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
		current := tt.input
		currentExpect := tt.expect
		for current == currentExpect && current != nil {
			if current != currentExpect {
				t.Fatalf("got %+v. expected %+v.",
					tt.input.Slice(),
					tt.expect.Slice(),
				)
			}
			current = current.next
			currentExpect = currentExpect.next
		}
	}
}
