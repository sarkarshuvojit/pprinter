package pprinter_test

import (
	"testing"

	"github.com/sarkarshuvojit/pprinter"
)

func printAll(p *pprinter.Pprinter) {
	p.Info("Test")
	p.Success("Test")
	p.Warning("Test")
	p.Error("Test")
}

func Test_Exists(t *testing.T) {
	ayuLightPrinter := pprinter.WithTheme(&pprinter.AyuLightTheme)

	printAll(ayuLightPrinter)
}
