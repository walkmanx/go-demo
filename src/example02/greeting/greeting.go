package greeting

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name")
	}
	message := fmt.Sprintf("hi，%v. Welcome!", name)
	return message, nil
}
