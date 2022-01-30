package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GeneraleToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issure    string
}

func JWTAuthService() *jwtService {
	return &jwtService{
		secretKey: getSecretKey(),
		issure:    "jwt",
	}
}

func getSecretKey() string {
	secret := os.Getenv("secret")
	if secret == "" {
		secret = "secret"
	}

	return secret
}

func (service *jwtService) GeneraleToken(email string, isUser bool) string {
	claims := &authCustomClaims{
		email,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    service.issure,
			IssuedAt:  jwt.TimeFunc().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}

	return t
}

func (service *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %s", token.Header["alg"])
		}

		return []byte(service.secretKey), nil
	})
}
