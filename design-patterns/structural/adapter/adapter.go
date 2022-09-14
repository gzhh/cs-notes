package adapter

// Target 是适配器的目标接口
type Target interface {
	Request() string
}

// Adaptee 是适配者的目标接口
type Adaptee interface {
	SpecificRequest() string
}

func NewAdaptee() Adaptee {
	return &adaptee{}
}

type adaptee struct{}

func (*adaptee) SpecificRequest() string {
	return "adaptee method"
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

// adapter 是转换 adapter 为 Target 接口的适配器
type adapter struct {
	Adaptee
}

func (a *adapter) Request() string {
	return a.SpecificRequest()
}
