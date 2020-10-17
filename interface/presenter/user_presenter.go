package presenter

import (
	"os"
	"service-back/consts"
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().AddDate(0, 0, 7).Unix(),
		"iat": time.Now().Unix(),
	})
	tokenString, _ := token.SignedString([]byte(os.Getenv(consts.SIGNIN_KEY)))
	oUser := outputdata.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	return &outputdata.Login{
		JWT:  tokenString,
		User: oUser,
	}
}
