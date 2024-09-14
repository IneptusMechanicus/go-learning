package greetings

import (
	"testing"
	"regexp"
)

func TestMessageDefault() {
	var name = "Ligma"
	var want = regexp.MustCompile(`\b` + name + `\b`)
}


