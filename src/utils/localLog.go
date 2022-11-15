package utils

import (
	"fmt"
	"strings"
	"time"
)

func Log(message ...string) {
	fmt.Print(logTime(), strings.Join(message, " "), "\n")
}

func Logt(message ...string) {
	fmt.Print(logTime(), strings.Join(message, " "))
}

func Logm(message ...string) {
	fmt.Print(strings.Join(message, " "))
}

func logTime() string {
	return "" + time.Now().Format("15:04:05.000") + " | "
}