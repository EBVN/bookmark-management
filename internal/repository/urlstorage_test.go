package repository

import (
	"context"
	"testing"

	redisPkg "github.com/EBVN/bookmark-management/pkg/redis"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
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
				res, err := r.Get(ctx, "abcxyz").Result()
				assert.NoError(t, err)
				assert.Equal(t, "https://google.com", res)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()

			mock := tc.setupMock(ctx)
			testRepo := NewURLStorage(mock)
			err := testRepo.StoreURL(ctx, "abcxyz", "https://google.com")
			assert.ErrorIs(t, err, tc.expectedErr)
			if tc.verifyFunc != nil {
				tc.verifyFunc(ctx, mock)
			}
		})
	}
}
