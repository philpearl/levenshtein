// Copyright (c) 2015, Arbo von Monkiewitsch All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package levenshtein

type Context struct {
	intSlice []int
}

func (c *Context) getIntSlice(l int) []int {
	if cap(c.intSlice) < l {
		c.intSlice = make([]int, l)
	}
	return c.intSlice[:l]
}

func Distance(s1, s2 string) int {
	c := Context{}
	return c.Distance(s1, s2)
}

// The Levenshtein distance between two strings is defined as the minimum
// number of edits needed to transform one string into the other, with the
// allowable edit operations being insertion, deletion, or substitution of
// a single character
// http://en.wikipedia.org/wiki/Levenshtein_distance
//
// This implemention is optimized to use O(min(m,n)) space.
// It is based on the optimized C version found here:
// http://en.wikibooks.org/wiki/Algorithm_implementation/Strings/Levenshtein_distance#C
func (c *Context) Distance(str1, str2 string) int {
	s1 := []rune(str1)
	s2 := []rune(str2)

	len_s1 := len(s1)
	len_s2 := len(s2)

	if len_s2 == 0 {
		return len_s1
	}

	column := c.getIntSlice(len_s1 + 1)
	// Column[0] will be initialised at the start of the first loop before it
	// is read, unless len_s2 is zero, which we deal with above
	for i := 1; i <= len_s1; i++ {
		column[i] = i
	}

	for x := 0; x < len_s2; x++ {
		s2Rune := s2[x]
		column[0] = x + 1
		lastdiag := x

		for y := 0; y < len_s1; y++ {
			olddiag := column[y+1]
			cost := 0
			if s1[y] != s2Rune {
				cost = 1
			}
			column[y+1] = min(
				column[y+1]+1,
				column[y]+1,
				lastdiag+cost,
			)
			lastdiag = olddiag
		}
	}

	return column[len_s1]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}
