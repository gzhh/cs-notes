package observer

import (
	"testing"
	"time"
)

func TestObserver(t *testing.T) {
	// Initialize a new Notifier.
	n := eventNotifier{
		observers: map[Observer]struct{}{},
	}

	// Register a couple of observers.
	n.Register(&eventObserver{id: 1})
	n.Register(&eventObserver{id: 2})

	// A simple loop publishing the current Unix timestamp to observers.
	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C
	for {
		select {
		case <-stop:
			return
		case t := <-tick:
			n.Notify(Event{Data: t.UnixNano()})
		}
	}
}
