package trees_and_graphs

import (
	"reflect"
	"testing"
)

func TestNewBST(t *testing.T) {
	tests := []struct {
		input []int
		want  *BTNode
	}{
		{input: []int{}, want: nil},
		{input: []int{1}, want: &BTNode{data: 1}},
		{input: []int{1, 2}, want: &BTNode{data: 1, right: &BTNode{data: 2}}},
		{input: []int{1, 2, 3}, want: &BTNode{data: 2, left: &BTNode{data: 1}, right: &BTNode{data: 3}}},
		{
			input: []int{3, 5, 2, 7, 1},
			want:  &BTNode{data: 3, left: &BTNode{data: 1, right: &BTNode{data: 2}}, right: &BTNode{data: 5, right: &BTNode{data: 7}}},
		},
	}
	for _, tt := range tests {
		if got := NewBST(tt.input); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("NewBST() = %v, want %v", got.String(), tt.want.String())
		}
	}
}

func TestBTNode_IsBalanced(t *testing.T) {
	tests := []struct {
		input *BTNode
		want  bool
	}{
		{input: nil, want: true},
		{input: &BTNode{data: 1}, want: true},
		{input: &BTNode{data: 1, right: &BTNode{data: 2}}, want: true},
		{input: &BTNode{data: 2, right: &BTNode{data: 3, right: &BTNode{data: 1}}}, want: false},
	}
	for _, tt := range tests {
		if got := tt.input.IsBalanced(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("input: %v, input.IsBalanced() = %v, want %v", tt.input.String(), got, tt.want)
		}
	}
}
