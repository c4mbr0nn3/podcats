package auth

import (
	"errors"
	"example/hello/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = utils.GetJwtSecret()

type JwtClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(id int, username string, email string) (string, error) {
	issuedTime := time.Now()
	claims := JwtClaims{
		id,
		username,
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(issuedTime.Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(issuedTime),
			NotBefore: jwt.NewNumericDate(issuedTime),
			Issuer:    "podcats",
			Subject:   "auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJwtToken(authToken string) error {
	token, err := jwt.ParseWithClaims(
		getTokenFromHeader(authToken),
		&JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)

	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		err = errors.New("invalid token")
		return err
	}
	if claims.ExpiresAt.Unix() < time.Now().Unix() {
		err = errors.New("token expired")
		return err
	}
	return nil
}

// get jwt token from gin context and map it to jwt claims
func GetJwtClaims(c *gin.Context) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(
		getTokenFromHeader(c.GetHeader("Authorization")),
		&JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		err = errors.New("invalid token")
		return nil, err
	}
	return claims, nil
}

func getTokenFromHeader(header string) string {
	return header[7:]
}
