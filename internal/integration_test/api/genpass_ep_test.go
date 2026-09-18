package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/EBVN/bookmark-management/internal/api"
	redisPkg "github.com/EBVN/bookmark-management/pkg/redis"
)

func TestGenPassEndpoint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupTestApiEngine func(api api.Engine) *httptest.ResponseRecorder

		expectedStatus   int
		expectedResponse string
	}{
		{
			name: "normal case",

			setupTestApiEngine: func(api api.Engine) *httptest.ResponseRecorder {
				req := httptest.NewRequest(http.MethodGet, "/genpass", nil)
				rec := httptest.NewRecorder()
				api.ServeHTTP(rec, req)
				return rec
			},

			expectedStatus:   http.StatusOK,
			expectedResponse: `{"message":"`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockRedis := redisPkg.InitMockRedis(t)

			testAPI := api.NewEngine(&api.Config{}, mockRedis)
			rec := tc.setupTestApiEngine(testAPI)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), tc.expectedResponse)
		})
	}
}
