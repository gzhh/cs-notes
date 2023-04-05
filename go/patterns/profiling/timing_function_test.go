package profiling

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	defer Duration(time.Now(), "something name")

	// TODO Something
	time.Sleep(1 * time.Second)
}
