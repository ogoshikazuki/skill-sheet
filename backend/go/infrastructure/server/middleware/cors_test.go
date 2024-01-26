package middleware_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ogoshikazuki/skill-sheet/infrastructure/server/middleware"
)

func TestCors(t *testing.T) {
	// CORS関連のヘッダ、繰り返し使うので定数化
	const (
		headerAccessControlAllowCredentials = "Access-Control-Allow-Credentials"
		headerAccessControlAllowHeaders     = "Access-Control-Allow-Headers"
		headerAccessControlAllowMethods     = "Access-Control-Allow-Methods"
		headerAccessControlAllowOrigin      = "Access-Control-Allow-Origin"
		headerAccessControlExposeHeaders    = "Access-Control-Expose-Headers"
		headerAccessControlRequestHeaders   = "Access-Control-Request-Headers"
		headerAccessControlRequestMethod    = "Access-Control-Request-Method"
		headerOrigin                        = "Origin"
		headerVary                          = "Vary"
	)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	requestPath := "/test"

	tests := map[string]struct {
		allowedOrigins               []string
		origin                       string
		requestMethod                string
		header                       string
		wantAccessControlAllowOrigin string
		isAllowed                    bool
	}{
		"正常系": {
			allowedOrigins:               []string{"http://example.com", "https://example.com"},
			origin:                       "http://example.com",
			requestMethod:                http.MethodPost,
			wantAccessControlAllowOrigin: "http://example.com",
			isAllowed:                    true,
		},
		"異常系:オリジンが許可されていない": {
			allowedOrigins:               []string{"http://example.com"},
			origin:                       "https://example.com",
			requestMethod:                http.MethodPost,
			wantAccessControlAllowOrigin: "",
			isAllowed:                    false,
		},
		"異常系:メソッドが許可されていない": {
			allowedOrigins:               []string{"http://example.com"},
			origin:                       "http://example.com",
			requestMethod:                http.MethodGet,
			wantAccessControlAllowOrigin: "",
			isAllowed:                    false,
		},
		"異常系:ヘッダが許可されていない": {
			allowedOrigins:               []string{"http://example.com"},
			origin:                       "http://example.com",
			requestMethod:                http.MethodPost,
			header:                       "Not-Allowed-Header",
			wantAccessControlAllowOrigin: "",
			isAllowed:                    false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			corsMiddleware := middleware.Cors(middleware.WithCorsAllowedOrigins(tt.allowedOrigins))
			handler := corsMiddleware(handler)

			// プリフライトリクエスト
			preflightReq, err := http.NewRequestWithContext(context.TODO(), http.MethodOptions, requestPath, http.NoBody)
			if err != nil {
				t.Fatal(err)
			}
			if tt.header != "" {
				preflightReq.Header.Add(headerAccessControlRequestHeaders, tt.header)
			}
			preflightReq.Header.Add(headerAccessControlRequestMethod, tt.requestMethod)
			preflightReq.Header.Add(headerOrigin, tt.origin)

			// プリフライトリクエストのテスト
			preFlightRecord := httptest.NewRecorder()
			handler.ServeHTTP(preFlightRecord, preflightReq)
			if preFlightRecord.Code != http.StatusNoContent {
				t.Errorf("handler returned wrong status code: got %v want %v", preFlightRecord.Code, http.StatusNoContent)
			}

			assertHeader(t, preFlightRecord.Header(), headerAccessControlAllowOrigin, []string{tt.wantAccessControlAllowOrigin})
			if tt.isAllowed {
				assertHeader(t, preFlightRecord.Header(), headerAccessControlAllowCredentials, []string{"true"})
				assertHeader(t, preFlightRecord.Header(), headerAccessControlAllowMethods, []string{http.MethodPost})
				assertHeader(t, preFlightRecord.Header(), headerVary, []string{
					headerOrigin,
					headerAccessControlRequestMethod,
				})
			} else {
				assertHeader(t, preFlightRecord.Header(), headerAccessControlAllowCredentials, nil)
				assertHeader(t, preFlightRecord.Header(), headerAccessControlAllowMethods, nil)
				assertHeader(t, preFlightRecord.Header(), headerVary, []string{headerOrigin})
			}

			if tt.isAllowed {
				// 実際のリクエスト
				req, err := http.NewRequestWithContext(context.TODO(), tt.requestMethod, requestPath, http.NoBody)
				if err != nil {
					t.Fatal(err)
				}
				req.Header.Add(headerOrigin, tt.origin)

				// 実際のリクエストのテスト
				record := httptest.NewRecorder()
				handler.ServeHTTP(record, req)

				assertHeader(t, record.Header(), headerAccessControlAllowOrigin, []string{tt.wantAccessControlAllowOrigin})
				if tt.isAllowed {
					assertHeader(t, record.Header(), headerAccessControlAllowCredentials, []string{"true"})
					assertHeader(t, record.Header(), headerVary, []string{headerOrigin})
				} else {
					assertHeader(t, record.Header(), headerAccessControlAllowCredentials, nil)
					assertHeader(t, record.Header(), headerVary, []string{headerOrigin})
				}
			}
		})
	}
}

func assertHeader(t *testing.T, header http.Header, key string, expectedValues []string) {
	if expectedValues == nil {
		if header.Get(key) != "" {
			t.Errorf("unexpected header %s: got %v want %v", key, header.Get(key), expectedValues)
		}
		return
	}

	values := strings.Split(header.Get(key), ", ")
	valueSet := make(map[string]struct{})
	for _, value := range values {
		valueSet[value] = struct{}{}
	}
	for _, expectedValue := range expectedValues {
		if _, ok := valueSet[expectedValue]; !ok {
			t.Errorf("unexpected header %s: got %v want %v", key, values, expectedValues)
		}
	}
}
