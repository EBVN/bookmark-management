package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	genpass_svc_mocks "github.com/EBVN/bookmark-management/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePassword(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name         string
		setupMockSvc func(ctx context.Context) *genpass_svc_mocks.GenPass

		expectedStatus   int
		expectedResponse string
	}{
		{
			name: "normal case",
			setupMockSvc: func(ctx context.Context) *genpass_svc_mocks.GenPass {
				mockSvc := genpass_svc_mocks.NewGenPass(t)
				mockSvc.On("GeneratePassword", DefaultLength).Return("password", nil)
				return mockSvc
			},

			expectedStatus:   http.StatusOK,
			expectedResponse: `{"password":"password"}`,
		},
		{
			name: "err case",
			setupMockSvc: func(ctx context.Context) *genpass_svc_mocks.GenPass {
				mockSvc := genpass_svc_mocks.NewGenPass(t)
				mockSvc.On("GeneratePassword", DefaultLength).Return("", errors.New("test err"))
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

			mockSvc := tc.setupMockSvc(ginCtx)
			genPassHandler := NewGenPass(mockSvc)
			genPassHandler.GeneratePassword(ginCtx)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.Equal(t, tc.expectedResponse, rec.Body.String())
		})
	}
}
