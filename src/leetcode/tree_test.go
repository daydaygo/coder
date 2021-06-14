package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTree(t *testing.T) {
	s := "[1,2,3,4,5,6,7]"
	root := NewTree3(s)
	po := []int{0, 0}        // 节点坐标
	m := make(map[int][]int) // key: x val: x相同的节点的值
	var a []int
	traversal(root, &a, &m, po)

	fmt.Println(root, a, m)
}

func traversal(root *TreeNode, a *[]int, m *map[int][]int, po []int) {
	if root == nil {
		return
	}

	x, y := po[0], po[1]

	traversal(root.Left, a, m, []int{x - 1, y - 1})
	(*m)[x] = append((*m)[x], root.Val)
	*a = append(*a, root.Val)
	traversal(root.Right, a, m, []int{x + 1, y - 1})
}

func TestNewTree(t *testing.T) {
	type args struct {
		preOrder []int
		inOrder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "",
			args: args{preOrder: []int{2, 1, 3}, inOrder: []int{1, 2, 3}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTree(tt.args.preOrder, tt.args.inOrder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newTree2(t *testing.T) {
	type args struct {
		a *[]int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// {args: args{a: &[]int{1, 2, 0, 0, 3, 0, 0}}},
		// {args: args{a: &[]int{1, 2, 3}}},
		// {args: args{a: &[]int{}}},
		// {args: args{a: &[]int{1, 0, 2, 3}}},
		{args: args{a: &[]int{5, 4, 7, 3, 0, 2, 0, -1, 0, 9}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTree2(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTree2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTree3(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// {args: args{s: "[1,2,3]"}},
		// {args: args{s: "[]"}},
		// {args: args{s: "[1,null,2,3]"}},
		{args: args{s: "[5,4,7,3,null,2,null,-1,null,9]"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTree3(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTree3() = %v, want %v", got, tt.want)
			}
		})
	}
}
