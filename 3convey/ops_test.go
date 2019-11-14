package _convey_test

import (
	"testing"

	. "github.com/morvanzhou/unittest-demo/3convey"
	. "github.com/smartystreets/goconvey/convey"
)


// 实时查看web端的所有test结果

func TestAdd(t *testing.T) {
	Convey("将两数相加", t, func() {
		So(Add(1, 2), ShouldEqual, 3)
	})
}

func TestSubtract(t *testing.T) {
	Convey("将两数相减", t, func() {
		So(Subtract(1, 2), ShouldEqual, -1)
	})
}

func TestMultiply(t *testing.T) {
	Convey("将两数相乘", t, func() {
		So(Multiply(3, 2), ShouldEqual, 6)
	})
}

func TestDivision(t *testing.T) {
	// 嵌套外层要传入 t
	Convey("将两数相除", t, func() {

		Convey("除以非 0 数", func() {
			num, err := Division(10, 2)
			So(err, ShouldBeNil)
			So(num, ShouldEqual, 5)
		})

		Convey("除以 0", func() {
			_, err := Division(10, 0)
			So(err, ShouldNotBeNil)
		})

	})
}


type inputs struct {
	a int
	b int
}
type expected struct {
	res int
	hasErr bool
}
func TestTableTest(t *testing.T) {
	tt := []struct {
		name string
		in inputs
		expected expected

	}{
		{"除以非 0 数", inputs{10, 2}, expected{5, false}},
		{"除以 0", inputs{10, 0}, expected{0, true}},
	}

	Convey("将两数相除", t, func() {
		for _, tc := range tt {
			Convey(tc.name, func() {
				num, err := Division(tc.in.a, tc.in.b)
				if tc.expected.hasErr {
					So(err, ShouldBeError)
				}
				So(num, ShouldEqual, tc.expected.res)
			})
		}

	})
}