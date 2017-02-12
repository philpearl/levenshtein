// Copyright (c) 2015, Arbo von Monkiewitsch All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package levenshtein

type Context struct {
	column []int
}

func (c *Context) reset(l int) {
	if cap(c.column) < l {
		c.column = make([]int, l)
	}
	c.column = c.column[:l]
	for i := range c.column {
		c.column[i] = i
	}
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
	var cost, lastdiag, olddiag int
	s1 := []rune(str1)
	s2 := []rune(str2)
	len_s1 := len(s1)
	len_s2 := len(s2)

	c.reset(len_s1 + 1)

	for x := 1; x <= len_s2; x++ {
		c.column[0] = x
		lastdiag = x - 1

		s2Rune := s2[x-1]

		for y := 1; y <= len_s1; y++ {
			olddiag = c.column[y]
			cost = 0
			if s1[y-1] != s2Rune {
				cost = 1
			}
			c.column[y] = min(
				c.column[y]+1,
				c.column[y-1]+1,
				lastdiag+cost)
			lastdiag = olddiag
		}
	}
	return c.column[len_s1]
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
