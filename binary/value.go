package binary

import (
	"fmt"
	"strings"
)

// Value defines a value that can be stored in a tree.
type Value interface {
	// Compare returns an integer comparing two values.
	// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
	Compare(b Value) int
}

type StringValue string

func (a StringValue) Compare(b Value) int {
	bb, ok := b.(StringValue)
	if !ok {
		panic(fmt.Sprintf("not a string: %v", b))
	}
	return strings.Compare(string(a), string(bb))
}
