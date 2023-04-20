package util

import (
	"testing"
)

/*
func ReverseInt(x interface{}) (int, error) {
	y, ok := x.(int)
	if !ok {
		return 0, errors.New("not int")
	}

	var result int
	for y != 0 {
		result = result*10 + y%10
		if result > 2147483647 || result < -2147483648 {
			return 0, nil
		}

		y /= 10
	}

	return result, nil
}
*/

/*
Что будет, если передать туда обычное число? Например, 123 ?
Что будет, если передать туда отрицательное число? Например, -649 ?
Что будет, если передать туда 0 ?
*/

func TestReverseInt(t *testing.T) {
	testDigit := []int{123, -649, 0}

	for digit := range testDigit {
		_, err := ReverseInt(digit)
		if err != nil {
			t.Errorf("fail to refers - %s", err)
		}

	}

	_, err := ReverseInt("123")
	if err == nil {
		t.Errorf("error is not exist")
	}

	if err.Error() != "not int" {
		t.Errorf("Improper error handling response")
	}

}
