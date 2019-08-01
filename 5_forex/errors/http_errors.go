package errors

type BadRequest struct {
	ErrMsg string
}

func (b BadRequest) Error() string {
	return b.ErrMsg
}

type InternalError string // typealias

func (i InternalError) Error() string {
	return string(i)
}
