package UserModel

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

type UserModelImpl struct {
	ID         int    `gorm:"column:u_id;PRIMAEY_KEY;" json:"id,omitempty"`
	Name       string `binding:"UserName" gorm:"column:u_name" json:"username,omitempty"`
	PassWord   string `binding:"PassWord" gorm:"column:u_password" json:"password,omitempty"`
	Email      string `binding:"Email" gorm:"column:u_email" json:"email,omitempty"`
	UpdateTime MyTime `gorm:"column:u_update_time;type:datetime" json:"update_time,omitempty"`
}

type UserReturnInfoImpl struct {
	ID   int    `gorm:"column:u_id;PRIMAEY_KEY;" json:"id,omitempty"`
	Name string `binding:"UserName" gorm:"column:u_name" json:"username,omitempty"`
	//PassWord   string `binding:"PassWord" gorm:"column:u_password" json:"password,omitempty"`
	Email      string `binding:"Email" gorm:"column:u_email" json:"email,omitempty"`
	UpdateTime MyTime `gorm:"column:u_update_time;type:datetime" json:"update_time,omitempty"`
}

type UserLoginInfoImpl struct {
	ID         int    `gorm:"column:u_id;PRIMAEY_KEY;" json:"id,omitempty"`
	Name       string `binding:"LoginUser" gorm:"column:u_name" json:"username,omitempty"`
	PassWord   string `binding:"PassWord" gorm:"column:u_password" json:"password,omitempty"`
	UpdateTime MyTime `gorm:"column:u_update_time;type:datetime" json:"update_time,omitempty"`
}

func NewUserLoginInfoImpl() *UserLoginInfoImpl {
	return &UserLoginInfoImpl{}
}

func NewUserInfoImpl() *UserReturnInfoImpl {
	return &UserReturnInfoImpl{}
}

type MyTime time.Time

func (t *MyTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = MyTime(t1)
	return err
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t MyTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}

func (t *MyTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = MyTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *MyTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
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
	return "t_user"
}

//UserReturnInfoImpl UserLoginInfoImpl

func (UserReturnInfoImpl) TableName() string {
	return "t_user"
}

func (UserLoginInfoImpl) TableName() string {
	return "t_user"
}