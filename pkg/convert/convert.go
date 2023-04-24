package convert

import (
	"fmt"
	"strconv"
	"strings"
)

func ByteToInt(in []byte) (out []int) {
	for _, n := range in {
		out = append(out, int(n))
	}
	return out
}

func IntToStr(in []int) (out string) {
	for i, n := range in {
		out += fmt.Sprintf("%d", n)
		if i != len(in)-1 {
			out += ","
		}
	}
	return out
}

func StrToByte(in string) (out []byte) {
	var intermediateint []int
	value := strings.Split(in, ",")
	for _, nStr := range value {
		nInt, _ := strconv.Atoi(nStr)
		intermediateint = append(intermediateint, nInt)
	}
	for _, nInt := range intermediateint {
		out = append(out, byte(nInt))
	}
	return out
}
