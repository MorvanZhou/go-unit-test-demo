package _testtable_test

import (
	"testing"

	"github.com/morvanzhou/unittest-demo/1testtable"

	"github.com/stretchr/testify/assert"
)



func TestIsOne(t *testing.T) {

	// test table
	tt := []struct {
		name     string		// test name
		in       int		// test input
		expected bool		// expected output
	}{
		// test cases
		{"1 put 1", 1, true},
		{"2 put 2", 2, false},
		{"3 put 2 failed", 2, true},
		{"4 put -2", -2, false},
	}

	for _, tc := range tt {

		// wrap test case, prevent break
		t.Run(tc.name, func(t *testing.T) {

			res := _testtable.IsOne(tc.in)
			assert.Equal(t, tc.expected, res)

		})
	}

}


func TestIsOneBreak(t *testing.T) {

	// test table
	tt := []struct {
		name     string		// test name
		in       int		// test input
		expected bool		// expected output
	}{
		// test cases
		{"1 put 1", 1, true},
		{"2 put 2", 2, false},
		{"3 put 2 failed", 2, true},
		{"4 put -2", -2, false},
	}

	for _, tc := range tt {

		// 3rd case will break test loop
		res := _testtable.IsOne(tc.in)
		assert.Equal(t, tc.expected, res)

	}
}