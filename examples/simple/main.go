package main

import (
	"github.com/sarkarshuvojit/pprinter/pprinter"
)

func main() {
	p := pprinter.WithTheme(&pprinter.PastelTheme)
	p.Info("Test")
	p.Success("Test")
	p.Warning("Test")
	p.Error("Test")
}
