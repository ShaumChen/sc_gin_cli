package response

type Response struct {
	data interface{}
}

func Resp() *Response {
	return &Response{}
}

func (r *Response) Json(obj interface{}) *Response {
	r.data = obj
	return r
}

func (r *Response) String(data string) *Response {
	r.data = data
	return r
}

func (r *Response) Byte(data []byte) *Response {
	r.data = data
	return r
}

func (r *Response) GetData() interface{} {
	return r.data
}
