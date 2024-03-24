package service

const HelloServiceName = "HelloService"

type HelloServices interface {
	Hello(request string, reply *string) error
}

type HelloServiceImpl struct{}

// Hello的逻辑 就是 将对方发送的消息前面添加一个Hello 然后返还给对方
// 由于我们是一个rpc服务, 因此参数上面还是有约束：
//
//	第一个参数是请求
//	第二个参数是响应
//
// 可以类比Http handler
func (p *HelloServiceImpl) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}
