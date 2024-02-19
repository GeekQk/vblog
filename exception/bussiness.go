package exception

import "net/http"

var (
	ErrBadRequest = NewApiException(http.StatusBadRequest, http.StatusText(http.StatusBadRequest)).WithHttpCode(http.StatusBadRequest)
	// 为认证, 没有登录, Token没传递
	ErrUnauthorized = NewApiException(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)).WithHttpCode(http.StatusUnauthorized)
	// 鉴权失败, 认证通过，但是没有权限操作 该接口
	ErrPermissionDeny = NewApiException(http.StatusForbidden, http.StatusText(http.StatusForbidden)).WithHttpCode(http.StatusForbidden)
)
