package greet

import "strings"

func Hello(name string) string {

	return "Hola, " + normalizeName(name) + " :-)"
}

func normalizeName(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return "mundo"
	}

	return name
}
