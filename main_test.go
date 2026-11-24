package main

import "testing"

func TestSingleElement(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}, 2},
		{[]int{3, 3, 7, 7, 10, 11, 11}, 10},
		{[]int{5}, 5},
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 2}, 2},
		{[]int{2, 3, 3}, 2},
	}
	for _, c := range cases {
		if got := SingleElement(c.in); got != c.want {
			t.Errorf("SingleElement(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
