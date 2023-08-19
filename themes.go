package pprinter

type Theme struct {
	ColorSuccess string
	ColorInfo    string
	ColorWarn    string
	ColorError   string
}

var (
	PastelTheme = Theme{
		ColorSuccess: "#27ae60",
		ColorInfo:    "#3498db",
		ColorWarn:    "#f39c12",
		ColorError:   "#c0392b",
	}

	AyuTheme = Theme{
		ColorSuccess: "#b3b846",
		ColorInfo:    "#66aabb",
		ColorWarn:    "#d18e44",
		ColorError:   "#c33e40",
	}
	AyuLightTheme = Theme{
		ColorSuccess: "#a3c08b",
		ColorInfo:    "#9eb0b7",
		ColorWarn:    "#d0a176",
		ColorError:   "#c07e7e",
	}

	MonokaiTheme = Theme{
		ColorSuccess: "#a6e22e",
		ColorInfo:    "#66d9ef",
		ColorWarn:    "#fd971f",
		ColorError:   "#f92672",
	}
	MonokaiLightTheme = Theme{
		ColorSuccess: "#c3e88d",
		ColorInfo:    "#9cc2d3",
		ColorWarn:    "#ffb86c",
		ColorError:   "#ff7b8a",
	}

	DraculaTheme = Theme{
		ColorSuccess: "#50fa7b",
		ColorInfo:    "#8be9fd",
		ColorWarn:    "#f1fa8c",
		ColorError:   "#ff5555",
	}
	DraculaLightTheme = Theme{
		ColorSuccess: "#7aa56c",
		ColorInfo:    "#85a7db",
		ColorWarn:    "#d4c36a",
		ColorError:   "#db615d",
	}

	SolarizedTheme = Theme{
		ColorSuccess: "#859900",
		ColorInfo:    "#268bd2",
		ColorWarn:    "#b58900",
		ColorError:   "#dc322f",
	}
	SolarizedLightTheme = Theme{
		ColorSuccess: "#a6a74f",
		ColorInfo:    "#7791b3",
		ColorWarn:    "#c5a632",
		ColorError:   "#d94c4c",
	}

	OneDarkTheme = Theme{
		ColorSuccess: "#50e3c2",
		ColorInfo:    "#61afef",
		ColorWarn:    "#e5c07b",
		ColorError:   "#ff6c6b",
	}
	OneDarkLightTheme = Theme{
		ColorSuccess: "#6dd3a8",
		ColorInfo:    "#89a1e2",
		ColorWarn:    "#e8d39b",
		ColorError:   "#ff9f9e",
	}
)
