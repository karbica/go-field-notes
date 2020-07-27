package binarytree

import (
	"testing"
)

func TestRoot(t *testing.T) {
	tree := NewTree()

	var want *Node
	if got, want := tree.Root(), want; got != want {
		t.Errorf("TestRoot got: %v, want: %v", got, want)
	}

	tree.Insert(0, 'a')
	if got, want := tree.Root().Key, 0; got != want {
		t.Errorf("TestRoot got: %v, want: %v", got, want)
	}

	tree.Clear()
	tree.Insert(1, 'b')
	if got, want := tree.Root().Key, 1; got != want {
		t.Errorf("TestRoot got: %v, want: %v", got, want)
	}
}

func TestInsert(t *testing.T) {
	tree := NewTree()
	tests := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}

	for i, test := range tests {
		tree.Insert(i, test)
	}

	if got, want := tree.root.Left.Key, 1; got != want {
		t.Errorf("TestInsert got: %v, want: %v", got, want)
	}

	if got, want := tree.root.Right.Key, 2; got != want {
		t.Errorf("TestInsert got: %v, want: %v", got, want)
	}

	if got, want := tree.root.Left.Left.Key, 3; got != want {
		t.Errorf("TreeInsert got: %v, want: %v", got, want)
	}

	if got, want := tree.root.Left.Right.Key, 4; got != want {
		t.Errorf("TreeInsert got: %v, want: %v", got, want)
	}

	tree.Clear()
	tree.Insert(0, 'a')
	tree.Remove(0)
	tree.Insert(1, 'b')
	if got, want := tree.root.Key, 1; got != want {
		t.Errorf("TestInsert got: %v, want: %v", got, want)
	}
}

func TestRemove(t *testing.T) {
	tree := NewTree()
	tests := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}

	for i, test := range tests {
		tree.Insert(i, test)
	}

	tree.Remove(2)
	if got, want := tree.root.Right.Key, 6; got != want {
		t.Errorf("TreeRemove got: %v, want: %v", got, want)
	}

	tree.Remove(1)
	if got, want := tree.root.Left.Key, 5; got != want {
		t.Errorf("TreeRemove got: %v, want: %v", got, want)
	}

	tree.Clear()
	tree.Insert(0, 'a')
	tree.Remove(0)
	if got := tree.root; got != nil {
		t.Errorf("TreeRemove got: %v, want: %v", got, nil)
	}
}

func TestLevelorder(t *testing.T) {
	tree := NewTree()
	tests := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	output := []rune{}

	for i, test := range tests {
		tree.Insert(i, test)
	}

	fn := func(n *Node) {
		v := n.Value.(rune)
		output = append(output, v)
	}

	tree.Levelorder(tree.root, fn)

	for i, got := range output {
		if want := tests[i]; got != want {
			t.Errorf("TestLevelorder got: %v, want: %v", got, want)
		}
	}
}

func TestSize(t *testing.T) {
	tree := NewTree()
	tests := []rune{'a', 'b', 'c'}

	for i, test := range tests {
		if got, want := tree.Size(), i; got != want {
			t.Errorf("TestSize got: %v, want: %v", got, want)
		}
		tree.Insert(i, test)
	}

	for i := range tests {
		if got, want := tree.Size(), len(tests)-i; got != want {
			t.Errorf("TestSize got: %v, want: %v", got, want)
		}
		tree.Remove(i)
	}
}

func TestEmpty(t *testing.T) {
	tree := NewTree()
	tests := []bool{true, false, false}

	for i, test := range tests {
		if got, want := tree.Empty(), test; got != want {
			t.Errorf("TestEmpty got: %v, want: %v", got, want)
		}
		tree.Insert(i, test)
	}

	for i := len(tests) - 1; i > 0; i-- {
		if got, want := tree.Empty(), tests[i]; got != want {
			t.Errorf("TestEmpty got: %v, want: %v", got, want)
		}
		tree.Remove(i)
	}
}

func TestClear(t *testing.T) {
	tree := NewTree()

	tree.Clear()
	if got, want := tree.Empty(), true; got != want {
		t.Errorf("TestClear got: %v, want: %v", got, want)
	}

	tree.Insert(0, 'a')
	tree.Clear()
	if got, want := tree.Empty(), true; got != want {
		t.Errorf("TestClear got: %v, want: %v", got, want)
	}
}
