package intset

import (
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
