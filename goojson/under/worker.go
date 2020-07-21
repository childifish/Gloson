package under

import (
	"errors"
	"strings"
)

var tc TypeChanger

func CleanSpace(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "\n", ""), " ", "")
}

func CheckJson(rawBytes []byte) error {
	if len(rawBytes) != 0 {
		if rawBytes[0] != 123 || rawBytes[len(rawBytes)-1] != 125 {
			return errors.New("not correct json")
		}
		return nil
	} else {
		return errors.New("nil json")
	}
}

func ClearMark(s string) string {
	b := tc.Str2bytes(s)
	if b[0] == 34 {
		b = b[1 : len(b)-1]
		return tc.Bytes2str(b)
	}
	return tc.Bytes2str(b)
}
