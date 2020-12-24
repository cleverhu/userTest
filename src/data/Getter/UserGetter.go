package Getter

import (
	"userTest/src/dbs"
	"userTest/src/models/UserModel"
)

var UserGetter IUserGetter

func init() {
	UserGetter = NewUserGetterImpl()
}

type IUserGetter interface {
	GetUserList() []*UserModel.UserInfoImpl
}

type UserGetterImpl struct {
}

func NewUserGetterImpl() *UserGetterImpl {
	return &UserGetterImpl{}
}

func (this *UserGetterImpl) GetUserList() (us []*UserModel.UserInfoImpl) {
	type Info struct {

	}
	dbs.Orm.Table("users").Find(&us)
	return
}
