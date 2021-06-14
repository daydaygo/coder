package leetcode

import "testing"

func Test_maxChunksToSorted(t *testing.T) {
    type args struct {
        arr []int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name: "t1",
            args: args{
                arr: []int{5,4,3,2,1},
            },
            want: 1,
        },
        {
            name: "t2",
            args: args{
                arr: []int{2,1,3,4,4},
            },
            want: 4,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := maxChunksToSorted(tt.args.arr); got != tt.want {
                t.Errorf("maxChunksToSorted() = %v, want %v", got, tt.want)
            }
        })
    }
}
