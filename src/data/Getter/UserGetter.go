package Getter

import (
	"fmt"
	"userTest/src/common"
	"userTest/src/dbs"
	"userTest/src/models/UserModel"
	"userTest/src/result"
)

var UserGetter IUserGetter

func init() {
	UserGetter = NewUserGetterImpl()
}

type IUserGetter interface {
	GetUserList() []*UserModel.UserReturnInfoImpl
	Login(u *UserModel.UserLoginInfoImpl) *result.ErrorResult
}

type UserGetterImpl struct {
}

func NewUserGetterImpl() *UserGetterImpl {
	return &UserGetterImpl{}
}

func (this *UserGetterImpl) GetUserList() (us []*UserModel.UserReturnInfoImpl) {
	dbs.Orm.Table("t_user").Find(&us)
	return
}

func (this *UserGetterImpl) Login(u *UserModel.UserLoginInfoImpl) *result.ErrorResult {

	if dbs.Orm.Find(&u, "u_name = ? or u_email = ?", u.Name, u.Name).RecordNotFound() {
		return &result.ErrorResult{
			Err:  fmt.Errorf("user not found"),
			Data: nil,
		}
	} else {
		if dbs.Orm.Table("t_user").Find(&u, "u_name = ? or u_email = ? and u_password = ?", u.Name, u.Name, common.MD5(u.PassWord)).RowsAffected == 1 {
			if token, err := common.CreateToken(u); err != nil {
				return &result.ErrorResult{
					Err:  fmt.Errorf("generate token error"),
					Data: nil,
				}
			} else {
				return &result.ErrorResult{
					Err:  nil,
					Data: token,
				}
			}
		} else {
			return &result.ErrorResult{
				Err:  fmt.Errorf("user password is wrong"),
				Data: nil,
			}
		}

	}
}
