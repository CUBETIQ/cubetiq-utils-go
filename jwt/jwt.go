package jwt

import (
	"errors"
	"fmt"
	"strings"
	"time"

	stringutil "github.com/cubetiq/cubetiq-utils-go/string"
	"github.com/golang-jwt/jwt/v5"
)

const (
	AUTHORIZATION_HEADER = "Authorization"
	TOKEN_PREFIX         = "bearer"
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
	ID            string `json:"id"`
	Username      string `json:"username"`
	jwt.MapClaims `json:"details"`
}

// UsernameJwtClaim adds username as a claim to the token
type UsernameJwtClaim struct {
	Username      string `json:"username"`
	jwt.MapClaims `json:"details"`
}

// FileJwtClaim adds ownerkey as a claim to the token
type FileJwtClaim struct {
	OwnerKey      string   `json:"owner_key"`
	FileIds       []string `json:"file_ids"`
	jwt.MapClaims `json:"details"`
}

// EncryptToken generates a jwt token
func (j *JwtWrapper) EncryptToken(userId string, username string) (signedToken string, err error) {
	// create the claims
	claims := &JwtClaim{
		ID:       userId,
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

// EncryptTokenByUsername generates a jwt token that takes only username
func (j *JwtWrapper) EncryptTokenByUsername(username string) (signedToken string, err error) {
	// create the claims
	claims := &UsernameJwtClaim{
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

// FileEncryptSharedToken generates a jwt token for shared files
func (j *JwtWrapper) FileEncryptSharedToken(ownerKey string, fileIds []string) (signedToken string, err error) {
	// create the claims
	var claims *FileJwtClaim
	if j.ExpirationHours == 0 {
		claims = &FileJwtClaim{
			OwnerKey: ownerKey,
			FileIds:  fileIds,
			MapClaims: jwt.MapClaims{
				"exp": nil,
				"iss": j.Issuer,
				"iat": j.IssueAt,
			},
		}
	} else {
		claims = &FileJwtClaim{
			OwnerKey: ownerKey,
			FileIds:  fileIds,
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

func (j *JwtWrapper) DecryptToken(tokenString string) (*JwtClaim, error) {
	return DecryptToken(tokenString, []byte(j.SecretKey))
}

// DecryptToken get claims from token
func DecryptToken(tokenString string, secretKey []byte) (*JwtClaim, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return &JwtClaim{
		ID:       claim["id"].(string),
		Username: claim["username"].(string),
	}, nil
}

// ExtractToken get token without Bearer or bearer
func ExtractToken(token string) (string, error) {
	// if token is empty then send error
	if stringutil.IsEmpty(token) {
		return "", errors.New("token is required")
	}

	// get split token
	getSplitToken := strings.Split(token, " ")
	// get token
	getBearer := stringutil.ToLower(getSplitToken[0])
	if getBearer != TOKEN_PREFIX {
		return "", errors.New("bearer is required")
	}

	// get token without Bearer
	getToken := getSplitToken[1]

	// validate again with a token that has three dots or not if not then send error
	if len(strings.Split(getToken, ".")) != 3 {
		return "", errors.New("token is invalid")
	}

	return getToken, nil
}
