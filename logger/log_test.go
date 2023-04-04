package logger

import (
	"fmt"
	"testing"
)

func TestPrintColor(t *testing.T) {
	printColor := "\033[38;5;%dm%s\033[39;49m - number : %v\n"
	for j := 0; j < 256; j++ {
		fmt.Printf(printColor, j, "Hello! ", j)
	}
}

func TestDebug(t *testing.T) {
	l := NewLogger()
	l.SetLevel(Debug)
	l.Debug("test debug")
	l.Info("test info")
	l.Error("test error")
	// l.Panic("test panic")
	l.Fatal("test Fatal")
}
