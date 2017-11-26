package binary

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestTree(t *testing.T) {
	stringValues := func(vals ...string) []StringValue {
		var x []StringValue
		for _, v := range vals {
			x = append(x, StringValue(v))
		}
		return x
	}
	values := func(vals ...string) []Value {
		var x []Value
		for _, v := range vals {
			x = append(x, StringValue(v))
		}
		return x
	}

	tests := []struct {
		desc          string
		values        []StringValue
		pre, in, post []Value
		len           int
		depth         int
	}{
		{
			desc:  "empty",
			len:   0,
			depth: 0,
		},
		{
			desc:   "one node",
			values: stringValues("a"),
			pre:    values("a"),
			in:     values("a"),
			post:   values("a"),
			len:    1,
			depth:  1,
		},
		{
			//     b
			//    / \
			//   a   c
			desc:   "balanced",
			values: stringValues("b", "a", "c"),
			pre:    values("b", "a", "c"),
			in:     values("a", "b", "c"),
			post:   values("a", "c", "b"),
			len:    3,
			depth:  2,
		},
		{
			//     c
			//    /
			//   b
			//  /
			// a
			desc:   "left-leaning",
			values: []StringValue{"c", "b", "a"},
			pre:    values("c", "b", "a"),
			in:     values("a", "b", "c"),
			post:   values("a", "b", "c"),
			len:    3,
			depth:  3,
		},
		{
			//     a
			//      \
			//       b
			//        \
			//         c
			desc:   "right-leaning",
			values: []StringValue{"a", "b", "c"},
			pre:    values("a", "b", "c"),
			in:     values("a", "b", "c"),
			post:   values("c", "b", "a"),
			len:    3,
			depth:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			bt := &Tree{}
			for _, v := range tt.values {
				bt.Add(v)
			}
			for _, v := range tt.values {
				if !bt.Contains(v) {
					t.Fatalf("tree should contain %v", v)
				}
			}
			if got, want := bt.Len(), tt.len; got != want {
				t.Errorf("got len %d want %d", got, want)
			}
			if got, want := bt.Depth(), tt.depth; got != want {
				t.Errorf("got depth %d want %d", got, want)
			}
			if got, want := bt.PreOrder(), tt.pre; len(got) > 0 && len(want) > 0 && !reflect.DeepEqual(got, want) {
				t.Errorf("got pre-order %v want %v", got, want)
			}
			if got, want := bt.InOrder(), tt.in; len(got) > 0 && len(want) > 0 && !reflect.DeepEqual(got, want) {
				t.Errorf("got in-order %v want %v", got, want)
			}
			if got, want := bt.PostOrder(), tt.post; len(got) > 0 && len(want) > 0 && !reflect.DeepEqual(got, want) {
				t.Errorf("got post-order %v want %v", got, want)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	const n = 1000000
	vals := randomStrings(n, 8)
	b.ResetTimer()
	b.Run("map", func(b *testing.B) {
		m := map[string]bool{}
		for i := 0; i < b.N; i++ {
			m[vals[i%n]] = true
		}
	})
	b.Run("Tree", func(b *testing.B) {
		m := &Tree{}
		for i := 0; i < b.N; i++ {
			m.Add(StringValue(vals[i%n]))
		}
	})
}

var found int

func BenchmarkFind(b *testing.B) {
	const n = 1000000
	vals := randomStrings(n, 8)
	b.Run("map", func(b *testing.B) {
		m := map[string]bool{}
		for i := 0; i < n; i++ {
			m[vals[i%n]] = true
		}
		b.ResetTimer()
		for i := 0; i < n; i++ {
			if _, ok := m[vals[i]]; ok {
				found++
			}
		}
	})
	b.Run("Tree", func(b *testing.B) {
		m := &Tree{}
		for i := 0; i < n; i++ {
			m.Add(StringValue(vals[i%n]))
		}
		b.ResetTimer()
		for i := 0; i < n; i++ {
			v := StringValue(vals[i])
			if m.Contains(v) {
				found++
			}
		}
	})
}

var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randomStrings(count, length int) []string {
	fill := func(s []rune) {
		for i := 0; i < length; i++ {
			idx := rand.Int31n(int32(len(alphabet)))
			s[i] = alphabet[idx]
		}
	}

	s := make([]rune, length)
	strings := make([]string, count)
	for i := 0; i < count; i++ {
		fill(s)
		strings[i] = string(s)
	}
	return strings
}
