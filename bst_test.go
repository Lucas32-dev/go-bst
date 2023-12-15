package bst

import "testing"

func TestNew(t *testing.T) {
	b := New()

	len := b.Length()
	if len != 0 {
		t.Errorf("New empty bst is not empty, len: %d", len)
	}
}

func TestLength(t *testing.T) {
	b := New()

	tests := []struct {
		len   int
		manip func()
	}{
		{
			len:   0,
			manip: func() {},
		},
		{
			len: 1,
			manip: func() {
				b.Insert(newValueTest(9))
			},
		},
		{
			len: 3,
			manip: func() {
				b.Insert(newValueTest(9)) // duplicated value
				b.Insert(newValueTest(2))
				b.Insert(newValueTest(8))
			},
		},
		{
			len: 2,
			manip: func() {
				b.Delete(9)
			},
		},
		{
			len: 4,
			manip: func() {
				b.Insert(newValueTest(3))
				b.Insert(newValueTest(4))
			},
		},
	}

	for _, test := range tests {
		test.manip()
		len := b.Length()

		if len != test.len {
			t.Errorf("bst returned wrong length, got: %d expected: %d", len, test.len)
		}
	}
}
