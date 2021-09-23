package jwt

import (
	"encoding/json"
	"time"
)

type JWTer interface {
	CreateStringHS256(jti, iss, sub string, aud []string, ttl time.Duration, payload PayloadData) (string, error)
	CreateStringRS256(jti, iss, sub string, aud []string, ttl time.Duration, payload PayloadData) (string, error)
	CreateBytesHS256(jti, iss, sub string, aud []string, ttl time.Duration, payload PayloadData) ([]byte, error)
	CreateBytesRS256(jti, iss, sub string, aud []string, ttl time.Duration, payload PayloadData) ([]byte, error)
	ValidateStrHS256(token string) (bool, error)
	ValidateStrRS256(token string) (bool, error)
	ValidateBytesHS256(token []byte) (bool, error)
	ValidateBytesRS256(token []byte) (bool, error)
}

type PayloadData struct {
	Login                         string          `json:"login"`
	Password                      string          `json:"password"`
	IsAdmin                       bool            `json:"is_admin"`
	Roles                         []string        `json:"roles"`
	RefreshToken                  []byte          `json:"refresh_token"`
	AdditionalData                json.RawMessage `json:"additional_data"`
	OmitEmptyAdditionalDataFilled json.RawMessage `json:"omit_empty_additional_data_filled,omitempty"`
	OmitEmptyAdditionalDataEmpty  json.RawMessage `json:"omit_empty_additional_data_empty,omitempty"`
	Avatar                        []byte          `json:"avatar"`
}
