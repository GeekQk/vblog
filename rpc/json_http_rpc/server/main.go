package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/GeekQk/vblog/rpc/json_http_rpc/service"
)

var _ service.HelloServices = (*HelloService)(nil)

type HelloService struct{}

// Hello的逻辑 就是 将对方发送的消息前面添加一个Hello 然后返还给对方
// 由于我们是一个rpc服务, 因此参数上面还是有约束：
//
//	第一个参数是请求
//	第二个参数是响应
//
// 可以类比Http handler
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

// 首先我们依然要解决JSON编解码的问题, 我们需要将HTTP接口的Handler参数传递给jsonrpc,
// 因此需要满足jsonrpc接口, 因此我们需要提前构建也给conn io.ReadWriteCloser, writer现成的 reader就是request的body, 直接内嵌就可以
func NewRPCReadWriteCloserFromHTTP(w http.ResponseWriter, r *http.Request) *RPCReadWriteCloser {
	return &RPCReadWriteCloser{w, r.Body}
}

type RPCReadWriteCloser struct {
	io.Writer
	io.ReadCloser
}

func main() {
	// 把我们的对象注册成一个rpc的 receiver
	// 其中rpc.Register函数调用会将对象类型中所有满足RPC规则的对象方法注册为RPC函数，
	// 所有注册的方法会放在“HelloService”服务空间之下
	rpc.RegisterName("HelloService", new(HelloService))
	// RPC的服务架设在“/jsonrpc”路径，
	// 在处理函数中基于http.ResponseWriter和http.Request类型的参数构造一个io.ReadWriteCloser类型的conn通道。
	// 然后基于conn构建针对服务端的json编码解码器。
	// 最后通过rpc.ServeRequest函数为每次请求处理一次RPC方法调用
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		conn := NewRPCReadWriteCloserFromHTTP(w, r)
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":8301", nil)
}

//本地测试:localhost:8301/jsonrpc
//测试: curl -X POST -d '{"method":"HelloService.Hello","params":["hello"],"id":1}' http://localhost:8301/jsonrpc
//返回结果:{"id":1,"result":"hello:hello","error":null}
