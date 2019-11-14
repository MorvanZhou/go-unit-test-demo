//go:generate mockgen -destination=mock_foo.go -package=_mock github.com/morvanzhou/unittest-demo/5mock YourInterface

package _mock

type YourInterface interface {
	YourInterFunc(ii int) int
}

type OtherInterface interface {
	OtherInterfaceFunc(ii int) int
}

func Myfunc(instance YourInterface, mi int) int {
	mi = mi - 1
	return instance.YourInterFunc(mi)
}
