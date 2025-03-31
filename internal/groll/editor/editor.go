package editor

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

type Editor struct {
	cursorX       int
	cursorY       int
	contents      string
	fd            int
	originalState *term.State
	debugMode     bool
}

func Init(debug bool, fd int, contents string) *Editor {
	return &Editor{
		cursorX:   0,
		cursorY:   0,
		contents:  contents,
		fd:        fd,
		debugMode: debug,
	}
}

func (e *Editor) DisplayContents() {
	fmt.Print(e.contents)
}

func (e *Editor) ClearScreen() {
	fmt.Print("\x1b[2J\x1b[H")
}

func (e *Editor) DrawCursor() {
	fmt.Print(fmt.Sprintf("\x1b[%d;%dH", e.cursorY+1, e.cursorX+1))
	fmt.Print("\x1b[6 q")

}

func (e *Editor) MakeRaw() error {
	orgiState, err := term.MakeRaw(e.fd)
	if err != nil {
		return err
	}
	fmt.Print("\x1b[?1049h")
	e.originalState = orgiState
	return nil
}

func (e *Editor) Restore() {
	if err := term.Restore(e.fd, e.originalState); err != nil {
		panic(err)
	}
}

func (e *Editor) Loop() {
	// TODO: Need to read in characters from user
	for {
		if !e.debugMode {
			e.ClearScreen()
			e.DisplayContents()
			e.DrawCursor()
		}
		if should_break := e.HandleInput(); should_break == true {
			break
		}
	}
}

func (e *Editor) HandleInput() bool {

	byte_read := make([]byte, 1)
	n, err := os.Stdin.Read(byte_read)
	if err != nil {
		fmt.Println(err)
	}

	if n == 0 {
		fmt.Println("Nothing read")
	}

	switch byte_read[0] {
	case 3:
		return true
	case 27:
		escape_bytes := make([]byte, 2)
		os.Stdin.Read(escape_bytes)
		if escape_bytes[1] == 65 {
		} else if escape_bytes[1] == 66 {
		} else if escape_bytes[1] == 67 {
		} else if escape_bytes[1] == 68 {
		}
		break
	case 127:
		if len(e.contents) == 0 {
			break
		}
		e.cursorX--
		e.contents = e.contents[:len(e.contents)-1]
		break
	case 21:
		if len(e.contents) == 0 {
			break
		}
		temp_x := e.cursorX
		e.cursorX = 0
		e.contents = e.contents[:len(e.contents)-temp_x]
		break
	case 9:
		e.cursorX += 4
		e.contents += "    "
		break
	case 13:
		e.contents += "\r\n"
		e.cursorX = 0
		e.cursorY++
		break
	default:
		if e.debugMode {
			fmt.Printf("%d:%c\r\n", byte_read[0], byte_read[0])
		}
		e.cursorX++
		e.contents += string(byte_read[0])
		break
	}

	return false

}
