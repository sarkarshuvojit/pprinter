package pprinter

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

func printWithColor(text string, color lipgloss.Color) {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(color)

	fmt.Println(style.Render(text))
}

func Info(message string, args ...string) {
	var fmtMsg string
	if len(args) == 0 {
		fmtMsg = message
	} else {
		fmtMsg = fmt.Sprintf(message, args)
	}
	printWithColor("i "+fmtMsg, "#3498db")
}

func Success(message string, args ...string) {
	var fmtMsg string
	if len(args) == 0 {
		fmtMsg = message
	} else {
		fmtMsg = fmt.Sprintf(message, args)
	}
	printWithColor("✓ "+fmtMsg, "#27ae60")
}

func Error(message string, args ...string) {
	var fmtMsg string
	if len(args) == 0 {
		fmtMsg = message
	} else {
		fmtMsg = fmt.Sprintf(message, args)
	}
	printWithColor("x "+fmtMsg, "#c0392b")
}

func Warning(message string, args ...string) {
	var fmtMsg string
	if len(args) == 0 {
		fmtMsg = message
	} else {
		fmtMsg = fmt.Sprintf(message, args)
	}
	printWithColor("! "+fmtMsg, "#f39c12")
}

func Terminate(message string, args ...string) {
	Error(message, args...)
	os.Exit(1)
}

func TerminateWithHelpText(message string, helpText string, args ...string) {
	Error(message, args...)
	Info(helpText)
	os.Exit(1)
}

func AddCursorCheckPoint() {
	// ref: https://en.wikipedia.org/wiki/ANSI_escape_code
	fmt.Print("\033[s")
}

func ResetToCursorCheckpoint() {
	// ref: https://en.wikipedia.org/wiki/ANSI_escape_code
	fmt.Print("\033[u\033[K")
}
