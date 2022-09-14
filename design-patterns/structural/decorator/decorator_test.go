package decorator

import "testing"

func TestDecorator(t *testing.T) {
	var c Component = &ConcreteComponent{}
	c = WrapAddDecorator(c, 2)
	c = WrapMulDecorator(c, 5)
	want := 10
	if got := c.Calc(); got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
