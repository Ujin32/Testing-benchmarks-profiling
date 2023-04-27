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
		"1": {[]int{1, 2, 1, 3}, true},
		"2": {[]int{112222, 2, 44, 55}, false},
		"3": {[]int{}, false},
		"4": {[]int{2}, false},
		"5": {[]int{0, 2, -1, 1}, false},
	}

	for name, testCase := range cases {
		t.Run(name, func(t *testing.T) {
			res := ContainsDuplicate(testCase.slice)
			require.Equal(t, testCase.want, res)
		})
	}

}
