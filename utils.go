package main

import "errors"

func reverseInt(x interface{}) (int, error) {
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

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := set[v]; ok {
			return true
		} else {
			set[v] = struct{}{}
		}
	}
	return false
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var revertedNum int
	for x > revertedNum {
		last := x % 10
		x /= 10
		revertedNum = revertedNum*10 + last
	}

	return x == revertedNum || x == revertedNum/10
}
