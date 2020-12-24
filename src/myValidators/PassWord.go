package myValidators

import "github.com/go-playground/validator/v10"

type PassWord string

func init() {
	validatorError["PassWord"] = "密码必须大于6位"
	register("PassWord", PassWord("required,min=6").toFunc())
}

func (this PassWord) toFunc() validator.Func {
	return func(fl validator.FieldLevel) bool {
		uPwd, ok := fl.Field().Interface().(string)
		return ok && this.validate(uPwd)
	}
}

func (this PassWord) validate(v string) bool {
	if err := myValidator.Var(v, string(this)); err != nil {
		return false
	}
	return true
}
