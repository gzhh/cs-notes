package messaging

import "errors"

//Message struct
type Message struct {
	Topic string
	Value interface{}
}

//Channel struct
type Channel struct {
	ch chan Message
}

//Mq struct
type Mq struct {
	topics map[string]*Channel
}

// var sessions map[string][]Session

//New func
func NewMq() *Mq {
	return &Mq{
		topics: map[string]*Channel{},
	}
}

//Subscribe method
func (s *Mq) Subscribe(topic string, handler func(m *Message)) error {
	// generate topic if not exist
	if _, exist := s.topics[topic]; exist {
		//TODO: make Session List for handle multiple Subscribe on Single Topic
		return errors.New("subscribe exist, topic:" + topic)
	}

	s.topics[topic] = &Channel{
		ch: make(chan Message),
	}

	go func() {
		for {
			c := <-s.topics[topic].ch
			handler(&c)
		}
	}()
	return nil
}

//Publish method
func (s *Mq) Publish(msg Message) error {
	if _, ok := s.topics[msg.Topic]; !ok {
		return errors.New("topic has been closed")
	}

	s.topics[msg.Topic].ch <- msg

	return nil
}
