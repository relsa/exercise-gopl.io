package intset

import (
	"math/rand"
	"testing"
)

func Test(t *testing.T) {
	tests := []int{0, 1, 2, 3, 100, 1000}

	is := &IntSet{}
	mis := &MapIntSet{}

	// Add {{{
	adds := []int{1, 3, 100}
	for _, a := range adds {
		is.Add(a)
		mis.Add(a)
	}

	for _, tc := range tests {
		if want, got := mis.Has(tc), is.Has(tc); want != got {
			t.Errorf("case Has(%d) after Add(x): want %v, but got %v", tc, want, got)
		}
	}
	// }}}

	// UnionWith {{{
	otherIs := &IntSet{}
	otherMis := &MapIntSet{}

	otherAdds := []int{1000}
	for _, a := range otherAdds {
		otherIs.Add(a)
		otherMis.Add(a)
	}

	is.UnionWith(otherIs)
	mis.UnionWith(otherMis)

	for _, tc := range tests {
		if want, got := mis.Has(tc), is.Has(tc); want != got {
			t.Errorf("case Has(%d) after UnionWith(x): want %v, but got %v", tc, want, got)
		}
	}
	// }}}

	// String {{{
	if want, got := mis.String(), is.String(); want != got {
		t.Errorf("case String(): want %v, but got %v", want, got)
	}
	// }}}
}

const (
	max  = 10000000
	seed = 0
)

func BenchmarkIntSet_Add(b *testing.B) {
	rand.Seed(seed)

	for i := 0; i < b.N; i++ {
		is := &IntSet{}
		for j := 0; j < 10000; j++ {
			is.Add(rand.Intn(max))
		}
	}
}

func BenchmarkMapIntSet_Add(b *testing.B) {
	rand.Seed(seed)

	for i := 0; i < b.N; i++ {
		is := &MapIntSet{}
		for j := 0; j < 10000; j++ {
			is.Add(rand.Intn(max))
		}
	}
}

func BenchmarkIntSet_Has(b *testing.B) {
	rand.Seed(seed)

	is := &IntSet{}
	for i := 0; i < 10000; i++ {
		is.Add(rand.Intn(max))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		is.Has(rand.Intn(max))
	}
}

func BenchmarkMapIntSet_Has(b *testing.B) {
	rand.Seed(seed)

	is := &MapIntSet{}
	for i := 0; i < 10000; i++ {
		is.Add(rand.Intn(max))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		is.Has(rand.Intn(max))
	}
}

func BenchmarkIntSet_UnionWith(b *testing.B) {
	rand.Seed(seed)

	isx := &IntSet{}
	isy := &IntSet{}
	for i := 0; i < 10000; i++ {
		isx.Add(rand.Intn(max))
		isy.Add(rand.Intn(max))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		isx.UnionWith(isy)
	}
}

func BenchmarkMapIntSet_UnionWith(b *testing.B) {
	rand.Seed(seed)

	isx := &MapIntSet{}
	isy := &MapIntSet{}

	for i := 0; i < 10000; i++ {
		isx.Add(rand.Intn(max))
		isy.Add(rand.Intn(max))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		isx.UnionWith(isy)
	}
}

func BenchmarkIntSet_String(b *testing.B) {
	rand.Seed(seed)

	is := &IntSet{}
	for i := 0; i < 10000; i++ {
		is.Add(rand.Intn(max))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		is.String()
	}
}

func BenchmarkMapIntSet_String(b *testing.B) {
	rand.Seed(seed)

	is := &MapIntSet{}
	for i := 0; i < 10000; i++ {
		is.Add(rand.Intn(max))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		is.String()
	}
}
