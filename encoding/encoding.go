package encoding

import (
	"errors"
	"strings"
)

const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
const length = uint64(len(alphabet))

func Encode(number uint64) string {
	var sb strings.Builder
	for number > 0 {
		i := number % length
		number = number / length
		char := alphabet[i]
		sb.WriteByte(char)
	}
	return sb.String()
}

func Decode(s string) (uint64, error) {
	var res, base uint64 = 0, 1
	for i := 0; i < len(s); i++ {
		digit := strings.IndexByte(alphabet, s[i])
		if digit == -1 {
			return 0, errors.New("Link contains character(s) not present in base58 encoding")
		}
		res += base * uint64(digit)
		base = base * length
	}
	return res, nil
}
