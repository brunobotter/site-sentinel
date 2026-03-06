package validator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinNumberValidate(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		require.NoError(t, MinNumber(1).Validate(1))
	})

	t.Run("invalid value", func(t *testing.T) {
		err := MinNumber(1).Validate(0)
		require.Error(t, err)
		require.EqualError(t, err, "must be greater than or equal to 1")
	})
}

func TestMaxNumberValidate(t *testing.T) {
	t.Run("valid value", func(t *testing.T) {
		require.NoError(t, MaxNumber(100).Validate(100))
	})

	t.Run("invalid value", func(t *testing.T) {
		err := MaxNumber(100).Validate(101)
		require.Error(t, err)
		require.EqualError(t, err, "must be less than or equal to 100")
	})
}
