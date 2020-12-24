package UserModel

import "time"

//func New(attrs ...UserModelAttrFunc) *UserModelImpl {
//	u := &UserModelImpl{}
//	UserModelAttrFuncs(attrs).Apply(u)
//	return u
//}
//
//func (u *UserModelImpl) Mutate(attrs ...UserModelAttrFunc) *UserModelImpl {
//	UserModelAttrFuncs(attrs).Apply(u)
//	return u
//}
func init() {

}

type UserModelAttrFunc func(*UserModelImpl)

type UserModelAttrFuncs []UserModelAttrFunc

func (this UserModelAttrFuncs) Apply(u *UserModelImpl) {
	for _, f := range this {
		f(u)
	}
}

func WithUserID(id int) UserModelAttrFunc {
	return func(impl *UserModelImpl) {
		impl.ID = id
	}
}

func WithUserName(name string) UserModelAttrFunc {
	return func(impl *UserModelImpl) {
		impl.Name = name
	}
}

func WithUserPassWord(password string) UserModelAttrFunc {
	return func(impl *UserModelImpl) {
		impl.PassWord = password
	}
}

func WithUserEmail(email string) UserModelAttrFunc {
	return func(impl *UserModelImpl) {
		impl.Email = email
	}
}

func WithUpdateTime(updateTime time.Time) UserModelAttrFunc {
	return func(impl *UserModelImpl) {
		impl.UpdateTime = updateTime
	}
}


