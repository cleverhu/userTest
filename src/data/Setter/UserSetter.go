package Setter

import (
	"crypto/md5"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
	"userTest/src/common"
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
	u.Name = strings.TrimSpace(u.Name)
	u.PassWord = strings.TrimSpace(u.PassWord)
	u.Email = strings.TrimSpace(u.Email)

	offset := common.Hashcode(u.Name)
	offset1 := common.Hashcode(u.Email)

	userExists := dbs.Rds.GetBit("users", offset).Val()
	emailExists := dbs.Rds.GetBit("emails", offset1).Val()
	fmt.Println(offset,offset1)
	fmt.Println(userExists,emailExists)
	if userExists == 1 || emailExists == 1 {
		return &result.ErrorResult{
			Err:  fmt.Errorf("user is existed, redis"),
			Data: nil,
		}
	}

	if dbs.Orm.First(&u, "u_email = ? or u_name = ?", u.Email, u.Name).RecordNotFound() {
		err := dbs.Orm.Transaction(func(tx *gorm.DB) error {

			u.PassWord = fmt.Sprintf("%x", md5.Sum([]byte(u.PassWord)))
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
			dbs.Rds.SetBit("users", offset, 1)
			dbs.Rds.SetBit("emails", offset1, 1)
			return &result.ErrorResult{
				Err:  nil,
				Data: u,
			}
		} else {
			return &result.ErrorResult{
				Err:  fmt.Errorf("add user transaction error"),
				Data: nil,
			}
		}
	} else {
		dbs.Rds.SetBit("users", offset, 1)
		dbs.Rds.SetBit("emails", offset1, 1)
		return &result.ErrorResult{
			Err:  fmt.Errorf("user is existed, mysql"),
			Data: nil,
		}
	}
}
