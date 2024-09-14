package greetings

import (
	"math/rand"
)

func RandomFormat() string {
	messageOptions := []string {
		"%v, what a surprise!",
		"Oh hi, %v",
		"Mister %v, welcome back!",
	};

	return messageOptions[rand.Intn(len(messageOptions))];
}
