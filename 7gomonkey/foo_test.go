package _gomonkey_test

import (
	"github.com/agiledragon/gomonkey"
	_gomonkey "github.com/morvanzhou/unittest-demo/7gomonkey"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestWrapMysqlFunc(t *testing.T) {
	tt := []struct {
		name    string
		input   int
		mockOut []string
		expect  string
	}{
		{"test1", 2, []string{"a2", "b2"}, "2a2"},
		{"test2", 3, []string{"a1", "b1"}, "3a1"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			patch := gomonkey.ApplyFunc(_gomonkey.MysqlFunc, func(_ int) []string {
				log.Printf("%v", tc.input)
				return tc.mockOut
			})
			defer patch.Reset()

			users := _gomonkey.WrapMysqlFunc(tc.input)

			assert.Equal(t, tc.expect, users[0])

		})
	}
}
