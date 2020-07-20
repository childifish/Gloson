package under

import (
	"errors"
	"strings"
)

var tc TypeChanger

func CleanSpace(s string)string  {
	return strings.ReplaceAll(strings.ReplaceAll(s,"\n","")," ","")
}

func CheckJson(rawBytes []byte)error {
	if rawBytes[0]!=123||rawBytes[len(rawBytes)-1]!=125{
		return errors.New("not correct json")
	}
	return nil
}




