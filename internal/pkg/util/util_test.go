package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseInt(t *testing.T) {
	testDigit := []int{123, -649, 0}

	for digit := range testDigit {
		_, err := ReverseInt(digit)
		require.NoError(t, err)

	}
}

func TestReverseString(t *testing.T) {
	_, err := ReverseInt("123")
	require.NotEqual(t, "", err)
	require.Error(t, err)
	require.Contains(t, err.Error(), "not int")
}
