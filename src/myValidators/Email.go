package myValidators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
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
	v = strings.TrimSpace(v)
	validatorError["Email"] = "必须输入邮箱"
	if err := myValidator.Var(v, string(this)); err != nil {
		return false
	}
	rxp := regexp.MustCompile(`[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)`)
	matches := rxp.FindAllStringSubmatch(v, -1)
	if len(matches) == 1 {
		if strings.Index(v, " ") != -1 {
			validatorError["Email"] = "邮箱不能包含空格"
			return false
		}
		return true
	}
	validatorError["Email"] = "邮箱输入错误"
	return false
}
