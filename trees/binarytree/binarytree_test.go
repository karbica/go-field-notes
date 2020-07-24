package binarytree

import (
	"testing"
)

func TestInsert(t *testing.T) {
	tree := New(0, 'a')

	tests := []rune{'b', 'c', 'd', 'e', 'f', 'g'}

	for i, test := range tests {
		tree.Insert(i+1, test)
	}

	if got, want := tree.Left.Key, 1; got != want {
		t.Errorf("TestInsert got: %v, want: %v", got, want)
	}

	if got, want := tree.Right.Key, 2; got != want {
		t.Errorf("TestInsert got: %v, want: %v", got, want)
	}

	if got, want := tree.Left.Left.Key, 3; got != want {
		t.Errorf("TreeInsert got: %v, want: %v", got, want)
	}

	if got, want := tree.Left.Right.Key, 4; got != want {
		t.Errorf("TreeInsert got: %v, want: %v", got, want)
	}
}
