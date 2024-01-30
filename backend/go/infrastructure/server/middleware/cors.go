package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

type corsOptions struct {
	allowedOrigins []string
}

type CORSOption func(*corsOptions)

func WithCORSAllowedOrigins(allowedOrigins []string) CORSOption {
	return func(opts *corsOptions) {
		opts.allowedOrigins = allowedOrigins
	}
}

func CORS(opts ...CORSOption) func(http.Handler) http.Handler {
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
