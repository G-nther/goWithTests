package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"

		assertCorrectMessate(t, got, want)
	})
	t.Run("saying 'Hello, World' with empty string", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessate(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessate(t, got, want)
	})
	t.Run("in french", func(t *testing.T) {
		got := Hello("Pierre", "French")
		want := "Bonjour, Pierre"

		assertCorrectMessate(t, got, want)
	})
}

func assertCorrectMessate(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
