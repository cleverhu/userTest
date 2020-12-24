package myValidators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type Email string

func init() {
	validatorError["Email"] = "用户邮箱错误"
	register("Email", Email("required").toFunc())
}

func (this Email) toFunc() validator.Func {
	return func(fl validator.FieldLevel) bool {
		uEmail, ok := fl.Field().Interface().(string)
		return ok && this.validate(uEmail)
	}
}

func (this Email) validate(v string) bool {
	validatorError["Email"] = "必须输入邮箱"
	if err := myValidator.Var(v, string(this)); err != nil {
		return false
	}
	rxp := regexp.MustCompile(`[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)`)
	matches := rxp.FindAllStringSubmatch(v, -1)
	if len(matches) == 1 {
		return true
	}
	validatorError["Email"] = "邮箱输入错误"
	return false
}
