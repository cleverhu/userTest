package Setter

import (
	"fmt"
	"userTest/src/dbs"
	"userTest/src/models/UserModel"
	"userTest/src/result"
)

var UserSetter IUserSetter

func init() {
	UserSetter = NewUserSetterImpl()
}

type IUserSetter interface {
	AddUser(impl *UserModel.UserModelImpl) *result.ErrorResult
}

type UserSetterImpl struct {
}

func NewUserSetterImpl() *UserSetterImpl {
	return &UserSetterImpl{}
}

func (this *UserSetterImpl) AddUser(impl *UserModel.UserModelImpl) *result.ErrorResult {


	if 	dbs.Orm.First(&impl,"u_email = ? or u_name = ?",impl.Email,impl.Name).RecordNotFound()  && dbs.Orm.Save(&impl).RowsAffected == 1 {
		return &result.ErrorResult{
			Err:  nil,
			Data: impl,
		}
	} else {
		return &result.ErrorResult{
			Err:  fmt.Errorf("user is existed"),
			Data: nil,
		}
	}
}
