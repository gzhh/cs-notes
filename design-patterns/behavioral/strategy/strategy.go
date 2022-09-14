package strategy

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

type Addition struct{}

func (Addition) Apply(lv, rv int) int {
	return lv + rv
}

type Multiplication struct{}

func (Multiplication) Apply(lv, rv int) int {
	return lv * rv
}
