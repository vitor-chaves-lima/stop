package logic

type Error struct {
	Code   string `json:"code"`
	Detail error  `json:"detail"`
}

func (e Error) Error() string {
	return e.Detail.Error()
}

func NewError(code string, detail error) *Error {
	return &Error{
		Code:   code,
		Detail: detail,
	}
}
