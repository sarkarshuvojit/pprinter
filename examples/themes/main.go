package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/sarkarshuvojit/pprinter"
)

func printAll(p *pprinter.Pprinter) {
	p.Info("Test")
	p.Success("Test")
	p.Warning("Test")
	p.Error("Test")
}

func main() {
	themes := map[string]*pprinter.Theme{
		"PastelTheme":         &pprinter.PastelTheme,
		"AyuTheme":            &pprinter.AyuTheme,
		"AyuLightTheme":       &pprinter.AyuLightTheme,
		"MonokaiTheme":        &pprinter.MonokaiTheme,
		"MonokaiLightTheme":   &pprinter.MonokaiLightTheme,
		"DraculaTheme":        &pprinter.DraculaTheme,
		"DraculaLightTheme":   &pprinter.DraculaLightTheme,
		"SolarizedTheme":      &pprinter.SolarizedTheme,
		"SolarizedLightTheme": &pprinter.SolarizedLightTheme,
		"OneDarkTheme":        &pprinter.OneDarkTheme,
		"OneDarkLightTheme":   &pprinter.OneDarkLightTheme,
	}
	for themeName := range themes {
		fmt.Println(themeName)
		fmt.Println(strings.Repeat("=", len(themeName)))
		printAll(pprinter.WithTheme(themes[themeName]))
		fmt.Println()
		time.Sleep(time.Second)
	}
}
