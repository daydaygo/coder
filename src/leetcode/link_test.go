package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLink(t *testing.T) {
	a := NewLink2([]int{1, 2, 3, 4})
	b := LinkToArray(a)
	fmt.Println(a, b)

	got := 1
	want := 1
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TestLink() = %v, want %v", got, want)
	}
}

func TestNewLink(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name     string
		args     args
		wantHead *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHead := NewLink(tt.args.a); !reflect.DeepEqual(gotHead, tt.wantHead) {
				t.Errorf("NewLink() = %v, want %v", gotHead, tt.wantHead)
			}
		})
	}
}
