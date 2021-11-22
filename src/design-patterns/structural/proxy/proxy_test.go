package proxy

import "testing"

func TestProxy(t *testing.T) {
	var proxy Subject
	proxy = &Proxy{}

	want := "pre:real:after"
	if got := proxy.Do(); got != want {
		t.Fatalf("want %s, got %s", want, got)
	}
}
