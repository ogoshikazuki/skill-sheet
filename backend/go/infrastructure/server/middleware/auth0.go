package middleware

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

type CustomClaims struct {
	Scope string `json:"scope"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func (c CustomClaims) HasScope(expectedScope string) bool {
	result := strings.Split(c.Scope, " ")
	for i := range result {
		if result[i] == expectedScope {
			return true
		}
	}

	return false
}

func Auth0(domain string, audience []string) func(next http.Handler) http.Handler {
	issuerURL, err := url.Parse("https://" + domain + "/")
	if err != nil {
		panic(err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5+time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		audience,
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
	)
	if err != nil {
		panic(err)
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithCredentialsOptional(true),
	)

	return func(next http.Handler) http.Handler {
		return middleware.CheckJWT(next)
	}
}

func IsAuthenticated(ctx context.Context) bool {
	_, ok := ctx.Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
	return ok
}

func HasScope(ctx context.Context, expectedScope string) bool {
	validatedClaims := ctx.Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

	customClaims := validatedClaims.CustomClaims.(*CustomClaims)

	return customClaims.HasScope(expectedScope)
}
