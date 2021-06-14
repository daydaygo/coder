package algo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(fib(9))
}

func TestMonotoneStack(t *testing.T) {
	want := []int{4, 2, 4, -1, -1}
	arg := []int{2, 1, 2, 4, 3}
	got := MonotoneStack(arg)
	reflect.DeepEqual(got, want)
}
