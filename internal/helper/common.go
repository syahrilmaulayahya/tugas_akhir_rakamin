package helper

import "github.com/go-playground/validator/v10"

type ErrorStruct struct {
	Err  error `json:"err,omitempty"`
	Code int   `json:"code,omitempty"`
}

var Validate = validator.New()
