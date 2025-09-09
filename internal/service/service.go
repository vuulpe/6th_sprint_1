package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", fmt.Errorf("пустая строка")
	}

	if isMorseCode(trimmed) {
		return morse.ToText(trimmed), nil
	} else {
		return morse.ToMorse(trimmed), nil
	}
}

func isMorseCode(input string) bool {
	for _, char := range input {
		if char != '.' && char != '-' && char != ' ' && char != '/' {
			return false
		}
	}
	return true
}
