package response

func (mr *Response) Status(status int) *Response {
	if mr != nil && status > 99 && status < 600 {
		mr.statusCode = status
	}
	return mr
}

func (mr *Response) Content(content any) *Response {
	if mr != nil {
		mr.contentData = content
	}
	return mr
}
