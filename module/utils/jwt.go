package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pivaros/go-image-recognition/constants"
)

type UserClaims struct {
	UserId string
	Role   constants.Role
	jwt.StandardClaims
}
