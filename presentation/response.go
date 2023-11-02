package presentation

import "log"

type ErrorRes struct {
	Message string `json:"message"`
	Err     error  `json:"error"`
}

type ErrorResOpt func(*ErrorRes)

func WithMessage(message string) ErrorResOpt {
	return func(res *ErrorRes) {
		res.Message = message
	}
}

func WithErr(err error) ErrorResOpt {
	return func(res *ErrorRes) {
		res.Err = err
	}
}

func NewErrorRes(opts ...ErrorResOpt) ErrorRes {
	var errorRes ErrorRes
	for _, opt := range opts {
		opt(&errorRes)
	}

	log.Println(errorRes.Err)

	return errorRes
}
