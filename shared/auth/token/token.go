package token

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type JWTTokenVerifier struct {
	PublicKey *rsa.PublicKey
}

func (v *JWTTokenVerifier) Verify(token string) (string, error)  {
	t, err := jwt.ParseWithClaims(
		token,
		&jwt.StandardClaims{},
		func(*jwt.Token) (interface{}, error) {
			return v.PublicKey, nil
		},
	)
	if err != nil {
		return "", fmt.Errorf("cannot parse: %v", err)
	}
	if !t.Valid {
		return "", fmt.Errorf("token not valid")
	}

	clm, ok := t.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", fmt.Errorf("token claim is not standard")
	}
	if clm.Valid() != nil {
		return "", fmt.Errorf("token claim is not valid")
	}

	return clm.Subject, nil

}