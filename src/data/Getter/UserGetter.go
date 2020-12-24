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
	GetUserList() []*UserModel.UserModelImpl
}

type UserGetterImpl struct {
}

func NewUserGetterImpl() *UserGetterImpl {
	return &UserGetterImpl{}
}

func (this *UserGetterImpl) GetUserList() (us []*UserModel.UserModelImpl) {
	dbs.Orm.Find(&us)
	return
}
