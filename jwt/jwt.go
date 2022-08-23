package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JwtWrapper wraps the signing key and the issuer
type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

// JwtClaim adds username as a claim to the token
type JwtClaim struct {
	Id       uint
	Username string
	jwt.MapClaims
}

// GenerateToken generates a jwt token
func (j *JwtWrapper) GenerateToken(userId uint, username string) (signedToken string, err error) {
	// create the claims
	claims := &JwtClaim{
		Id:       userId,
		Username: username,
		MapClaims: jwt.MapClaims{
			"exp": time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			"iss": j.Issuer,
		},
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// generate encoded token and send it as response
	signedToken, err = token.SignedString([]byte(j.SecretKey))

	if err != nil {
		return
	}

	return
}

// ValidateToken validates the jwt token
func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("Couldn't parse claims")
		return
	}

	if !claims.VerifyExpiresAt(time.Now().Local().Unix(), true) {
		err = errors.New("JWT is expired")
		return
	}

	return
}