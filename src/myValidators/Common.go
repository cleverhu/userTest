package myValidators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
)

var (
	myValidator    *validator.Validate
	validatorError map[string]string
)

func init() {
	validatorError = make(map[string]string)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		myValidator = v
	} else {
		log.Fatal("validator init failed")
	}
}

func register(tag string, fn validator.Func) {
	if err := myValidator.RegisterValidation(tag, fn); err != nil {
		log.Fatalf("validator %s error", tag)
	}
}

func CheckErrors(err error) {
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, err := range errs {
			if v, exists := validatorError[err.Tag()]; exists {
				panic(v)
			}
		}
	}
}
