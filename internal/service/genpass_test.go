package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenPass_GeneratePassword(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		expectedLen   int
		expectedError error
	}{
		{
			name: "normal case",

			expectedLen:   10,
			expectedError: nil,
		},
		{
			name: "normal case 2",

			expectedLen:   15,
			expectedError: nil,
		},
		{
			name: "edge case",

			expectedLen:   1000000,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// khoi tao service
			svc := NewGenPass()

			// run test
			password, err := svc.GeneratePassword(tc.expectedLen)

			// check result
			assert.ErrorIs(t, err, tc.expectedError)
			assert.Equal(t, tc.expectedLen, len(password))
		})
	}
}
