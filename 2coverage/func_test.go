package _coverage_test

import (
	"testing"

	"github.com/morvanzhou/unittest-demo/2coverage"

	"github.com/stretchr/testify/assert"
)

func TestPartCover1(t *testing.T) {
	tt := []struct {
		name     string
		in       int
		expected bool
	}{
		// test cases
		{"put 2", 2, false},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			res := _coverage.IsOne(tc.in)
			assert.Equal(t, tc.expected, res)

		})
	}

}

func TestPartCover2(t *testing.T) {
	tt := []struct {
		name     string
		in       int
		expected bool
	}{
		// test cases
		{"put 1", 1, true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			res := _coverage.IsOne(tc.in)
			assert.Equal(t, tc.expected, res)

		})
	}

}


