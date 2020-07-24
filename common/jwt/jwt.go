package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const Secret = "kih**&hgyshq##js"

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

type UserInfo struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func (user *UserInfo) GenerateToken() (string, error) {
	claim := jwt.MapClaims{
		"id":   user.Id,
		"name": user.Name,
		"nbf":  time.Now().Unix(),
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Unix() + 3*60*60,
		"iss":  "myxy99.cn",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokens, err := token.SignedString([]byte(Secret))
	return tokens, err
}

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	}
}

func (user *UserInfo) ParseToken(tokens string) (err error) {
	token, err := jwt.Parse(tokens, secret())
	if err != nil {
		return
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to mapclaim")
		return
	}
	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}
	user.Name = claim["name"].(string)
	user.Id = uint(claim["id"].(float64))
	return err
}
