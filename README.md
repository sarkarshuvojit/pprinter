# pprinter

Minimal utility package to print user friendly messages to the console.

Can be used as an alternative to `fmt.Println` while printing messages from your CLI tool.


## Features

- Minimal API
- Single dependency on [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss)
- Supports color schemes
- 10 inbuilt themes
- Support for custom themes
- Other utilities to print animated lines (like a progress bar)
## Installation

Install from [pkg.go.dev](https://pkg.go.dev/github.com/sarkarshuvojit/pprinter)

```bash
go get github.com/sarkarshuvojit/pprinter
```
## Usage/Examples

Simply import the package and intialise with a predefined theme & a custom theme of your own.

Then you can call the underlying methods.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/sarkarshuvojit/pprinter/pprinter"
)

func main() {
	pp := pprinter.WithTheme(pprinter.PastelTheme)
    p.Info("Test")
	p.Success("Test")
	p.Warning("Test")
	p.Error("Test")
}

```

Which would yeild the following output

[![asciicast](https://asciinema.org/a/603677.svg)](https://asciinema.org/a/603677)


### Other Examples 

[Themes Example](examples/themes/)

[![asciicast](https://asciinema.org/a/603681.svg)](https://asciinema.org/a/603681)

[Animated Text Example](examples/animated_text/)

[![asciicast](https://asciinema.org/a/603679.svg)](https://asciinema.org/a/603679)
