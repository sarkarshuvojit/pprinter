package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sarkarshuvojit/pprinter/pprinter"
)

func main() {
	p := pprinter.WithTheme(&pprinter.MonokaiTheme)
	p.Info("This progres bar will go on for 50s")
	p.AddCursorCheckPoint()

	for i := 0; i < 100; i += rand.Intn(10) {
		p.ResetToCursorCheckpoint()
		fmt.Printf("Progress %d%%", i)
		time.Sleep(time.Millisecond * 500)
	}

	p.ResetToCursorCheckpoint()
	p.Success("Complete")
}
