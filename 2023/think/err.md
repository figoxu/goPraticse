
error报错被层层包裹的时候，用is，as来判断err是否符合预期，原理是is，as实现是递归调用 Unwrap 来判断 err的 错误内容或者错误类型是否一致


Go 1.13以后添加了 error可以支持的 func Unwrap(err error) error 这个方法，相关的一些改变值得关注 包括我上面用到的：
fmt.Errorf(...)