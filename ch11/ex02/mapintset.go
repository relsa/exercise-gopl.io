package intset

import (
	"bytes"
	"fmt"
	"sort"
)

type MapIntSet struct {
	set map[int]bool
}

func (s *MapIntSet) Has(x int) bool {
	if s.set == nil {
		return false
	}
	return s.set[x]
}

func (s *MapIntSet) Add(x int) {
	if s.set == nil {
		s.set = make(map[int]bool)
	}
	s.set[x] = true
}

func (s *MapIntSet) UnionWith(t *MapIntSet) {
	if t.set == nil {
		return
	}
	if s.set == nil {
		s.set = make(map[int]bool)
	}

	for k, v := range t.set {
		if v {
			s.set[k] = true
		}
	}
}

func (s *MapIntSet) String() string {
	if s.set == nil {
		return "{}"
	}

	var ns []int
	for k, v := range s.set {
		if v {
			ns = append(ns, k)
		}
	}
	sort.Ints(ns)

	var buf bytes.Buffer
	buf.WriteByte('{')

	for i, n := range ns {
		if i > 0 {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", n)
	}

	buf.WriteByte('}')
	return buf.String()
}
