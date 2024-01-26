package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

type corsOptions struct {
	allowedOrigins []string
}

type CorsOption func(*corsOptions)

func WithCorsAllowedOrigins(allowedOrigins []string) CorsOption {
	return func(opts *corsOptions) {
		opts.allowedOrigins = allowedOrigins
	}
}

func Cors(opts ...CorsOption) func(http.Handler) http.Handler {
	options := corsOptions{}
	for _, opt := range opts {
		opt(&options)
	}

	return func(h http.Handler) http.Handler {
		return cors.New(cors.Options{
			AllowedMethods: []string{
				http.MethodPost,
			},
			AllowedOrigins:   options.allowedOrigins,
			AllowCredentials: true,
		}).Handler(h)
	}
}
