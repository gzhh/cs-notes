package messaging

import (
	"fmt"
	"testing"
	"time"
)

func TestMq(t *testing.T) {
	// New
	mque := NewMq()

	// Subscribe topic
	mque.Subscribe("Notify", func(m *Message) {
		fmt.Println("Received, Topic: Notify")
	})

	// Subscribe topic
	mque.Subscribe("Test", func(m *Message) {
		fmt.Println("Received, Topic: Test")
	})

	// Publish value into topic
	mque.Publish(
		Message{
			Topic: "Notify",
			Value: "blablabla",
		},
	)
	time.Sleep(1 * time.Second)

	// Publish value into topic
	mque.Publish(
		Message{
			Topic: "NotFound",
			Value: "blablabla",
		},
	)
	time.Sleep(1 * time.Second)

	// Publish value into topic
	mque.Publish(
		Message{
			Topic: "Test",
			Value: "blablabla",
		},
	)
	time.Sleep(1 * time.Second)
}
