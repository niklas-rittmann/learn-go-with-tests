package main

const englischHelloPrefix = "Hello, "
const germanHelloPrefix = "Hallo, "
const spanishHelloPrefix = "Hola, "
const german = "German"
const spanish = "Spanish"
const world = "World"

func Hello(name, language string) string {
	if name == "" {
		name = world
	}

	return greetingPrefix(language) + name
}

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
