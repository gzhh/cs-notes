package synchronization

import (
	"os"
	"testing"
	"time"
)

func TestSemaphoreWithTimeouts(t *testing.T) {
	tickets, timeout := 1, 3*time.Second
	s := NewSemaphore(tickets, timeout)

	if err := s.Acquire(); err != nil {
		panic(err)
	}

	// Do important work

	if err := s.Release(); err != nil {
		panic(err)
	}
}

// Non-Blocking
func TestSemaphoreWithoutTimeouts(t *testing.T) {
	tickets, timeout := 0, 0
	s := NewSemaphore(tickets, time.Duration(timeout))

	if err := s.Acquire(); err != nil {
		if err != ErrNoTickets {
			panic(err)
		}

		// No tickets left, can't work :(
		os.Exit(1)
	}
}
