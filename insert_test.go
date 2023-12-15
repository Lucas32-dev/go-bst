package bst

import "testing"

func TestInsert(t *testing.T) {
	b := New()

	tests := []struct {
		value int
		fails bool
	}{
		{
			value: 2,
			fails: false,
		},
		{
			value: 21,
			fails: false,
		},
		{
			value: 3,
			fails: false,
		},
		{
			value: 21,
			fails: true,
		},
		{
			value: 12,
			fails: false,
		},
	}

	for _, test := range tests {
		err := b.Insert(newValueTest(test.value))
		if err == nil && test.fails {
			t.Errorf("Insert of node value: %d passed when an error was expected", test.value)
			continue
		}

		if err != nil && !test.fails {
			t.Errorf("Insert of node value: %d failed, err: %s", test.value, err)
			continue
		}
	}
}
