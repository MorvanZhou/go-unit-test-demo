//go:generate mockgen -destination=mock_foo.go -package=_mock github.com/morvanzhou/unittest-demo/5mock MyInterface

package _mock

type MyInterface interface {
	MyInterFunc(ii int) int
}

type OtherInterface interface {
	OtherInterfaceFunc(ii int) int
}

func Myfunc(instance MyInterface, mi int) int {
	return instance.MyInterFunc(mi)
}
