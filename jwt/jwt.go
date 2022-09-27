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
	IssueAt         int64
	ExpirationHours int64
}

// JwtClaim adds username and user id as a claim to the token
type JwtClaim struct {
	Id            uint   `json:"id"`
	Username      string `json:"username"`
	jwt.MapClaims `json:"details"`
}

// FileJwtClaim adds ownerkey as a claim to the token
type FileJwtClaim struct {
	OwnerKey      string `json:"owner_key"`
	jwt.MapClaims `json:"details"`
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
			"iat": j.IssueAt,
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

// FileGenerateSharedToken generates a jwt token for shared files
func (j *JwtWrapper) FileGenerateSharedToken(ownerKey string) (signedToken string, err error) {
	// create the claims
	var claims *FileJwtClaim
	if j.ExpirationHours == 0 {
		claims = &FileJwtClaim{
			OwnerKey: ownerKey,
			MapClaims: jwt.MapClaims{
				"exp": nil,
				"iss": j.Issuer,
				"iat": j.IssueAt,
			},
		}
	} else {
		claims = &FileJwtClaim{
			OwnerKey: ownerKey,
			MapClaims: jwt.MapClaims{
				"exp": time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
				"iss": j.Issuer,
				"iat": j.IssueAt,
			},
		}
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
