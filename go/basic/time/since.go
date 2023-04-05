package main

import (
	"log"
	"time"
)

func Duration(invocation time.Time, name string) {
	elapsed := time.Since(invocation)

	log.Printf("%s called lasted %s", name, elapsed)
}

func main() {
	defer Duration(time.Now(), "something name")

	// TODO Something
	time.Sleep(1 * time.Second)
}
