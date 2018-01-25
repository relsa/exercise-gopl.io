package intset

import "testing"

func initIntSet(n ...int) *IntSet {
	is := new(IntSet)
	for _, x := range n {
		is.Add(x)
	}
	return is
}

func TestIntersectWith(t *testing.T) {
	ts := []struct {
		sNums []int
		tNums []int
		want  []int
	}{
		{
			sNums: []int{},
			tNums: []int{},
			want:  []int{},
		},
		{
			sNums: []int{1, 10, 100},
			tNums: []int{1},
			want:  []int{1},
		},
		{
			sNums: []int{1},
			tNums: []int{1, 10, 100},
			want:  []int{1},
		},
		{
			sNums: []int{1, 10, 100},
			tNums: []int{1, 10, 100},
			want:  []int{1, 10, 100},
		},
	}

	for _, tc := range ts {
		sIntSet := initIntSet(tc.sNums...)
		tIntSet := initIntSet(tc.tNums...)
		sIntSet.IntersectWith(tIntSet)

		if sIntSet.Len() != len(tc.want) {
			t.Errorf("len=%d, want %d", sIntSet.Len(), len(tc.want))
		}

		for _, n := range tc.want {
			if !sIntSet.Has(n) {
				t.Errorf("not have %d", n)
			}
		}
	}
}

func TestDiffrenceWith(t *testing.T) {
	ts := []struct {
		sNums []int
		tNums []int
		want  []int
	}{
		{
			sNums: []int{},
			tNums: []int{},
			want:  []int{},
		},
		{
			sNums: []int{1, 10, 100},
			tNums: []int{1},
			want:  []int{10, 100},
		},
		{
			sNums: []int{1},
			tNums: []int{1, 10, 100},
			want:  []int{},
		},
		{
			sNums: []int{1, 10, 100},
			tNums: []int{1, 10, 100},
			want:  []int{},
		},
	}

	for _, tc := range ts {
		sIntSet := initIntSet(tc.sNums...)
		tIntSet := initIntSet(tc.tNums...)
		sIntSet.DifferenceWith(tIntSet)

		if sIntSet.Len() != len(tc.want) {
			t.Errorf("len=%d, want %d", sIntSet.Len(), len(tc.want))
		}

		for _, n := range tc.want {
			if !sIntSet.Has(n) {
				t.Errorf("not have %d", n)
			}
		}
	}
}

func TestSymmentricDifference(t *testing.T) {
	ts := []struct {
		sNums []int
		tNums []int
		want  []int
	}{
		{
			sNums: []int{},
			tNums: []int{},
			want:  []int{},
		},
		{
			sNums: []int{1, 10, 100},
			tNums: []int{1},
			want:  []int{10, 100},
		},
		{
			sNums: []int{1},
			tNums: []int{1, 10, 100},
			want:  []int{10, 100},
		},
		{
			sNums: []int{1, 10, 100},
			tNums: []int{1, 10, 100},
			want:  []int{},
		},
	}

	for _, tc := range ts {
		sIntSet := initIntSet(tc.sNums...)
		tIntSet := initIntSet(tc.tNums...)
		sIntSet.SymmetricDifference(tIntSet)

		if sIntSet.Len() != len(tc.want) {
			t.Errorf("len=%d, want %d", sIntSet.Len(), len(tc.want))
		}

		for _, n := range tc.want {
			if !sIntSet.Has(n) {
				t.Errorf("not have %d", n)
			}
		}
	}
}
