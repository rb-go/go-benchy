package gbrlsnchs

import (
	"crypto/rsa"
	"time"

	jwtb "github.com/gbrlsnchs/jwt/v3"
	"github.com/riftbit/go-benchy/jwt"
)

// https://github.com/gbrlsnchs/jwt
// очень неплохой пакет, не сильно медленнее чем cristalhq_jwt,
// разница в аллокациях не большая, нужно смотреть дальше на удобство и функционал

type JWT struct {
	key        *jwtb.HMACSHA
	keyRS      *jwtb.RSASHA
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

type CustomPayload struct {
	jwtb.Payload
	jwt.PayloadData
}

func NewJWT(key []byte, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *JWT {
	return &JWT{
		key:        jwtb.NewHS256(key),
		keyRS:      jwtb.NewRS256(jwtb.RSAPrivateKey(privateKey), jwtb.RSAPublicKey(publicKey)),
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func prepareToken(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) CustomPayload {
	return CustomPayload{
		Payload: jwtb.Payload{
			Issuer:         iss,
			Subject:        sub,
			Audience:       aud,
			ExpirationTime: jwtb.NumericDate(time.Now().Add(ttl)),
			NotBefore:      jwtb.NumericDate(time.Now()),
			IssuedAt:       jwtb.NumericDate(time.Now()),
			JWTID:          jti,
		},
		PayloadData: payload,
	}
}

func (j *JWT) CreateStringHS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (string, error) {
	claims := prepareToken(jti, iss, sub, aud, ttl, payload)
	token, err := jwtb.Sign(claims, j.key)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func (j *JWT) CreateBytesHS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) ([]byte, error) {
	claims := prepareToken(jti, iss, sub, aud, ttl, payload)
	return jwtb.Sign(claims, j.key)
}

func (j *JWT) CreateStringRS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (string, error) {
	claims := prepareToken(jti, iss, sub, aud, ttl, payload)
	token, err := jwtb.Sign(claims, j.keyRS)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

func (j *JWT) CreateBytesRS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) ([]byte, error) {
	claims := prepareToken(jti, iss, sub, aud, ttl, payload)
	return jwtb.Sign(claims, j.keyRS)
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
