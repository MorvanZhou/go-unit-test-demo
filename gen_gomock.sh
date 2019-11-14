# use mockgen
mockgen -source=5mock/foo.go -destination=5mock/mock_foo.go -package=_mock

# use go generate
# 在package的go文件首行要提前定义好 //go:generate
go generate github.com/morvanzhou/unittest-demo/5mock
