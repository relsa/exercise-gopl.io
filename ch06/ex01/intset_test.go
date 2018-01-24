package intset

import "testing"

func initIntSet(n ...int) *IntSet {
	is := new(IntSet)
	for _, x := range n {
		is.Add(x)
	}
	return is
}

func TestLen(t *testing.T) {
	ts := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{},
			want: 0,
		},
		{
			nums: []int{1},
			want: 1,
		},
		{
			nums: []int{1, 10, 100, 1000},
			want: 4,
		},
	}

	for _, tc := range ts {
		is := initIntSet(tc.nums...)
		got := is.Len()
		if got != tc.want {
			t.Errorf("got len=%d, want %d", got, tc.want)
		}
	}
}

func TestRemove(t *testing.T) {
	ts := []struct {
		nums []int
		rm   []int
		want []int
	}{
		{
			nums: []int{},
			rm:   []int{1},
			want: []int{},
		},
		{
			nums: []int{1, 10, 100, 1000},
			rm:   []int{1, 100},
			want: []int{10, 1000},
		},
	}

	for _, tc := range ts {
		is := initIntSet(tc.nums...)
		for _, n := range tc.rm {
			is.Remove(n)
		}

		if is.Len() != len(tc.want) {
			t.Errorf("len=%d, want %d", is.Len(), len(tc.want))
		}

		for _, n := range tc.want {
			if !is.Has(n) {
				t.Errorf("not have %d", n)
			}
		}
	}
}

func TestClear(t *testing.T) {
	is := initIntSet(1, 10, 100, 1000)
	is.Clear()
	if is.Len() != 0 {
		t.Errorf("not cleared")
	}
}

func TestCopy(t *testing.T) {
	ts := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{},
			want: []int{},
		},
		{
			nums: []int{1},
			want: []int{1},
		},
		{
			nums: []int{1, 10, 100, 1000},
			want: []int{1, 10, 100, 1000},
		},
	}

	for _, tc := range ts {
		is := initIntSet(tc.nums...)

		if is.Len() != len(tc.want) {
			t.Errorf("len=%d, want %d", is.Len(), len(tc.want))
		}

		for _, n := range tc.want {
			if !is.Has(n) {
				t.Errorf("not have %d", n)
			}
		}
	}
}
