package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePassword(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		expectedLen int
		expectedErr error
	}{
		{
			name: "normal case - valid length",

			expectedLen: 10,
		},
		{
			name: "normal case 2 - valid length",

			expectedLen: 15,
		},
		{
			name: "edge case - valid length",

			expectedLen: 1000000,
		},
		{
			name: "edge case 2 - valid length",

			expectedLen: 2000000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			svc := NewGenPass()

			// run test
			password, err := svc.GeneratePassword(tc.expectedLen)
			assert.ErrorIs(t, err, tc.expectedErr)
			assert.Equal(t, tc.expectedLen, len(password))
		})
	}
}
