package intset

import "testing"

func initIntSet(n ...int) *IntSet {
	is := new(IntSet)
	for _, x := range n {
		is.Add(x)
	}
	return is
}

func TestAddAll(t *testing.T) {
	ts := []struct {
		nums []int
		adds []int
		want []int
	}{
		{
			nums: []int{},
			adds: []int{1},
			want: []int{1},
		},
		{
			nums: []int{1},
			adds: []int{1, 10},
			want: []int{1, 10},
		},
		{
			nums: []int{1},
			adds: []int{10, 100, 1000},
			want: []int{1, 10, 100, 1000},
		},
	}

	for _, tc := range ts {
		is := initIntSet(tc.nums...)
		is.AddAll(tc.adds...)

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
