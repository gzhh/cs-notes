package adapter

import "testing"

func TestAdapter(t *testing.T) {
	want := "adaptee method"
	adaptee := NewAdaptee()
	adapter := NewAdapter(adaptee)
	if got := adapter.Request(); got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
