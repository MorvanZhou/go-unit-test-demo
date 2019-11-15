# golang Unit-test demos

## 使用 test table 进行批量测试
[1testtable](/1testtable/func_test.go)
对一个功能, 汇总所有的case, 统一集中测试

## 检查测试代码覆盖率
[2coverage](/2coverage/func_test.go)
查看我们是否已经对所有代码都进行了测试, 因为有时我们只能测试到一部分的代码. 比如通常我们会忽略 err 部分的代码测试.

## convey 及时同步测试结果
[3convey](/3convey/ops_test.go)
使用web端, 无需手动触发, 及时同步最新所有的测试结果, 只要有不通过就能进行通知提醒.

## gin 框架测试
[4gin](/4gin/gin_test.go)
如果使用了go 的 gin 框架搭建了server, 就需要封装 http 请求, 对 API 进行测试.

## mock 模拟 struct 返回测试
[5mock](/5mock/foo_test.go)
在团队开发中, 有时候需要依赖别人开发的功能, 当别人还没开发完毕时, 我们可以模拟别人的函数或者 struct 功能进行自己功能的测试.

## mock API 请求
[6api](/6apimock/api_test.go)
如果有调用正在开发中的其他微服务, 或者想摆脱其他微服务的依赖, 模拟其他微服务 API 的返回信息, 不用等待真实的微服务.
