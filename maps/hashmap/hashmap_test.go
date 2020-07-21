package hashmap

import (
	"testing"
)

func TestSize(t *testing.T) {
	hmap := New()

	tests := []int{0, 1, 2, 3, 4, 5, 6}

	for _, test := range tests {
		got := hmap.Size()
		want := test
		if got != want {
			t.Errorf("TestSize got: %v, want: %v", got, want)
		}
		hmap.Set(test, test)
	}

	tests = []int{6, 5, 4, 3, 2, 1, 0}

	for _, test := range tests {
		hmap.Delete((test))
		got := hmap.Size()
		want := test
		if got != want {
			t.Errorf("TestSize got: %v, want: %v", got, want)
		}
	}
}

func TestEmpty(t *testing.T) {
	hmap := New()

	tests := []bool{true, false, false}

	for i, test := range tests {
		got := hmap.Empty()
		want := test
		if got != want {
			t.Errorf("TestEmpty got: %v, want: %v", got, want)
		}
		hmap.Set(i, struct{}{})
	}

	tests = []bool{false, false, true}

	for i, test := range tests {
		hmap.Delete(i)
		got := hmap.Empty()
		want := test
		if got != want {
			t.Errorf("TestEmpty got: %v, want: %v", got, want)
		}
	}
}

func TestGet(t *testing.T) {
	hmap := New()

	tests := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}

	for i, test := range tests {
		hmap.Set(i, test)
		got, _ := hmap.Get(i)
		want := test
		if got != want {
			t.Errorf("TestGet got: %v, want: %v", got, want)
		}
	}

	for range tests {
		_, got := hmap.Get('z')
		want := false
		if got != want {
			t.Errorf("TestGet got: %v, want: %v", got, want)
		}
	}
}

func TestSet(t *testing.T) {
	hmap := New()

	tests := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}

	for i, test := range tests {
		hmap.Set(i, test)
		got, _ := hmap.Get(i)
		want := test
		if got != want {
			t.Errorf("TestSet got: %v, want: %v", got, want)
		}
	}

	for i, test := range tests {
		hmap.Set(i, test)
		got := hmap.Size()
		want := len(tests)
		if got != want {
			t.Errorf("TestSet got: %v, want: %v", got, want)
		}
	}
}

func TestDelete(t *testing.T) {
	hmap := New()

	tests := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}

	for i, test := range tests {
		hmap.Set(i, test)
		hmap.Delete(i)
		_, got := hmap.Get(i)
		want := false
		if got != want {
			t.Errorf("TestDelete got: %v, want: %v", got, want)
		}
	}
}

func TestKeys(t *testing.T) {
	hmap := New()

	tests := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g'}

	for i, test := range tests {
		hmap.Set(test, i)
	}

	keys := hmap.Keys()

	if got, want := len(keys), hmap.Size(); got != want {
		t.Errorf("TestKeys got: %v, want: %v", got, want)
	}

	for _, key := range keys {
		_, got := hmap.Get(key)
		want := true
		if got != want {
			t.Errorf("TestKeys got: %v, want: %v", got, want)
		}
	}
}
