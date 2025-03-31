package response

import (
	"encoding/json"
	"net/http"
)

type HttpResponse struct {
	StatusCode int          `json:"status_code"`
	Message    string       `json:"message,omitempty"`
	Data       any          `json:"data,omitempty"`
	Validate   []Validation `json:"validate,omitempty"`
}

type Validation struct {
	Field    string   `json:"field,omitempty"`
	Messages []string `json:"messages,omitempty"`
}

type Response struct {
	writer      http.ResponseWriter
	err         *HttpResponse
	headers     map[string]string
	statusCode  int
	contentType string
	contentData any
}

func To(w http.ResponseWriter) *Response {
	return &Response{
		writer:     w,
		statusCode: http.StatusOK,
	}
}

func (mr *Response) SendJSON() error {
	if mr.writer == nil {
		return nil
	}
	mr.writeContentTypeAndHeaders("application/json; charset=utf-8")
	if mr.err != nil {
		return json.NewEncoder(mr.writer).Encode(mr.err)
	}
	return json.NewEncoder(mr.writer).Encode(&HttpResponse{
		StatusCode: mr.statusCode,
		Data:       mr.contentData,
	})
}

func (mr *Response) writeContentTypeAndHeaders(contentType string) {
	mr.writer.Header().Add("Content-Type", contentType)
	mr.writeHeaders()
	mr.writeStatusCode()
}

func (mr *Response) writeHeaders() {
	for k, v := range mr.headers {
		mr.writer.Header().Add(k, v)
	}
}

func (mr *Response) writeStatusCode() {
	if mr.err != nil {
		mr.writer.WriteHeader(mr.err.StatusCode)
	} else {
		mr.writer.WriteHeader(mr.statusCode)
	}
}
