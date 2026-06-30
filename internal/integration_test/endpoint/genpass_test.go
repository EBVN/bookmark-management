package endpoint

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/EBVN/golang-course/internal/api"
)

func TestGenPass(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupTestHTTP func(api api.Engine) *httptest.ResponseRecorder

		expectedStatusCode      int
		expectedResponseContain string
	}{
		{
			name: "Generate password successfully",

			setupTestHTTP: func(api api.Engine) *httptest.ResponseRecorder {
				req := httptest.NewRequest("GET", "/genpass", nil)
				respRec := httptest.NewRecorder()
				api.ServeHTTP(respRec, req)
				return respRec
			},
			expectedStatusCode:      http.StatusOK,
			expectedResponseContain: `{"password":`,
		},
		{
			name: "Wrong endpoint",

			setupTestHTTP: func(api api.Engine) *httptest.ResponseRecorder {
				req := httptest.NewRequest("GET", "/genpass2", nil)
				respRec := httptest.NewRecorder()
				api.ServeHTTP(respRec, req)
				return respRec
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			apiEngine := api.NewEngine()

			rec := tc.setupTestHTTP(apiEngine)

			assert.Equal(t, tc.expectedStatusCode, rec.Code)
			if len(tc.expectedResponseContain) > 0 {
				assert.Equal(t, len(rec.Body.String()), len(tc.expectedResponseContain)+12+3)
			}
			assert.Contains(t, rec.Body.String(), tc.expectedResponseContain)
		})
	}

}
