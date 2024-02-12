package middlewares

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/api"
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/internal/domain/model"
	"github.com/MicahParks/keyfunc"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/zerolog/log"
)

const (
	CustomClaimsTag = "customClaims"
	PermissionsTag  = "permissions"
)

func NewJWKSProvider(jwksBaseURL string) *keyfunc.JWKS {

	jwksURL, err := url.Parse(fmt.Sprintf("%s/.well-known/jwks.json", jwksBaseURL))

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse Auth0 jwks URL.")
	}

	// Create the keyfunc options. Refresh the JWKS every hour and log errors.

	options := keyfunc.Options{
		RefreshInterval: time.Hour,
		RefreshErrorHandler: func(err error) {
			log.Fatal().Err(err).Msg("Failed to refresh JWKS.")
		},
	}

	// Create the JWKS from the resource at the given URL.
	jwks, err := keyfunc.Get(jwksURL.String(), options)
	if err != nil {
		log.Fatal().Err(err).Str("url", jwksURL.String()).Msg("Failed to create JWKS from resource at the given URL.")
	}
	return jwks
}

func Authenticate(provider *keyfunc.JWKS) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authorizationHeader := ctx.Request.Header.Get("Authorization")

		if authorizationHeader == "" {
			error := api.NewUnauthorizedError(fmt.Errorf("No token provided."))
			ctx.AbortWithStatusJSON(error.HTTPStatus, error)
			return
		}

		jwtToken, err := stripBearerPrefixFromTokenString(authorizationHeader)

		if err != nil {
			error := api.NewUnauthorizedError(err)
			ctx.AbortWithStatusJSON(error.HTTPStatus, error)
			return
		}

		claims := new(Auth0TokenClaims)

		token, err := jwt.ParseWithClaims(jwtToken, claims, provider.Keyfunc)

		if err != nil {
			error := api.NewUnauthorizedError(fmt.Errorf("Unable to parse token: %v.", err))
			ctx.AbortWithStatusJSON(error.HTTPStatus, error)
			return
		}

		if !token.Valid {
			error := api.NewUnauthorizedError(fmt.Errorf("Invalid token."))
			ctx.AbortWithStatusJSON(error.HTTPStatus, error)
			return
		}

		ctx.Set(PermissionsTag, claims.Permissions)
		ctx.Set(CustomClaimsTag, claims.CustomClaims)
		ctx.Set(model.ContextTagToken, authorizationHeader)

		SetUserContext(ctx)
		ctx.Next()
	}
}

func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	}
	return "", fmt.Errorf("Invalid token.")
}
