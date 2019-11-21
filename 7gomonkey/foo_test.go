package _gomonkey_test

import (
	"github.com/agiledragon/gomonkey"
	_gomonkey "github.com/morvanzhou/unittest-demo/7gomonkey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrapMysqlFunc(t *testing.T) {
	tt := []struct {
		name    string
		input   int      // 外部输入
		mockOut []string // mock mysql 的输出
		expect  []string // 结合外部输入和 mysql 输出的结果
	}{
		{"test1", 2, []string{"a2", "b2"}, []string{"2a2", "2b2"}},
		{"test2", 3, []string{"a1", "b1"}, []string{"3a1", "3b1"}},
		{"test3", 21, []string{"d10", "f11"}, []string{"21d10", "21f11"}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			patch := gomonkey.ApplyFunc(_gomonkey.MysqlFunc, func(_ int) []string {
				return tc.mockOut
			})
			defer patch.Reset()

			users := _gomonkey.WrapMysqlFunc(tc.input)

			assert.Equal(t, tc.expect, users)

		})
	}
}
