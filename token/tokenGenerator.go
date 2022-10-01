package token

import (
	"log"
	"os"
	"time"
	"github.com/dgrijalva/jwt-go"
)

type UserClaim struct {
	Email string
	FirstName string
	LastName string
	UserId string
	IsAdmin bool
	jwt.StandardClaims
}

var SECRET_KEY = os.Getenv("SECRET")

func TokenGenerator(email string, firstName string, lastName string, userId string, isAdmin bool) (accessToken string, refreshToken string, err error) {
	claims := &UserClaim{
		Email: email,
		FirstName: firstName,
		LastName: lastName,
		UserId: userId,
		IsAdmin: isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &UserClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(48)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		return "", "", err
	}

	refreshToken, rerr := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if rerr != nil {
		log.Panic(rerr)
		return
	}

	return token, refreshToken, err
}

func ValidateToken(signedToken string) (claims *UserClaim, msg string) {
	token, err := jwt.ParseWithClaims(signedToken, &UserClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*UserClaim)

	if !ok {
		msg = "the token is invalid"
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "the token is expired"
		return
	}

	return claims, msg
}

