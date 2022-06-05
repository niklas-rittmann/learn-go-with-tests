package main

const englischHelloPrefix = "Hello, "
const germanHelloPrefix = "Hallo, "
const spanishHelloPrefix = "Hola, "
const german = "German"
const spanish = "Spanish"
const world = "World"

// Returns a greeting by proving the language and name
func Hello(name, language string) string {
	if name == "" {
		name = world
	}

	return greetingPrefix(language) + name
}

// Select the language prefix based on input
func greetingPrefix(language string) (languagePrefix string) {
	switch language {
	case german:
		languagePrefix = germanHelloPrefix
	case spanish:
		languagePrefix = spanishHelloPrefix
	default:
		languagePrefix = englischHelloPrefix
	}
	return
}
