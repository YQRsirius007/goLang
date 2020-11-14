package pos

//
// import (
// 	"context"
// 	"time"
// )
//
// type dialOptions struct {
// 	insecure bool
// 	timeout time.Duration
// }
// type ClientConn struct{
// 	authority string
// 	dopts dialOptions // 所有可选字段
// 	csMgr *connectivityStateManager
// }
// type DialOption interface{
// 	apply(*dialOptions)
// }
// func DialContext(ctx context.Context, target string, opts ...DialOptions)(){
// 	cc := &ClientConn{
// 		dopts: defaultDialOptions(), // 默认值选项
// 	}
//
// 	// 修改改选为用户的默认值
// 	for _, opt := range opts{
// 		opt.apply(&cc.dopts)
// 	}
// }
//
// // DialContext函数是一个grpc连接的创建函数 内部主要是构建ClientConn这个结构体 并作为返回值
// // defaultDialOptions函数返回的是系统提供给dopts字段的默认值 如果用户想要自定义可选属性可以通过可变参数opts来控制
// // 无论dopts字段如何增减，构造函数不需要改动 defaultDialOptions也可以从一个公有字段变为一个私有字段更加内聚 对调用者友好
