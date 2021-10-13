package lestrrat_go

import (
	"crypto/rsa"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	jwtb "github.com/lestrrat-go/jwx/jwt"
	"github.com/riftbit/go-benchy/jwt"
)

// https://

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

func prepareToken(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (jwtb.Token, error) {
	t := jwtb.New()
	err := t.Set(jwtb.JwtIDKey, jti)
	if err != nil {
		return nil, err
	}
	err = t.Set(jwtb.IssuerKey, iss)
	if err != nil {
		return nil, err
	}
	err = t.Set(jwtb.SubjectKey, sub)
	if err != nil {
		return nil, err
	}
	err = t.Set(jwtb.AudienceKey, aud)
	if err != nil {
		return nil, err
	}
	err = t.Set(jwtb.IssuedAtKey, time.Now())
	if err != nil {
		return nil, err
	}
	err = t.Set(jwtb.ExpirationKey, time.Now().Add(ttl))
	if err != nil {
		return nil, err
	}
	err = t.Set(jwtb.NotBeforeKey, time.Now())
	if err != nil {
		return nil, err
	}

	err = t.Set("login", payload.Login)
	if err != nil {
		return nil, err
	}
	err = t.Set("password", payload.Password)
	if err != nil {
		return nil, err
	}
	err = t.Set("is_admin", payload.IsAdmin)
	if err != nil {
		return nil, err
	}
	err = t.Set("roles", payload.Roles)
	if err != nil {
		return nil, err
	}
	err = t.Set("refresh_token", payload.RefreshToken)
	if err != nil {
		return nil, err
	}
	err = t.Set("additional_data", payload.AdditionalData)
	if err != nil {
		return nil, err
	}
	err = t.Set("omit_empty_additional_data_filled", payload.OmitEmptyAdditionalDataFilled)
	if err != nil {
		return nil, err
	}
	err = t.Set("omit_empty_additional_data_empty", payload.OmitEmptyAdditionalDataEmpty)
	if err != nil {
		return nil, err
	}
	err = t.Set("avatar", payload.Avatar)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (j *JWT) CreateStringHS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (string, error) {
	t, err := prepareToken(jti, iss, sub, aud, ttl, payload)
	if err != nil {
		return "", err
	}
	signed, err := jwtb.Sign(t, jwa.HS256, j.key)
	if err != nil {
		return "", err
	}
	return string(signed), nil
}

func (j *JWT) CreateBytesHS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) ([]byte, error) {
	t, err := prepareToken(jti, iss, sub, aud, ttl, payload)
	if err != nil {
		return nil, err
	}
	signed, err := jwtb.Sign(t, jwa.HS256, j.key)
	if err != nil {
		return nil, err
	}
	return signed, nil
}

func (j *JWT) CreateStringRS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (string, error) {
	t, err := prepareToken(jti, iss, sub, aud, ttl, payload)
	if err != nil {
		return "", err
	}
	signed, err := jwtb.Sign(t, jwa.RS256, j.privateKey)
	if err != nil {
		return "", err
	}
	return string(signed), nil
}

func (j *JWT) CreateBytesRS256(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) ([]byte, error) {
	t, err := prepareToken(jti, iss, sub, aud, ttl, payload)
	if err != nil {
		return nil, err
	}
	signed, err := jwtb.Sign(t, jwa.RS256, j.privateKey)
	if err != nil {
		return nil, err
	}
	return signed, nil
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
