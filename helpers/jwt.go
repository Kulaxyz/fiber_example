package helpers

import (
	"github.com/Kulaxyz/partner/config"
	"github.com/Kulaxyz/partner/models"
	"github.com/dgrijalva/jwt-go"
)

func NewToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["user_id"] = user.ID

	return token.SignedString([]byte(config.Config("JWT_SECRET")))
}
