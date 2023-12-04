package common

type BusinessCode int
type Result struct {
	Code    BusinessCode `json:"code"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data"`
}

func (r *Result) Success(data any) *Result {
	r.Code = 200
	r.Message = "success"
	r.Data = data
	return r
}

func (r *Result) Fail(code BusinessCode, message string) *Result {
	r.Code = code
	r.Message = message
	return r
}
