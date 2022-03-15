package token

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JWTTokenGen 产生一个JWT token
type JWTTokenGen struct {
	issuer 	string
	nowFunc func() time.Time
	key 	*rsa.PrivateKey
}

// NewJWTTokenGen 是JWTTokenGen对应的构造方法
func NewJWTTokenGen(issuer string, key *rsa.PrivateKey) *JWTTokenGen {
	return &JWTTokenGen{
		issuer:	 	issuer,
		nowFunc : 	time.Now,
		key: 		key,
	}
}

// GenerateToken 由accountID和过期时间expire得到Token
func (t *JWTTokenGen) GenerateToken(accountID string, expire time.Duration) (string, error) {
	nowSec := t.nowFunc().Unix()
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		Issuer: t.issuer,
		IssuedAt: nowSec,
		ExpiresAt: nowSec + int64(expire.Seconds()),
		Subject: accountID,
	})

	return tkn.SignedString(t.key)
}