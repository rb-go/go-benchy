package lestrrat_go

import (
	"time"

	"github.com/lestrrat-go/jwx/jwa"

	jwtb "github.com/lestrrat-go/jwx/jwt"
	"github.com/rb-pkg/benchy/jwt"
)

// https://

type JWT struct {
	key []byte
}

func NewJWT(key []byte) *JWT {
	return &JWT{key: key}
}

func (j *JWT) CreateString(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) (string, error) {

	t := jwtb.New()
	err := t.Set(jwtb.JwtIDKey, jti)
	if err != nil {
		return "", err
	}
	err = t.Set(jwtb.IssuerKey, iss)
	if err != nil {
		return "", err
	}
	err = t.Set(jwtb.SubjectKey, sub)
	if err != nil {
		return "", err
	}
	err = t.Set(jwtb.AudienceKey, aud)
	if err != nil {
		return "", err
	}
	err = t.Set(jwtb.IssuedAtKey, time.Now())
	if err != nil {
		return "", err
	}
	err = t.Set(jwtb.ExpirationKey, time.Now().Add(ttl))
	if err != nil {
		return "", err
	}
	err = t.Set(jwtb.NotBeforeKey, time.Now())
	if err != nil {
		return "", err
	}

	err = t.Set("login", payload.Login)
	if err != nil {
		return "", err
	}
	err = t.Set("password", payload.Password)
	if err != nil {
		return "", err
	}
	err = t.Set("is_admin", payload.IsAdmin)
	if err != nil {
		return "", err
	}
	err = t.Set("roles", payload.Roles)
	if err != nil {
		return "", err
	}
	err = t.Set("refresh_token", payload.RefreshToken)
	if err != nil {
		return "", err
	}
	err = t.Set("additional_data", payload.AdditionalData)
	if err != nil {
		return "", err
	}
	err = t.Set("omit_empty_additional_data_filled", payload.OmitEmptyAdditionalDataFilled)
	if err != nil {
		return "", err
	}
	err = t.Set("omit_empty_additional_data_empty", payload.OmitEmptyAdditionalDataEmpty)
	if err != nil {
		return "", err
	}
	err = t.Set("avatar", payload.Avatar)
	if err != nil {
		return "", err
	}

	signed, err := jwtb.Sign(t, jwa.HS256, j.key)
	if err != nil {
		return "", err
	}

	return string(signed), nil
}

func (j *JWT) CreateBytes(jti, iss, sub string, aud []string, ttl time.Duration, payload jwt.PayloadData) ([]byte, error) {

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

	signed, err := jwtb.Sign(t, jwa.HS256, j.key)
	if err != nil {
		return nil, err
	}

	return signed, nil
}

func (j *JWT) ValidateStr(token string) (bool, error) {
	return true, nil
}

func (j *JWT) ValidateBytes(token []byte) (bool, error) {
	return true, nil
}
