package presenter

import (
	"service-back/domain/model"
	outputdata "service-back/usecase/output/data"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserPresenter struct {
}

func NewUserPresenter() *UserPresenter {
	return &UserPresenter{}
}

func (p *UserPresenter) Login(user *model.User) *outputdata.Login {
	jwtToken := jwt.New(jwt.SigningMethodHS256)
	claims := jwtToken.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["exp"] = time.Now().AddDate(0, 0, 7).Unix()
	claims["iat"] = time.Now().Unix()
	t, _ := jwtToken.SignedString([]byte("secret"))
	return &outputdata.Login{
		JwtToken: t,
	}
}
