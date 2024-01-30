package exception

import "net/http"

var (
	ErrBadRequest = NewApiException(http.StatusBadRequest, http.StatusText(http.StatusBadRequest)).WithHttpCode(http.StatusBadRequest)
)
