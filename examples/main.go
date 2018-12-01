// Test
package main

import (
	"fmt"
	"github.com/rayqiu/segtree"
)

func main() {
	arr := []interface{}{3, 5, 10, 20, 45, 60}
	f := func(a interface{}, b interface{}) interface{} {
		return a.(int) + b.(int)
	}
	s, _ := segtree.NewSegmentTree(arr, f, 0)
	v := s.GetSegment(2, 5, 0)
	fmt.Println(v)
	plusOne := func(a interface{}) interface{} {
		return a.(int) + 1
	}
	s.UpdateSegment(4, 5, plusOne, 0)
	fmt.Println(s)
	v = s.GetSegment(2, 5, 0)
	fmt.Println(v)

	f = func(a interface{}, b interface{}) interface{} {
		return a.(int) * b.(int)
	}
	s2, _ := segtree.NewSegmentTree(arr, f, 1)
	v = s2.GetSegment(1, 4, 0)
	fmt.Println(v)
	timesFive := func(a interface{}) interface{} {
		return a.(int) * 5
	}
	s2.UpdateSegment(1, 4, timesFive, 0)
	fmt.Println(s2)
	v = s2.GetSegment(1, 4, 0)
	fmt.Println(v)

	arr = []interface{}{"hello", "world", "this", "is", "a", "test"}
	f = func(a interface{}, b interface{}) interface{} {
		return a.(string) + b.(string)
	}
	s3, _ := segtree.NewSegmentTree(arr, f, "")
	v = s3.GetSegment(1, 4, 0)
	fmt.Println(v)
}
