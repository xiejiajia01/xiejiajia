package service

import "github.com/google/wire"

// ProviderSet is service providers.
/*添加第二个参数 NewVerifyCodeService。这个函数是用来生成 VerifyCodeService 服务的，定义在internal/service/verifycode.go 中。以上代码的意思就是告知 wire 依赖注入系统，
如果需要 VerifyCodeService 的话，使用 NewVerifyCodeService 函数来构建。 */
var ProviderSet = wire.NewSet(NewGreeterService, NewVerifyCodeService)
