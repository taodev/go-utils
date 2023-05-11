package random

import (
	"math/rand"
)

const (
	// LOWER 小写字母
	LOWER = 0x1
	// UPPER 大写字母
	UPPER = 0x2
	// NUMBER 数字
	NUMBER = 0x4
	// SYMBOL 符号
	SYMBOL = 0x8
)

const (
	lowerString  = "abcdefghijklmnopqrstuvwxyz"
	upperString  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberString = "1234567890"
	symbolString = "`~!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"
	randString   = lowerString + upperString + numberString + symbolString
)

func String(l int, flag int) string {
	str := ""

	if flag == 0 {
		str = randString
	} else {
		if flag&LOWER != 0 {
			str += lowerString
		}

		if flag&UPPER != 0 {
			str += upperString
		}

		if flag&NUMBER != 0 {
			str += numberString
		}

		if flag&SYMBOL != 0 {
			str += symbolString
		}
	}

	n := len(str)
	index := 0

	v := make([]byte, l)
	for i := 0; i < l; i++ {
		index = rand.Intn(n)
		v[i] = str[index]
	}

	return string(v)
}
