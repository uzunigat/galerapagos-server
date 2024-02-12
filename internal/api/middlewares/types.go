package middlewares

import (
	"github.com/golang-jwt/jwt/v4"
)

type ClientType string

const (
	ClientTypeSystem ClientType = "MACHINE"
	ClientTypeUser              = "USER"
)

type UserMetadata struct {
	UserId string `json:"account_id"`
	Email  string `json:"email"`
}

type ClientMetadata struct {
	ClientName string `json:"client_name"`
}

type CustomClaims struct {
	UserMetadata   UserMetadata   `json:"user_metadata,omitempty"`
	ClientMetadata ClientMetadata `json:"client_metadata,omitempty"`
	ClientType     ClientType     `json:"client_type"`
	Roles          []string       `json:"roles"`
}

type Auth0TokenClaims struct {
	CustomClaims    CustomClaims  `json:"https://hear.com"`
	Permissions     []interface{} `json:"permissions"`
	Scope           string        `json:"scope"`
	AuthorizedParty string        `json:"azp"`
	jwt.RegisteredClaims
}
