package UserModel

import (
	"fmt"
	"time"
)

type UserModelImpl struct {
	ID         int    `gorm:"column:u_id;PRIMAEY_KEY;" json:"id,omitempty"`
	Name       string `binding:"UserName" gorm:"column:u_name" json:"username,omitempty"`
	PassWord   string `binding:"PassWord" gorm:"column:u_password" json:"password,omitempty"`
	Email      string `binding:"Email" gorm:"column:u_email" json:"email,omitempty"`
	UpdateTime MyTime `gorm:"column:u_update_time;type:datetime" json:"update_time,omitempty"`
}

type UserInfoImpl struct {
	ID   int    `gorm:"column:u_id;PRIMAEY_KEY;" json:"id,omitempty"`
	Name string `binding:"UserName" gorm:"column:u_name" json:"username,omitempty"`
	//PassWord   string `binding:"PassWord" gorm:"column:u_password" json:"password,omitempty"`
	Email      string `binding:"Email" gorm:"column:u_email" json:"email,omitempty"`
	UpdateTime MyTime `gorm:"column:u_update_time;type:datetime" json:"update_time,omitempty"`
}

func NewUserInfoImpl() *UserInfoImpl {
	return &UserInfoImpl{}
}

type MyTime time.Time

func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

//func (t MyTime) Value() (driver.Value, error) {
//	// MyTime 转换成 time.Time 类型
//	tTime := time.Time(t)
//	return tTime.Format("2006-01-02 15:04:05"), nil
//}

func New(attrs ...UserModelAttrFunc) *UserModelImpl {
	u := &UserModelImpl{}
	UserModelAttrFuncs(attrs).Apply(u)
	return u
}

func (u *UserModelImpl) Mutate(attrs ...UserModelAttrFunc) *UserModelImpl {
	UserModelAttrFuncs(attrs).Apply(u)
	return u
}

func (UserModelImpl) TableName() string {
	return "users"
}
