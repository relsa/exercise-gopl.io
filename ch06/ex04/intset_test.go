package intset

import "testing"

func initIntSet(n ...int) *IntSet {
	is := new(IntSet)
	for _, x := range n {
		is.Add(x)
	}
	return is
}

func TestElems(t *testing.T) {
	ts := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{},
			want: []int{},
		},
		{
			nums: []int{1, 10, 100, 1000},
			want: []int{1, 10, 100, 1000},
		},
	}

	for _, tc := range ts {
		intSet := initIntSet(tc.nums...)
		got := intSet.Elems()

		if len(got) != len(tc.want) {
			t.Errorf("len=%d, want %d", len(got), len(tc.want))
		}

		for i := range tc.want {
			if got[i] != tc.want[i] {
				t.Errorf("[%d]: got %d, want %d", i, got[i], tc.want[i])
			}
		}
	}
}
