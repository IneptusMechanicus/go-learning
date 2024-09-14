package greetings

import (
	"fmt"
	"errors"
)

func Message(param string) (string, error) {
	if param == "" {
		return "", errors.New("Empty param!");
	}
	message := fmt.Sprintf("Testing using: %v" , param);
	return message, nil;
}

func MultiMessage(names []string) (map[string]string, error) {
	var messages = make(map[string]string);

	for _, name := range names {
		if messages[name] == "" {
			return nil, errors.New("Empty name in list!");
		}

		messages[name] = RandomFormat();
	}

	return messages, nil;
}
