package result

import (
	"fmt"
	"userTest/src/myValidators"
)

type ErrorResult struct {
	Err  error
	Data interface{}
}

func Result(vs ...interface{}) *ErrorResult {
	length := len(vs)
	if length == 1 {
		if vs[0] == nil {
			return &ErrorResult{
				Err:  nil,
				Data: nil,
			}
		}

		if err, ok := vs[0].(error); ok {
			if err != nil {
				return &ErrorResult{
					Err:  err,
					Data: nil,
				}
			}
		}
	}

	if length == 2 {
		if vs[1] == nil {
			return &ErrorResult{
				Err:  nil,
				Data: vs[0],
			}
		}

		if err, ok := vs[1].(error); ok {
			if err != nil {
				return &ErrorResult{
					Err:  err,
					Data: nil,
				}
			}
		}
	}

	return &ErrorResult{
		Err:  fmt.Errorf("unformat error"),
		Data: nil,
	}
}

func (this *ErrorResult) Unwrap() interface{} {
	if this.Err != nil {
		myValidators.CheckErrors(this.Err)
		panic(this.Err.Error())
	}

	return this.Data
}
