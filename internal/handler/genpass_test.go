package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	genpass_mocks "github.com/EBVN/bookmark-management/internal/service/mocks"
)

var testErr = errors.New("test error")

func TestGenPass_GeneratePassword(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupRequest     func(c *gin.Context)
		setupMockService func(ctx context.Context) *genpass_mocks.GenPass

		expectedStatus   int
		expectedResponse string
	}{
		{
			name: "normal case",

			setupRequest: func(c *gin.Context) {
				c.Request = httptest.NewRequest(http.MethodGet, "/genpass", nil)
			},

			setupMockService: func(ctx context.Context) *genpass_mocks.GenPass {
				mockSvc := genpass_mocks.NewGenPass(t)
				mockSvc.On("GeneratePassword", DefaultPasswordLength).Return("1234567890", nil)
				return mockSvc
			},

			expectedStatus:   http.StatusOK,
			expectedResponse: `{"message":"1234567890"}`,
		},
		{
			name: "err case",

			setupRequest: func(c *gin.Context) {
				c.Request = httptest.NewRequest(http.MethodGet, "/genpass", nil)
			},

			setupMockService: func(ctx context.Context) *genpass_mocks.GenPass {
				mockSvc := genpass_mocks.NewGenPass(t)
				mockSvc.On("GeneratePassword", DefaultPasswordLength).Return("", testErr)
				return mockSvc
			},

			expectedStatus:   http.StatusInternalServerError,
			expectedResponse: `{"error":"Internal Server Error"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rec := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(rec)

			// setup
			tc.setupRequest(ginCtx)

			// setup mock
			mockSvc := tc.setupMockService(ginCtx)

			// setup handler
			genPassHandler := NewGenPass(mockSvc)

			// run test
			genPassHandler.GeneratePassword(ginCtx)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.Equal(t, tc.expectedResponse, rec.Body.String())

		})
	}

}
