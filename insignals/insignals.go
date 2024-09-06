package insignals

import (
	// "fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Insignals struct {
	Quit bool
	Keys map[byte]bool
}

func New() Insignals {
	return Insignals{
		Quit: false,
		Keys: make(map[byte]bool),
	}

}

func (i *Insignals) Update() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			i.Quit = true
		case *sdl.KeyboardEvent:
			k := event.(*sdl.KeyboardEvent)
			if k.State == sdl.PRESSED && k.Repeat == 0 {
				i.Keys[byte(k.Keysym.Scancode)] = true
			}
			if k.State == sdl.RELEASED && k.Repeat == 0 {
				i.Keys[byte(k.Keysym.Scancode)] = false
			}
		}

	}

}
