package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"userTest/src/models/UserModel"
)

const (
	SECRET = "deeplyThink"
)

func CreateToken(u *UserModel.UserLoginInfoImpl) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":      u.ID,
		"username": u.Name,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	token, err := at.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	}
	return token, nil
}
