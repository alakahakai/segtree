// A segment tree implementation in Go
// Supports different types by using generic interface{} for combine function and emptyValue
package segtree

import (
	"errors"
	"math"
)

// SegmentTree struct
//  Arr: Array to store Segment Tree nodes
//  Fn: Combine function for values in segment
//  EmptyValue: Empty value to return when query is out of segment
type SegmentTree struct {
	Arr        []SegmentTreeNode
	Fn         func(a interface{}, b interface{}) interface{}
	EmptyValue interface{}
}

// SegmentTreeNode struct
//  Start: Segment start index
//  End: Segment end index
//  Value: Value that represents the segment
type SegmentTreeNode struct {
	Start int
	End   int
	Value interface{}
}

// Function to create a segment tree.
//  arr: Original data Array
//  fn: Combine function
//  emptyValue: Empty value to return when query is out of range
func NewSegmentTree(arr []interface{}, fn func(a interface{}, b interface{}) interface{}, emptyValue interface{}) (*SegmentTree, error) {
	debug("Creating new segment tree for Array:", arr)
	var size int
	if len(arr) == 0 {
		return nil, errors.New("Source data array cannot be empty!")
	} else if len(arr) == 1 {
		size = 1
	} else {
		depth := int(math.Ceil(math.Log2(float64(len(arr)))))
		size = 2*int(math.Pow(2, float64(depth))) - 1
	}
	stArr := make([]SegmentTreeNode, size)
	st := SegmentTree{
		Arr:        stArr,
		Fn:         fn,
		EmptyValue: emptyValue,
	}
	st.construct(arr, 0, len(arr)-1, 0)
	debug("Segment tree: %+v", st)
	return &st, nil
}

// Internal function to populate the Segment Tree Array
//  arr: Original data array
//  ss: Start index
//  se: End index
//  si: index in the segment tree Array for query use
func (st *SegmentTree) construct(arr []interface{}, ss int, se int, si int) *SegmentTreeNode {
	// Only one element
	if ss == se {
		st.Arr[si] = SegmentTreeNode{
			Start: ss,
			End:   se,
			Value: arr[ss],
		}
		return &st.Arr[si]
	} else {
		mid := (ss + se) / 2
		left := st.construct(arr, ss, mid, si*2+1)
		right := st.construct(arr, mid+1, se, si*2+2)
		var newstart int
		var newend int
		if left.Start < right.Start {
			newstart = left.Start
		} else {
			newstart = right.Start
		}
		if left.End > right.End {
			newend = left.End
		} else {
			newend = right.End
		}
		v := st.Fn(left.Value, right.Value)
		st.Arr[si] = SegmentTreeNode{
			Start: newstart,
			End:   newend,
			Value: v,
		}
		return &st.Arr[si]
	}
}

// Function to query segment tree
//  ss: Start index
//  se: End index
//  si: Start index in the segment tree array
func (st *SegmentTree) GetSegment(ss int, se int, si int) interface{} {
	debug("Query %+v for range: %d to %d", st.Arr[si], ss, se)
	if se < ss || st.Arr[si].Start > se || st.Arr[si].End < ss {
		debug("%+v is ignored", st.Arr[si])
		return st.EmptyValue
	} else if ss <= st.Arr[si].Start && se >= st.Arr[si].End {
		debug("%+v is selected", st.Arr[si])
		return st.Arr[si].Value
	} else {
		debug("%+v is being looked into", st.Arr[si])
		lv := st.GetSegment(ss, se, si*2+1)
		rv := st.GetSegment(ss, se, si*2+2)
		return st.Fn(lv, rv)
	}
}

// Function to update segment tree
//  ss: Start index
//  se: End index
//  fn: Function to update the value
//  si: Start index in the segment tree array
func (st *SegmentTree) UpdateSegment(ss int, se int, fn func(interface{}) interface{}, si int) interface{} {
	debug("Update %+v for range: %d to %d", st.Arr[si], ss, se)
	if se < ss || st.Arr[si].Start > se || st.Arr[si].End < ss {
		debug("%+v is ignored", st.Arr[si])
		return st.Arr[si].Value
	} else if st.Arr[si].Start == st.Arr[si].End {
		debug("%+v is being updated", st.Arr[si])
		v := fn(st.Arr[si].Value)
		st.Arr[si].Value = v
		return v
	} else {
		debug("%+v is selected to be updated, or is being looked into", st.Arr[si])
		lv := st.UpdateSegment(ss, se, fn, si*2+1)
		rv := st.UpdateSegment(ss, se, fn, si*2+2)
		v := st.Fn(lv, rv)
		st.Arr[si].Value = v
		return st.Arr[si].Value
	}
}
