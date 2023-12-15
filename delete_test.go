package bst

import (
	"math/rand"
	"testing"
)

func TestDelete(t *testing.T) {
	b := New()

	b.Insert(newValueTest(23))
	b.Insert(newValueTest(12))
	b.Insert(newValueTest(7))
	b.Insert(newValueTest(9))
	b.Insert(newValueTest(30))

	tests := []struct {
		value int
		ok    bool
	}{
		{
			value: 23,
			ok:    true,
		},
		{
			value: 23, // already deleted
			ok:    false,
		},
		{
			value: 24,
			ok:    false,
		},
		{
			value: 7,
			ok:    true,
		},
		{
			value: 9,
			ok:    true,
		},
		{
			value: 7, // already deleted
			ok:    false,
		},
		{
			value: 30,
			ok:    true,
		},
	}

	for _, test := range tests {
		ok := b.Delete(test.value)
		if ok != test.ok {
			t.Errorf("bst delete failed for node value: %d, got: %v, expected: %v ", test.value, ok, test.ok)
		}
	}

	b.PrintValues()
}

func BenchmarkDelete(b *testing.B) {
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
		tree.Delete(values[i])
	}
}
