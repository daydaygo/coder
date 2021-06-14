package leetcode

import (
	"reflect"
	"testing"
)

func TestMonotoneStack(t *testing.T) {
	type args struct {
		a []int
	}
	var tests = []struct {
		name string
		args args
		want []int
	}{
		{
			name: "t1",
			args: args{
				a: []int{2, 1, 2, 4, 3},
			},
			want: []int{4, 2, 4, -1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MonotoneStack(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MonotoneStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
