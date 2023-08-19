package pprinter

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

type pprinterService interface {
	// Standard ones
	Info(message string, args ...string)
	Success(message string, args ...string)
	Error(message string, args ...string)
	Warning(message string, args ...string)

	// Show Error & Exit
	Terminate(message string, args ...string)
	TerminateWithHelpText(message string, helpText string, args ...string)

	// Cursor Movement
	AddCursorCheckPoint()
	ResetToCursorCheckpoint()
}

type Pprinter struct {
	selectedTheme *Theme
}

func WithTheme(t *Theme) *Pprinter {
	return &Pprinter{
		selectedTheme: t,
	}
}

func printWithColor(text string, color lipgloss.Color) {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(color)

	fmt.Println(style.Render(text))
}

func (p Pprinter) Info(message string, args ...string) {
	var fmtMsg string
	if len(args) == 0 {
		fmtMsg = message
	} else {
		fmtMsg = fmt.Sprintf(message, args)
	}
	printWithColor("i "+fmtMsg, "#3498db")
}

func (p Pprinter) Success(message string, args ...string) {
	var fmtMsg string
	if len(args) == 0 {
		fmtMsg = message
	} else {
		fmtMsg = fmt.Sprintf(message, args)
	}
	printWithColor("✓ "+fmtMsg, "#27ae60")
}

func (p Pprinter) Error(message string, args ...string) {
	var fmtMsg string
	if len(args) == 0 {
		fmtMsg = message
	} else {
		fmtMsg = fmt.Sprintf(message, args)
	}
	printWithColor("x "+fmtMsg, "#c0392b")
}

func (p Pprinter) Warning(message string, args ...string) {
	var fmtMsg string
	if len(args) == 0 {
		fmtMsg = message
	} else {
		fmtMsg = fmt.Sprintf(message, args)
	}
	printWithColor("! "+fmtMsg, "#f39c12")
}

func (p Pprinter) Terminate(message string, args ...string) {
	p.Error(message, args...)
	os.Exit(1)
}

func (p Pprinter) TerminateWithHelpText(message string, helpText string, args ...string) {
	p.Error(message, args...)
	p.Info(helpText)
	os.Exit(1)
}

func (p Pprinter) AddCursorCheckPoint() {
	// ref: https://en.wikipedia.org/wiki/ANSI_escape_code
	fmt.Print("\033[s")
}

func (p Pprinter) ResetToCursorCheckpoint() {
	// ref: https://en.wikipedia.org/wiki/ANSI_escape_code
	fmt.Print("\033[u\033[K")
}
