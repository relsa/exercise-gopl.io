package cyclic

import (
	"testing"
)

func TestIsCyclicStructPtr(t *testing.T) {
	type link struct {
		value string
		tail  *link
	}

	a, b := &link{value: "a"}, &link{value: "b"}
	a.tail, b.tail = b, a

	got := IsCyclic(a)
	want := true
	if got != want {
		t.Errorf("fail to detect cyclic: got %v, want %v", got, want)
	}

	b.tail = nil

	got = IsCyclic(a)
	want = false
	if got != want {
		t.Errorf("fail to detect cyclic: got %v, want %v", got, want)
	}

}

func TestIsCyclicStructSlice(t *testing.T) {
	type link struct {
		vals []*link
	}

	a, b := &link{}, &link{}
	a.vals = append(a.vals, b)
	b.vals = append(b.vals, a)

	got := IsCyclic(a)
	want := true
	if got != want {
		t.Errorf("fail to detect cyclic: got %v, want %v", got, want)
	}

	b.vals = []*link{}

	got = IsCyclic(a)
	want = false
	if got != want {
		t.Errorf("fail to detect cyclic: got %v, want %v", got, want)
	}
}

func TestIsCyclicStructMap(t *testing.T) {
	type link struct {
		vals map[string]*link
	}

	a, b := &link{make(map[string]*link)}, &link{make(map[string]*link)}
	a.vals["a"] = b
	b.vals["b"] = a

	got := IsCyclic(a)
	want := true
	if got != want {
		t.Errorf("fail to detect cyclic: got %v, want %v", got, want)
	}

	b.vals = make(map[string]*link)

	got = IsCyclic(a)
	want = false
	if got != want {
		t.Errorf("fail to detect cyclic: got %v, want %v", got, want)
	}
}
