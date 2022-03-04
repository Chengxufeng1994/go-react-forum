package auth

import (
	"errors"
	"fmt"
	"github.com/Chengxufeng1994/go-react-forum/global"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
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

func TokenValid(r *http.Request) error {
	var bearerToken string
	authorization := r.Header.Get("Authorization")
	if len(strings.Split(authorization, " ")) == 2 {
		bearerToken = strings.Split(authorization, " ")[1]
	}

	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		jwtSecret := global.GRF_VP.GetString("server.jwt-secret")
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// TODO: pretty claims
		fmt.Println(claims)
	}
	return nil
}
