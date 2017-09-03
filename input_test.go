package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestMakeGroupedSlice(t *testing.T) {
	tab := []struct {
		in  string
		out []bool
	}{
		{in: "-", out: []bool{false}},
		{in: "+", out: []bool{true}},
		{in: "-+", out: []bool{false, true}},
		{in: "+++", out: []bool{true}},
		{in: "-+----+---++++-", out: []bool{false, true, false, true, false, true, false}},
	}

	for _, tc := range tab {
		out := makeGroupedSlice(tc.in)
		assert.Equal(t, tc.out, out)
	}
}
