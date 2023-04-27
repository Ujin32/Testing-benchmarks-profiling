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

func TestContainsDuplicate(t *testing.T) {
	cases := map[string]struct {
		slice []int
		want  bool
	}{
		"Is duplicate ":           {[]int{1, 2, 1, 3}, true},
		"Is not duplicate ":       {[]int{112222, 2, 44, 55}, false},
		"Empty slice":             {[]int{}, false},
		"One value in slice":      {[]int{2}, false},
		"Negative value in slice": {[]int{0, 2, -1, 1}, false},
	}

	for name, testCase := range cases {
		t.Run(name, func(t *testing.T) {
			res := ContainsDuplicate(testCase.slice)
			require.Equal(t, testCase.want, res)
		})
	}

}

func TestIsPalindrom(t *testing.T) {
	palidromCase := func(i int, want bool) func(t *testing.T) {
		return func(t *testing.T) {
			res := IsPalindrome(i)
			require.Equal(t, want, res)
		}

	}

	t.Run("zero", palidromCase(0, true))
	t.Run("palindrom", palidromCase(131, true))
	t.Run("is not palindrom", palidromCase(134, false))
	t.Run("negative", palidromCase(-423, false))
	t.Run("negative palidrom", palidromCase(-323, false))

}
