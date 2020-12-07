package models

type ResponseData struct {
	Data  interface{}
	Error error
}

func ResponseResult(data interface{}, error error) *ResponseData {
	return &ResponseData{
		Data:  data,
		Error: error,
	}
}
