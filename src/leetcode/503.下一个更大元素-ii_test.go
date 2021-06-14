package leetcode

import (
    "reflect"
    "testing"
)

func Test_nextGreaterElements(t *testing.T) {
    type args struct {
        nums []int
    }
    tests := []struct {
        name string
        args args
        want []int
    }{
        {
            name: "t1",
            args: args{
                nums: []int{1,2,1},
            },
            want: []int{2,-1,2},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := nextGreaterElements(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("nextGreaterElements() = %v, want %v", got, tt.want)
            }
        })
    }
}
