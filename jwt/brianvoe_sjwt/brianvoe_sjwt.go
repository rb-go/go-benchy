package brianvoe

import (
	"crypto/rsa"
	"time"

	sjwt "github.com/brianvoe/sjwt"
	"github.com/rb-pkg/benchy/jwt"
)

// https://github.com/brianvoe/sjwt
// Унылое говно!!! не поддерживает другие типы шифрования, работает пздц медленно,
// работает только со строками в качестве токена

type JWT struct {
	key        []byte
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewJWT(key []byte, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *JWT {
	return &JWT{
		key:        key,
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func prepareToken(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (sjwt.Claims, error) {
	// Set Claims
	claims, err := sjwt.ToClaims(payload)
	if err != nil {
		return nil, err
	}
	claims.Set(sjwt.TokenID, jti)            // UUID generated
	claims.SetIssuer(iss)                    // Issuer of the token
	claims.SetSubject(sub)                   // Subject of the token
	claims.SetIssuedAt(time.Now())           // IssuedAt in time, value is set in unix
	claims.SetNotBeforeAt(time.Now())        // Token valid in 1 hour
	claims.SetExpiresAt(time.Now().Add(ttl)) // Token expires in 24 hours
	claims.SetAudience(aud)                  // Audience the toke is for
	return claims, err
}

func (j *JWT) CreateStringHS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (string, error) {
	claims, err := prepareToken(jti, iss, sub, aud, ttl, payload)
	if err != nil {
		return "", err
	}
	// Generate jwt
	return claims.Generate(j.key), nil
}

func (j *JWT) CreateBytesHS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) ([]byte, error) {
	claims, err := prepareToken(jti, iss, sub, aud, ttl, payload)
	if err != nil {
		return nil, err
	}
	// Generate jwt
	return []byte(claims.Generate(j.key)), nil
}

// CreateStringRS256 NOT SUPPORTED!!!
func (j *JWT) CreateStringRS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (string, error) {
	time.Sleep(time.Second * 2)
	return j.CreateStringHS256(jti, iss, sub, aud, ttl, payload)
}

// CreateBytesRS256 NOT SUPPORTED!!!
func (j *JWT) CreateBytesRS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) ([]byte, error) {
	time.Sleep(time.Second * 2)
	return j.CreateBytesHS256(jti, iss, sub, aud, ttl, payload)
}

func (j *JWT) ValidateStrHS256(token string) (bool, error) {
	return true, nil
}

func (j *JWT) ValidateBytesHS256(token []byte) (bool, error) {
	return true, nil
}

// ValidateStrRS256 NOT SUPPORTED!!!
func (j *JWT) ValidateStrRS256(token string) (bool, error) {
	time.Sleep(time.Second * 2)
	return j.ValidateStrHS256(token)
}

// ValidateBytesRS256 NOT SUPPORTED!!!
func (j *JWT) ValidateBytesRS256(token []byte) (bool, error) {
	time.Sleep(time.Second * 2)
	return j.ValidateBytesHS256(token)
}
