package my_util

import "strconv"

func StringToInt64Many(stringMany []string) (numbers []int64, err error) {
	var parseInt int64
	for _, s := range stringMany {
		parseInt, err = strconv.ParseInt(s, 10, 64)
		if err != nil {
			return
		} else {
			numbers = append(numbers, parseInt)
		}
	}
	return
}

func Int64ToStringMany(int64Many []int64) (stringMany []string) {
	var s string
	for _, number := range int64Many {
		s = strconv.FormatInt(number, 10)
		stringMany = append(stringMany, s)
	}
	return
}
