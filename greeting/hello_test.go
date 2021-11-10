package greeting_test

import (
	"example/greeting"
	"testing"
)

func TestHello(t *testing.T) {
	got := greeting.Hello()
	want := "Hello, World!"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
