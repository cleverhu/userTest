package myValidators

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

type UserName string

func init() {
	validatorError["UserName"] = "用户名在6-20位之间"
	register("UserName", UserName("required,min=6,max=20").toFunc())
}

func (this UserName) toFunc() validator.Func {
	return func(fl validator.FieldLevel) bool {
		uName, ok := fl.Field().Interface().(string)
		return ok && this.validate(uName)
	}
}

func (this UserName) validate(v string) bool {
	v = strings.TrimSpace(v)
	if err := myValidator.Var(v, string(this)); err != nil {
		return false
	}
	if strings.Index(v, " ") != -1 {
		validatorError["UserName"] = "用户名不能包含空格"
		return false
	}
	return true
}
