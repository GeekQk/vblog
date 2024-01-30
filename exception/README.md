# 业务异常


## 为什么需要业务异常

判断 令牌过期

业务异常的使用方式:
```go
err := Biz_Call()
if err == TokenExpired {

}
```

字符串比对, 可能造成误杀
```go
access token expired %f minutes

hasPrefix("access token expired")
```

设计一套业务专用的业务异常, 通常设计为异常码(Error Code):
```go
// err.ErrorCode == xxxx
if exception.IsTokenExpired(err) {

}
```

## 怎么设计业务异常

本书需要兼容Error的场景:
```go
func XXX() error
```

go 里面的Error是个接口
```go
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
	Error() string
}
```

fmt包里面提供的Error实现
```go
type wrapError struct {
	msg string
	err error
}

func (e *wrapError) Error() string {
	return e.msg
}

func (e *wrapError) Unwrap() error {
	return e.err
}

type wrapErrors struct {
	msg  string
	errs []error
}

func (e *wrapErrors) Error() string {
	return e.msg
}

func (e *wrapErrors) Unwrap() []error {
	return e.errs
}
```


如何定义自定义异常: APIException
```go
func NewAPIException(code int, msg string) *APIException {
	return &APIException{
		code: code,
		msg:  msg,
	}
}

// error的自定义实现
type APIException struct {
	code int
	msg  string
}

func (e *APIException) Error() string {
	return e.msg
}

func (e *APIException) Code() int {
	return e.code
}
```


## 定义业务异常

1. 定义 TokenExired 5000

```go
// 这个模块定义的业务异常
// token expired %f minitues
// 约定俗成:  ErrXXXXX 来进行自定义异常定义, 方便快速在包里搜索
var (
	ErrAccessTokenExpired  = exception.NewAPIException(5000, "access token expired")
	ErrRefreshTokenExpired = exception.NewAPIException(5001, "refresh token expired")
)
```

2. 使用自定义异常
```go
if aDelta > 0 {
    return ErrAccessTokenExpired.WithMessagef("access token expired %f minutes", aDelta)
}
```

## 如果判断异常是否相等喃

1. 基于断言后根据Code来进行业务异常判断
```go
if e, ok := err.(*exception.APIException); ok {
    t.Log(e.String())
    // 判断该异常是不是 TokenExpired异常
    if e.Code == token.ErrAccessTokenExpired.Code {
        t.Log(e.String())
    }
}
```

```go
// 	 exception.IsException(err, token.ErrAccessTokenExpired)
// 给一个异常判断的方法
func IsException(err error, e *APIException) bool {
	if targe, ok := err.(*APIException); ok {
		return targe.Code == e.Code
	}

	return false
}
```