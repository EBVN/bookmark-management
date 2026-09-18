package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	genpass_mocks "github.com/EBVN/bookmark-management/internal/service/mocks"
)

func TestShortenLink_Redirect(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupMockSvc     func(ctx context.Context) *genpass_mocks.ShortenUrl
		setupTestRequest func(c *gin.Context)

		expectedCode int
		expectedUrl  string
	}{
		{
			name: "normal case",
			setupMockSvc: func(ctx context.Context) *genpass_mocks.ShortenUrl {
				mockSvc := genpass_mocks.NewShortenUrl(t)
				mockSvc.On("GetUrlFromCode", ctx, "1234567").Return("https://google.com", nil)
				return mockSvc
			},

			setupTestRequest: func(c *gin.Context) {
				c.Request = httptest.NewRequest(http.MethodGet, "/v1/links/redirect/1234567", nil)
				//c.Params = gin.Params{{Key: "code", Value: "1234567"}}
			},

			expectedCode: http.StatusMovedPermanently,
			expectedUrl:  "https://google.com",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rec := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(rec)

			// setup
			tc.setupTestRequest(ginCtx)

			mockSvc := tc.setupMockSvc(ginCtx)
			handler := NewShortenLink(mockSvc)

			handler.Redirect(ginCtx)

			// assert
			assert.Equal(t, tc.expectedCode, rec.Code)
			assert.Equal(t, tc.expectedUrl, rec.Header().Get("Location"))
		})
	}
}
