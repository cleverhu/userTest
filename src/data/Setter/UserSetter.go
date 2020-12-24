package Setter

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
	"userTest/src/dbs"
	"userTest/src/models/LogModel"
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

func (this *UserSetterImpl) AddUser(u *UserModel.UserModelImpl) *result.ErrorResult {
	if dbs.Orm.First(&u, "u_email = ? or u_name = ?", u.Email, u.Name).RecordNotFound() {
		err := dbs.Orm.Transaction(func(tx *gorm.DB) error {
			if err := tx.Save(&u).Error; err != nil {
				return err
			}

			l := LogModel.New("add user", time.Now())
			if err := tx.Save(&l).Error; err != nil {
				return err
			}
			return nil
		})
		if err == nil {
			return &result.ErrorResult{
				Err:  nil,
				Data: u,
			}
		}else{
			return &result.ErrorResult{
				Err:  fmt.Errorf("add user transaction error"),
				Data: nil,
			}
		}
	} else {
		return &result.ErrorResult{
			Err:  fmt.Errorf("user is existed"),
			Data: nil,
		}
	}
}
