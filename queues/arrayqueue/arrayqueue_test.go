package arrayqueue

import (
	"testing"
)

func TestLen(t *testing.T) {
	queue := New()

	tests := [][]interface{}{
		{"a", 1},
		{"b", 2},
		{"c", 3},
		{"d", 4},
		{"e", 5},
		{"f", 6},
		{"g", 7},
	}

	if queue.Len() != 0 {
		t.Errorf("TestLen got: %v, want: %v", queue.Len(), 0)
	}

	for _, test := range tests {
		queue.Enqueue(test[0])
		got := queue.Len()
		want := test[1]
		if got != want {
			t.Errorf("TestLen got: %v, want: %v", got, want)
		}
	}
}

func TestEmpty(t *testing.T) {
	queue := New()

	tests := [][]interface{}{
		{"a", false},
		{"b", false},
		{"c", false},
	}

	if queue.Empty() != true {
		t.Errorf("TestEmpty received: %v, expected: %v", queue.Empty(), true)
	}

	for _, test := range tests {
		queue.Enqueue(test[0])
		if queue.Empty() != test[1] {
			t.Errorf("TestEmpty received: %v, expected: %v", queue.Empty(), test[1])
		}
	}

	queue.Dequeue()

	if queue.Empty() != false {
		t.Errorf("TestEmpty received: %v, expected: %v", queue.Empty(), false)
	}
}

func TestPeek(t *testing.T) {
	queue := New()

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
		queue.Enqueue(test[0])
	}

	for _, test := range tests {
		if received, _ := queue.Peek(); received != test[0] {
			t.Errorf("TestPeek received: %v, expected: %v", received, test[0])
		}
		queue.Dequeue()
	}
}

func TestEnqueue(t *testing.T) {
	queue := New()

	tests := [][]interface{}{
		{"a"},
		{"b"},
		{"c"},
		{"d"},
		{"e"},
		{"f"},
		{"g"},
	}

	for i, test := range tests {
		queue.Enqueue(test[0])
		got, _ := queue.items.Get(i)
		want := test[0]
		if got != want {
			t.Errorf("TestEnqueue got: %v, want: %v", got, want)
		}
	}
}

func TestDequeue(t *testing.T) {
	queue := New()

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
		queue.Enqueue(test[0])
	}

	for _, test := range tests {
		if received, _ := queue.Dequeue(); received != test[0] {
			t.Errorf("TestDequeue received: %v, expected: %v", received, test[0])
		}
	}
}
