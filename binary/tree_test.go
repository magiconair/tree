package binary

import (
	"math/rand"
	"testing"
)

func TestTree(t *testing.T) {
	tests := []struct {
		desc   string
		values []StringValue
		len    int
		depth  int
	}{
		{
			desc:   "empty",
			values: []StringValue{},
			len:    0,
			depth:  0,
		},
		{
			desc:   "one node",
			values: []StringValue{"a"},
			len:    1,
			depth:  1,
		},
		{
			desc:   "balanced",
			values: []StringValue{"b", "a", "c"},
			len:    3,
			depth:  2,
		},
		{
			desc:   "left-leaning",
			values: []StringValue{"c", "b", "a"},
			len:    3,
			depth:  3,
		},
		{
			desc:   "right-leaning",
			values: []StringValue{"a", "b", "c"},
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
			if got, want := bt.Len(), tt.len; got != want {
				t.Fatalf("got len %d want %d", got, want)
			}
			if got, want := bt.Depth(), tt.depth; got != want {
				t.Fatalf("got depth %d want %d", got, want)
			}
			for _, v := range tt.values {
				if !bt.Contains(v) {
					t.Fatalf("tree should contain %v", v)
				}
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
