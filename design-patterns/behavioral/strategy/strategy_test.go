package strategy

import "testing"

func TestStrategy(t *testing.T) {
	add := Operation{Addition{}}
	want := 5
	if got := add.Operate(2, 3); got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	multi := Operation{Multiplication{}}
	want = 6
	if got := multi.Operate(2, 3); got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
