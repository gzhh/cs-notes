package simplefactory

import "testing"

func TestHelloAPI(t *testing.T) {
	// get instance with factory
	api := NewAPI(HelloAPI)
	want := "Hello, World"
	if got := api.Say("World"); got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestHiAPI(t *testing.T) {
	api := NewAPI(HiAPI)
	want := "Hi, World"
	if got := api.Say("World"); got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
