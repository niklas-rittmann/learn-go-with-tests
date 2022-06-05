package main

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Niklas", "")
		want := "Hello, Niklas"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying hello to World", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)

	})
	t.Run("in German", func(t *testing.T) {
		got := Hello("Niklas", "German")
		want := "Hallo, Niklas"
		assertCorrectMessage(t, got, want)

	})
}

func TestLanguagePrefix(t *testing.T) {
	t.Run("in Spanish", func(t *testing.T) {
		got := greetingPrefix("Spanish")
		want := "Hola, "
		assertCorrectMessage(t, got, want)

	})
	t.Run("in German", func(t *testing.T) {
		got := greetingPrefix("German")
		want := "Hallo, "
		assertCorrectMessage(t, got, want)

	})
	t.Run("Default", func(t *testing.T) {
		got := greetingPrefix("")
		want := "Hello, "
		assertCorrectMessage(t, got, want)

	})

}
