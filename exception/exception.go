package exception

import (
	"fmt"

	"github.com/infraboard/mcube/tools/pretty"
)

func NewApiException(code int, msg string) *APIExeption {
	return &APIExeption{
		Code:   code,
		Reason: msg,
	}
}

type APIExeption struct {
	HttpCode int    `json:"-"`
	Code     int    `json:"code"`
	Reason   string `json:"reason"`
	Message  string `json:"message"`
}

func (e *APIExeption) Error() string {
	return fmt.Sprintf("%s,%s", e.Reason, e.Message)
}

// 设计为链式调用 返回APIExeption
func (e *APIExeption) WithMessage(msg string) *APIExeption {
	e.Message = msg
	return e
}

// 设计为链式调用 返回APIExeption http code
func (e *APIExeption) WithHttpCode(code int) *APIExeption {
	e.HttpCode = code
	return e
}

func (e *APIExeption) WithMessagef(format string, a ...any) *APIExeption {
	e.Message = fmt.Sprintf(format, a...)
	return e
}

func IsException(err error, e APIExeption) bool {
	if target, ok := err.(*APIExeption); ok {
		return target.Code == e.Code
	}
	return false
}

func (e *APIExeption) String() string {
	return pretty.ToJSON(e)
}
