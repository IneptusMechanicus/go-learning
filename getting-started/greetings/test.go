package submod

import (
	"fmt"
	"errors"
)

func Message(param string) (string, error){
	if param == "" {
		return "", errors.New("Empty param!")
	}

	message := fmt.Sprintf("Testing using: %v" , param)

	return message, nil
}
