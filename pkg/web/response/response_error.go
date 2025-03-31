package response

import (
	"fmt"
	"net/http"
)

func (mr *HttpResponse) String() string {
	if mr == nil {
		return "status_code:, message:"
	}
	s := fmt.Sprintf("status_code: %d, message: %s", mr.StatusCode, mr.Message)
	return s
}

func (r *Response) Err(err *HttpResponse) *Response {
	if r != nil {
		r.err = err
	}
	return r
}

func (r *Response) MessageErr(message string) *Response {
	r.err = &HttpResponse{
		StatusCode: r.statusCode,
		Message:    message,
	}

	return r
}

func (r *Response) UnauthorizedErr(message string) *Response {
	if r != nil {
		r.err = &HttpResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    message,
		}
	}
	return r
}
