package errors

import (
	"errors"
	"fmt"
)

type HttpError struct {
	Errors map[string][]string `json:"errors"`

	Code int `json:"code"`
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("%s", "HTTPError")
}

func FromError(err error) *HttpError {
	if err == nil {
		return nil
	}
	// 提取指定类型的错误实例‌
	if se := new(HttpError); errors.As(err, &se) {
		return se
	}
	return &HttpError{Code: 500}
}

func NewHttpError(code int, filed string, detail string) *HttpError {
	return &HttpError{
		Errors: map[string][]string{
			filed: {detail},
		},
		Code: code,
	}
}
