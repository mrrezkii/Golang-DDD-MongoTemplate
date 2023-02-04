package helper

import (
	"errors"
	"strings"
)

func ConvertIntToBool(number int) bool {
	return number == 1
}

func StringToBool(s string) (bool, error) {
	switch strings.ToLower(s) {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, errors.New("invalid bool")
	}

}
