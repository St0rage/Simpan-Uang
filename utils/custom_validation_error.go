package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

type ApiError struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

func CustomValidationErr(err error) map[string]string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make(map[string]string, len(ve))
		for _, fe := range ve {
			out[strcase.ToSnake(fe.Field())] = MsgForTag(fe.Tag(), fe.Param())
		}
		return out
	}
	return nil
}

func MsgForTag(tag string, param string) string {
	switch tag {
	case "json":
		return "Format tidak valid"
	case "required":
		return "Tidak boleh kosong"
	case "email":
		return "Email tidak valid"
	case "min":
		return "Minimal harus " + param + " karakter"
	case "max":
		return "Maksimal " + param + " karakter"
	case "eqfield":
		return "Password tidak sesuai"
	case "numeric":
		return "Harus berupa angka"
	case "gt":
		return "Minimal Rp " + param
	}

	return ""
}
