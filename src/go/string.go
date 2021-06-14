package main

import (
	"fmt"
)

func String() {
	s := "hello"
	for k, v := range s {
		// k int; v int32
		if v == 'h' {
			fmt.Println(v)
		}
		fmt.Println(k, v)
	}

	// c := 'c' // int32
	// var c2 byte = 'c' // uint8

	// s := `[1, null, 2, 3]`
	// var a []int
	// json.Unmarshal([]byte(s), &a) // a: []int{1, 0, 2, 3}
	// b, _ := json.Marshal(a) // b: [1, 0, 2, 3]
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9' // fmt.Println('0')
}

func repeatByte(k int, b []byte) []byte {
	var ans []byte
	for i := 0; i < k; i++ {
		ans = append(ans, b...)
	}
	return ans
}

func repeatString(k int, s string) string {
	var a string
	for i := 0; i < k; i++ {
		a += s
	}
	return a
}
