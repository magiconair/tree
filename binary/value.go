package binary

import (
	"fmt"
	"strings"
)

// Value defines a value that can be stored in a tree.
type Value interface {
	// Compare returns an integer comparing two values.
	// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
	Compare(v Value) int
}

type StringValue string

func (a StringValue) Compare(other Value) int {
	b, ok := other.(StringValue)
	if !ok {
		panic(fmt.Sprintf("not a string: %v", other))
	}
	return strings.Compare(string(a), string(b))
}
