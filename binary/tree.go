package binary

// Tree implements a binary search tree using pointers.
//
// The values need to provide a comparator which implements the sort order.
//
// The implementation uses empty nodes for terminating branches instead of null
// pointers. This simplifies the traversal algorithms at the expense of
// additional recursions. For small trees and high concurrency this may be
// relevant but since this implementation focuses on readability this can be
// optimized later if necessary.
type Tree struct {
	v       Value
	l, r, p *Tree
}

// Creates a new tree from the given values.
// Use pre-order to recreate a saved tree.
func NewTree(values ...Value) *Tree {
	t := &Tree{}
	for _, v := range Values {
		t.Add(v)
	}
	return t
}

// Add adds a value if it does not exist already. It returns true if the value
// was added.
func (t *Tree) Add(v Value) bool {
	if v == nil {
		return false
	}

	if t.empty() {
		t.v = v
		t.p = t
		t.l = &Tree{}
		t.r = &Tree{}
		return true
	}

	switch x := v.Compare(t.v); {
	case x == 0: // duplicate
		return false // duplicate
	case x < 0: // v < t.v
		return t.l.Add(v)
	default: // t.v > v
		return t.r.Add(v)
	}
}

// Del removes a value. It returns true, if the value was removed.
func (t *Tree) Del(v Value) bool {
	if v == nil {
		return false
	}
	panic("not impl")

}

// Contains returns true if the value is in the tree.
func (t *Tree) Contains(v Value) bool {
	if t.empty() {
		return false
	}
	switch x := v.Compare(t.v); {
	case x == 0: // v == t.v
		return true
	case x < 0: // v < t.v
		return t.l.Contains(v)
	default: // t.v > v
		return t.r.Contains(v)
	}
}

// PreOrder returns the values in pre-order (node, left, right).
func (t *Tree) PreOrder() []Value {
	if t.empty() {
		return nil
	}
	var x []Value
	x = append(x, t.v)
	x = append(x, t.l.PreOrder()...)
	x = append(x, t.r.PreOrder()...)
	return x
}

// InOrder returns the values in in-order (left, node, right).
func (t *Tree) InOrder() []Value {
	if t.empty() {
		return nil
	}
	var x []Value
	x = append(x, t.l.InOrder()...)
	x = append(x, t.v)
	x = append(x, t.r.InOrder()...)
	return x
}

// PostOrder returns the values in post-order (left, right, node).
func (t *Tree) PostOrder() []Value {
	if t.empty() {
		return nil
	}
	var x []Value
	x = append(x, t.l.PostOrder()...)
	x = append(x, t.r.PostOrder()...)
	x = append(x, t.v)
	return x
}

// Len returns the number of elements in the tree.
func (t *Tree) Len() int {
	if t.empty() {
		return 0
	}
	return t.l.Len() + t.r.Len() + 1
}

// Depth returns the maximum depth of the tree.
func (t *Tree) Depth() int {
	if t.empty() {
		return 0
	}
	if t.leaf() {
		return 1
	}
	return max(t.l.Depth(), t.r.Depth()) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t *Tree) empty() bool {
	return t.v == nil
}

func (t *Tree) leaf() bool {
	return t.l.empty() && t.r.empty()
}
