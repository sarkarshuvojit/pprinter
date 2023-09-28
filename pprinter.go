package pprinter

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

type pprinterService interface {
	Info(message string, args ...string)
	Success(message string, args ...string)
	Error(message string, args ...string)
	Warning(message string, args ...string)
	Terminate(message string, args ...string)
	TerminateWithHelpText(message string, helpText string, args ...string)
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

// Prints message with an leading 'i'
// Uses Theme.ColorInfo to color the text
func (p Pprinter) Info(message string, args ...string) {
	var fmtMsg string
	if len(args) == 0 {
		fmtMsg = message
	} else {
		fmtMsg = fmt.Sprintf(message, args)
	}
	printWithColor("i "+fmtMsg, lipgloss.Color(p.selectedTheme.ColorInfo))
}

// Prints message with an leading '✓'
// Uses Theme.ColorSuccess to color the text
func (p Pprinter) Success(message string, args ...string) {
	var fmtMsg string
	if len(args) == 0 {
		fmtMsg = message
	} else {
		fmtMsg = fmt.Sprintf(message, args)
	}
	printWithColor("✓ "+fmtMsg, lipgloss.Color(p.selectedTheme.ColorSuccess))
}

// Prints message with an leading 'x'
// Uses Theme.ColorError to color the text
func (p Pprinter) Error(message string, args ...string) {
	var fmtMsg string
	if len(args) == 0 {
		fmtMsg = message
	} else {
		fmtMsg = fmt.Sprintf(message, args)
	}
	printWithColor("x "+fmtMsg, lipgloss.Color(p.selectedTheme.ColorError))
}

// Prints message with an leading '!'
// Uses Theme.ColorSuccess to color the text
func (p Pprinter) Warning(message string, args ...string) {
	var fmtMsg string
	if len(args) == 0 {
		fmtMsg = message
	} else {
		fmtMsg = fmt.Sprintf(message, args)
	}
	printWithColor("! "+fmtMsg, lipgloss.Color(p.selectedTheme.ColorWarn))
}

// Calls Error(string, ...string)
// Uses Theme.ColorSuccess to color the text
// Then calls os.Exit(1)
func (p Pprinter) Terminate(message string, args ...string) {
	p.Error(message, args...)
	os.Exit(1)
}

// Prints `message` as Error
// Prints `helpText` as Info
// Then calls os.Exit(1)
func (p Pprinter) TerminateWithHelpText(message string, helpText string, args ...string) {
	p.Error(message, args...)
	p.Info(helpText)
	os.Exit(1)
}

// Sets a checkpoint for cursor to return back to
// Can be used to rewrite over a single line again and again in combination with ResetToCursorCheckpoint
// ref: https://en.wikipedia.org/wiki/ANSI_escape_code
func (p Pprinter) AddCursorCheckPoint() {
	fmt.Print("\033[s")
}

// Moves cursor to last checkpoint which was added calling AddCursorCheckPoint
// Any further prints will happen after last CheckPoint
// ref: https://en.wikipedia.org/wiki/ANSI_escape_code
func (p Pprinter) ResetToCursorCheckpoint() {
	fmt.Print("\033[u\033[K")
}
