package bst

import (
	"math/rand"
	"testing"
)

type nodeValueTest struct {
	v int
}

func (n nodeValueTest) Value() int {
	return n.v
}

func newValueTest(v int) nodeValueTest {
	return nodeValueTest{v}
}

func TestSearch(t *testing.T) {
	b := New()

	value, ok := b.Search(32)
	if ok || value != nil {
		t.Errorf("Search on empty tree failed, returned node value: %v, ok: %v", value, ok)
	}

	b.Insert(newValueTest(6))
	b.Insert(newValueTest(2))
	b.Insert(newValueTest(3))
	b.Insert(newValueTest(8))
	b.Insert(newValueTest(9))

	tests := []struct {
		value int
		ok    bool
	}{
		{
			value: 6,
			ok:    true,
		},
		{
			value: 9,
			ok:    true,
		},
		{
			value: 1,
			ok:    false,
		},
		{
			value: 12,
			ok:    false,
		},
		{
			value: 3,
			ok:    true,
		},
	}

	for _, test := range tests {
		value, ok := b.Search(test.value)
		if ok != test.ok {
			t.Errorf("Search returned wrong ok status for node value: %d, got: %v, expected: %v", test.value, ok, test.ok)
		}

		if value == nil && ok {
			t.Errorf("Search returned ok but with nil node value")
			continue
		}

		if value != nil && !ok {
			t.Errorf("Search returned not ok but with non nil node value")
			continue
		}

		if value != nil && value.Value() != test.value {
			t.Errorf("Search returned wrong node value, got: %d, expected: %d", value.Value(), test.value)
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	tree := New()

	values := make([]int, b.N)

	var value int
	for i := 0; i < b.N; i++ {
		value = rand.Int()
		values[i] = value

		// ignore errors due to the possibilty
		// of getting duplicated values from
		// rand.Int()
		_ = tree.Insert(newValueTest(value))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree.Search(values[i])
	}
}
