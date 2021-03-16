package infrastruct

import "net/http"

type CustomErrors struct {
	msg  string
	Code int
}

func (c *CustomErrors) Error() string {
	return c.msg
}

func NewError(msg string, code int) *CustomErrors {
	return &CustomErrors{
		msg:  msg,
		Code: code,
	}
}

var (
	ErrorBadRequest = NewError("bad Request", http.StatusBadRequest)
	ErrorNotFound   = NewError("value not enough", http.StatusNotFound)
)
