package decorator

type Component interface {
	Calc() int
}

type ConcreteComponent struct{}

func (*ConcreteComponent) Calc() int {
	return 0
}

type AddDecorator struct {
	Component
	num int
}

func (d *AddDecorator) Calc() int {
	return d.Component.Calc() + d.num
}

func WrapAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

type MulDecorator struct {
	Component
	num int
}

func (d *MulDecorator) Calc() int {
	return d.Component.Calc() * d.num
}

func WrapMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}
