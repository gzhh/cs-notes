package simplefactory

import "fmt"

type APIType int

const (
	HelloAPI APIType = 1 << iota
	HiAPI
)

type API interface {
	Say(name string) string
}

func NewAPI(t APIType) API {
	switch t {
	case HelloAPI:
		return &helloAPI{}
	case HiAPI:
		return &hiAPI{}
	default:
		return nil
	}
}

type helloAPI struct{}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

type hiAPI struct{}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}
