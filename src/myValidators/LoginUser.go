package myValidators

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

type LoginUser string

func init() {
	validatorError["LoginUserName"] = "用户名为空"
	register("LoginUser", LoginUser("required,min=0,max=20").toFunc())
}

func (this LoginUser) toFunc() validator.Func {
	return func(fl validator.FieldLevel) bool {
		uName, ok := fl.Field().Interface().(string)
		return ok && this.validate(uName)
	}
}

func (this LoginUser) validate(v string) bool {
	v = strings.TrimSpace(v)
	if err := myValidator.Var(v, string(this)); err != nil {
		return false
	}
	if v == "" {
		return true
	}

	if strings.Index(v, " ") != -1 {
		validatorError["LoginUserName"] = "用户名不能包含空格"
		return false
	}
	return true
}
