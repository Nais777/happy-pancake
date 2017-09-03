package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestMakeGroupedSlice(t *testing.T) {
	tab := []struct {
		in                string
		outCount          int
		outLastUpsideDown bool
	}{
		{in: "-", outCount: 1, outLastUpsideDown: true},
		{in: "+", outCount: 1, outLastUpsideDown: false},
		{in: "-+", outCount: 2, outLastUpsideDown: false},
		{in: "+++", outCount: 1, outLastUpsideDown: false},
		{in: "-+----+---++++-", outCount: 7, outLastUpsideDown: true},
	}

	for _, tc := range tab {
		grpCount, upDown := getGroupCount(tc.in)
		assert.Equal(t, tc.outCount, grpCount)
		assert.Equal(t, tc.outLastUpsideDown, upDown)
	}
}
