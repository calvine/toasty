package util

import (
	"runtime"
)

func GetLineEnding() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}
