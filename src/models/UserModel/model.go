package UserModel

import "time"

type UserModelImpl struct {
	ID         int       `gorm:"column:u_id;PRIMAEY_KEY;" json:"id,omitempty"`
	Name       string    `binding:"UserName" gorm:"column:u_name" json:"username,omitempty"`
	PassWord   string    `binding:"PassWord" gorm:"column:u_password" json:"password,omitempty"`
	Email      string    `binding:"Email" gorm:"column:u_email" json:"email,omitempty"`
	UpdateTime time.Time `gorm:"column:u_update_time;type:datetime" json:"update_time,omitempty"`
}

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
