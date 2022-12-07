package utils

import (
	"bank-application/initializers"
	"bank-application/models"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Level    int    `json:"level"`
	jwt.StandardClaims
}

var jwtSecretKey []byte
var jwtExpirationTime int

func init() {
	jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	jwtExpirationTime, _ = strconv.Atoi(os.Getenv("JWT_EXPIRATION_TIME"))
}

func GenerateJWT(user models.User) (tokenString string, err error) {
	var userClaim models.Claim
	initializers.DB.Order("level asc").First(&userClaim, user.ID)

	expirationTime := time.Now().Add(time.Minute * time.Duration(jwtExpirationTime))
	claims := &JWTClaim{
		Email:    user.Email,
		Username: user.Username,
		Level:    userClaim.Level,
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.New().String(),
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtSecretKey)
	return
}

func ValidateToken(signedToken string, level int) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	if claims.Level != level || claims.Level == 0 {
		err = errors.New("unauthorized")
	}

	return
}
