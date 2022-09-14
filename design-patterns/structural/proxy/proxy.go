package proxy

type Subject interface {
	Do() string
}

type RealSubject struct{}

func (RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real RealSubject
}

func (p Proxy) Do() string {
	var res string

	// TODO 在调用真实对象之前需要操作的
	res += "pre:"

	// 调用真实对象
	res += p.real.Do()

	// TODO 调用真实对象之后操作的
	res += ":after"

	return res
}
