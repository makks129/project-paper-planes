package utils

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Log(messages ...any) {
	if os.Getenv("GO_ENV") == "test" {
		return
	}
	ss := []string{}
	for _, m := range messages {
		ss = append(ss, fmt.Sprint(m))
	}
	fmt.Print(logTime(), strings.Join(ss, " "), "\n")
}

func Logt(message ...string) {
	if os.Getenv("GO_ENV") == "test" {
		return
	}
	fmt.Print(logTime(), strings.Join(message, " "))
}

func Logm(message ...string) {
	if os.Getenv("GO_ENV") == "test" {
		return
	}
	fmt.Print(strings.Join(message, " "))
}

func Error(error error) {
	if os.Getenv("GO_ENV") == "test" {
		return
	}
	fmt.Print(logTime(), error.Error(), "\n")
}

func logTime() string {
	return "" + time.Now().UTC().Format("15:04:05.000") + " | "
}
