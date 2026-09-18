package repository

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"

	redisPkg "github.com/EBVN/bookmark-management/pkg/redis"
)

func TestUrlStorage_StoreURL(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupMock func(ctx context.Context) *redis.Client

		expectedErr error
		verifyFunc  func(ctx context.Context, r *redis.Client)
	}{
		{
			name: "normal case",
			setupMock: func(ctx context.Context) *redis.Client {
				return redisPkg.InitMockRedis(t)
			},

			expectedErr: nil,
			verifyFunc: func(ctx context.Context, r *redis.Client) {
				res, err := r.Get(ctx, "abcdef").Result()
				assert.NoError(t, err)
				assert.Equal(t, "https://google.com", res)
			},
		},
		{
			name: "connection error case",
			setupMock: func(ctx context.Context) *redis.Client {
				mockR := redisPkg.InitMockRedis(t)
				_ = mockR.Close()
				return mockR
			},

			expectedErr: redis.ErrClosed,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			mock := tc.setupMock(ctx)

			testRepo := NewURLStorage(mock)
			err := testRepo.StoreURL(ctx, "abcdef", "https://google.com")
			assert.ErrorIs(t, tc.expectedErr, err)
			if tc.verifyFunc != nil {
				tc.verifyFunc(ctx, mock)
			}
		})
	}
}
