package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/BryanMcGahan/groll/internal/groll/editor"
)

func main() {

	debugFlag := flag.Bool("debug", false, "Start in Debug Mode")
	flag.Parse()

	currentEditor := editor.Init(*debugFlag, int(os.Stdin.Fd()), "")
	if err := currentEditor.MakeRaw(); err != nil {
		panic(err)
	}
	defer func() {
		fmt.Print("\x1b[?1049l")
		currentEditor.Restore()
	}()

	currentEditor.ClearScreen()
	currentEditor.ClearScreen()
	currentEditor.Loop()

}
