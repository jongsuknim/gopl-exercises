package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Len() int {
	l := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				l++
			}
		}
	}
	return l
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word >= len(s.words) {
		return
	} else {
		s.words[word] &= ^(1 << bit)
	}
}

func (s *IntSet) Clear() {
	s.words = make([]uint64, 0)
}

func (s *IntSet) Copy() *IntSet {
	clone := IntSet{words: make([]uint64, len(s.words))}
	copy(clone.words, s.words)
	return &clone
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) AddAll(values ...int) {
	for _, v := range values {
		s.Add(v)
	}
}

func (s *IntSet) InstersectWith(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] &= t.words[i]
		} else {
			s.words = s.words[:i]
			return
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] &= ^t.words[i]
		} else {
			return
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	u := s.Copy()
	u.InstersectWith(t)

	s.UnionWith(t)
	s.DifferenceWith(u)
}

func (s *IntSet) Elems() []int64 {
	retval := make([]int64, 0)

	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				retval = append(retval, int64(64*i+j))
			}
		}
	}
	return retval
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())

	z := x.Copy()
	fmt.Println("z", z.String())
	fmt.Println("z.Len", z.Len())

	z.Remove(100)
	z.Remove(144)
	fmt.Println(z.String())

	z.Clear()
	fmt.Println(z.String())

	z.AddAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 1000, 10000, 100000)
	fmt.Println(z.String())

	z.SymmetricDifference(&x)
	fmt.Println(z.String())
	fmt.Println(z.Elems())
}
