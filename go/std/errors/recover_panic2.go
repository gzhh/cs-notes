package main

import (
	"runtime"
	"time"

	zlog "github.com/rs/zerolog/log"
)

func CatchPanic(funcName string, fields ...map[string]any) {
	if err := recover(); err != nil {
		stack := make([]byte, 8096)
		stack = stack[:runtime.Stack(stack, false)]
		logger := zlog.Logger.Log().Bytes("stack", stack).Str("level", "fatal").Interface("error", err)
		if len(fields) > 0 {
			logger.Fields(fields[0])
		}
		logger.Msgf("recovered from panic -%s", funcName)
	}
}

func main() {
	go func() {
		defer CatchPanic("myBackgroundJob")
		// do some heavy lifting

		panic("Panic in goroutine")
	}()

	time.Sleep(time.Second)
}
