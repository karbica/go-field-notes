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

	tree = NewTree(0, 'a')
	tree.Remove(0)
	tree.Insert(1, 'b')
	if got, want := tree.Root.Key, 1; got != want {
		t.Errorf("TestInsert got: %v, want: %v", got, want)
	}
}

func TestRemove(t *testing.T) {
	tree := NewTree(0, 'a')
	root := tree.Root

	tests := []rune{'b', 'c', 'd', 'e', 'f', 'g'}

	for i, test := range tests {
		tree.Insert(i+1, test)
	}

	tree.Remove(2)
	if got, want := root.Right.Key, 6; got != want {
		t.Errorf("TreeRemove got: %v, want: %v", got, want)
	}

	tree.Remove(1)
	if got, want := root.Left.Key, 5; got != want {
		t.Errorf("TreeRemove got: %v, want: %v", got, want)
	}

	tree = NewTree(0, 'a')
	tree.Remove(0)
	if got := tree.Root; got != nil {
		t.Errorf("TreeRemove got: %v, want: %v", got, nil)
	}
}

func TestLevelorder(t *testing.T) {
	tree := NewTree(0, 'a')
	tests := []rune{'b', 'c', 'd', 'e', 'f', 'g'}
	output := []rune{}
	want := append([]rune{tree.Root.Value.(rune)}, tests...)

	for i, test := range tests {
		tree.Insert(i, test)
	}

	fn := func(n *Node) {
		v := n.Value.(rune)
		output = append(output, v)
	}

	tree.Levelorder(tree.Root, fn)

	for i, got := range output {
		if got != want[i] {
			t.Errorf("TestLevelorder got: %v, want: %v", got, want)
		}
	}
}
