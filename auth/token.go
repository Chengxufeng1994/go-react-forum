package auth

import (
	"github.com/Chengxufeng1994/go-react-forum/global"
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserId uint32
	jwt.StandardClaims
}

func CreateToken(userId uint32) (string, error) {
	claims := &Token{UserId: userId}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := global.GRF_VP.GetString("server.jwt-secret")
	return token.SignedString([]byte(jwtSecret))
}
