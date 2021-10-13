package cristalhq

import (
	"crypto/rsa"
	"time"

	jwtb "github.com/cristalhq/jwt/v3"
	"github.com/riftbit/go-benchy/jwt"
)

// https://github.com/cristalhq/jwt

type JWT struct {
	key        []byte
	builder    *jwtb.Builder
	signer     jwtb.Signer
	builderRS  *jwtb.Builder
	signerRS   jwtb.Signer
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

type UserClaims struct {
	jwtb.RegisteredClaims
	jwt.PayloadData
}

func NewJWT(key []byte, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *JWT {
	signer, _ := jwtb.NewSignerHS(jwtb.HS256, key)
	signerRS, _ := jwtb.NewSignerRS(jwtb.RS256, privateKey)
	res := &JWT{
		key:        key,
		builder:    jwtb.NewBuilder(signer),
		signer:     signer,
		builderRS:  jwtb.NewBuilder(signerRS),
		signerRS:   signerRS,
		privateKey: privateKey,
		publicKey:  publicKey,
	}
	return res
}

func prepareToken(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) UserClaims {
	return UserClaims{
		RegisteredClaims: jwtb.RegisteredClaims{
			ID:        jti,
			Issuer:    iss,
			Subject:   sub,
			Audience:  aud,
			ExpiresAt: jwtb.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwtb.NewNumericDate(time.Now()),
			NotBefore: jwtb.NewNumericDate(time.Now()),
		},
		PayloadData: payload,
	}
}

func (j *JWT) CreateStringHS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (string, error) {
	claims := prepareToken(jti, iss, sub, aud, ttl, payload)
	token, err := j.builder.Build(claims)
	if err != nil {
		return "", err
	}
	return token.String(), nil
}

func (j *JWT) CreateBytesHS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) ([]byte, error) {
	claims := prepareToken(jti, iss, sub, aud, ttl, payload)
	token, err := j.builder.Build(claims)
	if err != nil {
		return nil, err
	}
	return token.Raw(), nil
}

func (j *JWT) CreateStringRS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (string, error) {
	claims := prepareToken(jti, iss, sub, aud, ttl, payload)
	token, err := j.builderRS.Build(claims)
	if err != nil {
		return "", err
	}
	return token.String(), nil
}

func (j *JWT) CreateBytesRS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) ([]byte, error) {
	claims := prepareToken(jti, iss, sub, aud, ttl, payload)
	token, err := j.builderRS.Build(claims)
	if err != nil {
		return nil, err
	}
	return token.Raw(), nil
}

func (j *JWT) ValidateStrHS256(token string) (bool, error) {
	return true, nil
}

func (j *JWT) ValidateBytesHS256(token []byte) (bool, error) {
	return true, nil
}

func (j *JWT) ValidateStrRS256(token string) (bool, error) {
	return true, nil
}

func (j *JWT) ValidateBytesRS256(token []byte) (bool, error) {
	return true, nil
}
