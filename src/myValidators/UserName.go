package myValidators

import "github.com/go-playground/validator/v10"

type UserName string

func init() {
	validatorError["UserName"] = "用户名在6-20位之间"
	register("UserName", PassWord("required,min=6,max=20").toFunc())
}

func (this UserName) toFunc() validator.Func {
	return func(fl validator.FieldLevel) bool {
		uName, ok := fl.Field().Interface().(string)
		return ok && this.validate(uName)
	}
}

func (this UserName) validate(v string) bool {
	if err := myValidator.Var(v, string(this)); err != nil {
		return false
	}
	return true
}
