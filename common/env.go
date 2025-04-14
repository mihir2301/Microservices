package common

import (
	"syscall"
)

func EnvString(key, fallback string) string {
	val, ok := syscall.Getenv(key)
	if ok {
		return val
	}
	return fallback
}
