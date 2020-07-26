package binarytree

import (
	"testing"
)

func TestInsert(t *testing.T) {
	tree := NewTree(0, 'a')
	root := tree.Root

	tests := []rune{'b', 'c', 'd', 'e', 'f', 'g'}

	for i, test := range tests {
		tree.Insert(i+1, test)
	}

	if got, want := root.Left.Key, 1; got != want {
		t.Errorf("TestInsert got: %v, want: %v", got, want)
	}

	if got, want := root.Right.Key, 2; got != want {
		t.Errorf("TestInsert got: %v, want: %v", got, want)
	}

	if got, want := root.Left.Left.Key, 3; got != want {
		t.Errorf("TreeInsert got: %v, want: %v", got, want)
	}

	if got, want := root.Left.Right.Key, 4; got != want {
		t.Errorf("TreeInsert got: %v, want: %v", got, want)
	}
}
