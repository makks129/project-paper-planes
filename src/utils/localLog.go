package utils

import (
	"fmt"
	"strings"
	"time"
)

func Log(messages ...any) {
	ss := []string{}
	for _, m := range messages {
		ss = append(ss, fmt.Sprint(m))
	}
	fmt.Print(logTime(), strings.Join(ss, " "), "\n")
}

func Logt(message ...string) {
	fmt.Print(logTime(), strings.Join(message, " "))
}

func Logm(message ...string) {
	fmt.Print(strings.Join(message, " "))
}

func Error(error error) {
	fmt.Print(logTime(), error.Error(), "\n")
}

func logTime() string {
	return "" + time.Now().Format("15:04:05.000") + " | "
}
