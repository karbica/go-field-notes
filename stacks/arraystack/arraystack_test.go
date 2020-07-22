package arraystack

import (
	"testing"
)

func TestLen(t *testing.T) {
	stack := New()

	tests := []int{0, 1, 2, 3, 4, 5, 6}

	for _, test := range tests {
		got := stack.Len()
		want := test
		if got != want {
			t.Errorf("TestLen got: %v, want: %v", got, want)
		}
		stack.Push(nil)
	}

	tests = []int{6, 5, 4, 3, 2, 1, 0}

	for _, test := range tests {
		stack.Pop()
		got := stack.Len()
		want := test
		if got != want {
			t.Errorf("TestLen got: %v, want: %v", got, want)
		}
	}
}

func TestEmpty(t *testing.T) {
	stack := New()

	tests := []bool{true, false, false}

	for _, test := range tests {
		got := stack.Empty()
		want := test
		if got != want {
			t.Errorf("TestEmpty got: %v, want %v", got, want)
		}
		stack.Push(nil)
	}

	tests = []bool{false, false, true}

	for _, test := range tests {
		stack.Pop()
		got := stack.Empty()
		want := test
		if got != want {
			t.Errorf("TestEmpty got: %v, want: %v", got, want)
		}
	}
}

func TestPeek(t *testing.T) {
	stack := New()

	tests := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}

	for _, test := range tests {
		stack.Push(test)
		got, _ := stack.Peek()
		want := test
		if got != want {
			t.Errorf("TestPeek got: %v, want: %v", got, want)
		}
	}
}

func TestPush(t *testing.T) {
	stack := New()

	tests := []rune{'a', 'b', 'c', 'e', 'f', 'g'}

	for _, test := range tests {
		stack.Push(test)
	}

	for i, test := range tests {
		got, _ := stack.items.Get(i)
		want := test
		if got != want {
			t.Errorf("TestPush got: %v, want: %v", got, want)
		}
	}
}

func TestPop(t *testing.T) {
	stack := New()

	tests := []rune{'a', 'b', 'c', 'e', 'f', 'g'}

	for _, test := range tests {
		stack.Push(test)
	}

	for i := range tests {
		got, _ := stack.Pop()
		want := tests[len(tests)-i-1]
		if got != want {
			t.Errorf("TestPop got: %v, want: %v", got, want)
		}
	}

	for range tests {
		_, err := stack.Pop()
		got := err == nil
		want := false
		if got != want {
			t.Errorf("TestPop got: %v, want: %v", got, want)
		}
	}
}
