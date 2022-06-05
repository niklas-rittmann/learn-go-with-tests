package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	}
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Niklas")
		want := "Hello, Niklas"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Saying hello to World", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)

	})
}
