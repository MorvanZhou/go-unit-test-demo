package _monkey_test

import (
	"bou.ke/monkey"
	_gomonkey "github.com/morvanzhou/unittest-demo/7monkey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrapMysqlFunc(t *testing.T) {
	tt := []struct {
		name    string
		input   int    // 外部输入
		mockOut string // mock mysql 的输出
		expect  string // 拼接外部输入和 mysql 输出的结果 (input+mockOut)
	}{
		{"test1", 2, "a2", "2a2"},
		{"test2", 3, "b1", "3b1"},
		{"test3", 21, "d10", "21d10"},
	}

	for _, tc := range tt {

		t.Run(tc.name, func(t *testing.T) {
			monkey.Patch(_gomonkey.MysqlFunc, func(age int) string {
				return tc.mockOut
			})

			user := _gomonkey.WrapMysqlFunc(tc.input)

			assert.Equal(t, tc.expect, user)
		})

	}

}
