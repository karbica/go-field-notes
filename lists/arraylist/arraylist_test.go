package arraylist

import (
	"testing"
)

func TestLen(t *testing.T) {
	list := New()

	tests := [][]interface{}{
		{0},
		{1},
		{2},
		{3},
		{4},
		{5},
		{6},
		{7},
	}

	for _, test := range tests {
		got := list.Len()
		want := test[0]
		if got != want {
			t.Errorf("TestLen got: %v, want: %v", got, want)
		}
		list.Push(nil)
	}
}

func TestEmpty(t *testing.T) {
	list := New()

	tests := [][]interface{}{
		{true},
		{false},
		{false},
	}

	for _, test := range tests {
		got := list.Empty()
		want := test[0]
		if got != want {
			t.Errorf("TestEmpty got: %v, want: %v", got, want)
		}
		list.Push(nil)
	}
}

func TestGet(t *testing.T) {
	list := New()

	tests := [][]interface{}{
		{"a"},
		{"b"},
		{"c"},
		{"d"},
		{"e"},
		{"f"},
		{"g"},
	}

	for _, test := range tests {
		list.Push(test[0])
	}

	for i, test := range tests {
		got, _ := list.Get(i)
		want := test[0]
		if got != want {
			t.Errorf("TestGet got: %v, want: %v", got, want)
		}
	}

	for range tests {
		_, err := list.Get(len(tests))
		got := err != nil
		want := true
		if got != want {
			t.Errorf("TestGet got: %v, want: %v", got, want)
		}
	}
}

func TestPush(t *testing.T) {
	list := New()

	tests := [][]interface{}{
		{"a"},
		{"b"},
		{"c"},
		{"d"},
		{"e"},
		{"f"},
		{"g"},
	}

	for _, test := range tests {
		list.Push(test[0])
		got := list.elements[len(list.elements)-1]
		want := test[0]
		if got != want {
			t.Errorf("TestPush got: %v, want %v", got, want)
		}
	}
}

func TestPop(t *testing.T) {
	list := New()

	tests := [][]interface{}{
		{"a"},
		{"b"},
		{"c"},
		{"d"},
		{"e"},
		{"f"},
		{"g"},
	}

	for _, test := range tests {
		list.Push(test[0])
	}

	for i := range tests {
		got, _ := list.Pop()
		want := tests[len(tests)-i-1][0]
		if got != want {
			t.Errorf("TestPop got: %v, want: %v", got, want)
		}
	}

	list = New()
	for range tests {
		_, err := list.Pop()
		got := err != nil
		want := true
		if got != want {
			t.Errorf("TestGet got: %v, want: %v", got, want)
		}
	}
}

func TestShift(t *testing.T) {
	list := New()

	tests := [][]interface{}{
		{"a"},
		{"b"},
		{"c"},
		{"d"},
		{"e"},
		{"f"},
		{"g"},
	}

	for _, test := range tests {
		list.Push(test[0])
	}

	for _, test := range tests {
		got, _ := list.Shift()
		want := test[0]
		if got != want {
			t.Errorf("TestShift got: %v, want: %v", got, want)
		}
	}

	list = New()
	for range tests {
		_, err := list.Shift()
		got := err != nil
		want := true
		if got != want {
			t.Errorf("TestGet got: %v, want: %v", got, want)
		}
	}
}

func TestUnshift(t *testing.T) {
	list := New()

	tests := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}

	for _, test := range tests {
		list.Unshift(test)
		got, _ := list.Get(0)
		want := test
		if got != want {
			t.Errorf("TestUnshift got: %v, want: %v", got, want)
		}
	}
}

func TestWithinRange(t *testing.T) {
	list := New()

	tests := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}

	for i, test := range tests {
		list.Push(test)
		if got, want := list.withinRange(i), true; got != want {
			t.Errorf("TestWithinRange got: %v, want: %v", got, want)
		}
	}

	for i, test := range tests {
		list.Push(test)
		if got, want := list.withinRange(i+10), false; got != want {
			t.Errorf("TestWithinRange got: %v, want: %v", got, want)
		}
	}

	for i, test := range tests {
		list.Push(test)
		if got, want := list.withinRange(i+10*-1), false; got != want {
			t.Errorf("TestWithinRange got: %v, want: %v", got, want)
		}
	}
}
