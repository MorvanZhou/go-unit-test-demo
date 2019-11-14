package _mock_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/morvanzhou/unittest-demo/5mock"
	"github.com/stretchr/testify/assert"
)

func TestMyfunc(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// get mocked interface
	m := _mock.NewMockMyInterface(ctrl)

	// set test case
	m.EXPECT().
		MyInterFunc(gomock.Eq(1)).		// when use MyInterFunc(1)
		Return(101).					// should return 101
		AnyTimes()

	m.EXPECT().
		MyInterFunc(gomock.Not(1)).		// when use MyInterFunc(other than 1)
		Return(100).					// should return 100
		AnyTimes()

	assert.Equal(t, 101, _mock.Myfunc(m, 1))
	assert.Equal(t, 100, _mock.Myfunc(m, 98))
}

func TestMyfuncOrder(t *testing.T) {
	tt := []struct {
		name   string
		input  int
		expect int
	}{
		{"test1", 1, 1},
		{"test2", 2, -1},
		{"test3", 33, -1},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := _mock.NewMockMyInterface(ctrl)

	gomock.InOrder(
		m.EXPECT().MyInterFunc(gomock.Eq(1)).Return(1).AnyTimes(),
		m.EXPECT().MyInterFunc(gomock.Any()).Return(-1).AnyTimes(),
	)

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expect, _mock.Myfunc(m, tc.input))
		})
	}
}

