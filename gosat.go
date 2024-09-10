package main

import (
	"local/gosat/insignals"
	"local/gosat/vector"
	"local/gosat/visual"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	ins := insignals.New()

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("the best window", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// implement of rederer flags like render to texture
	// flags:
	// software rendere = 0x1
	// acceleratet renderer = 0x2
	// Present() with VSync = 0x4
	// Enable Render to Texture = 0x8
	rend, err := sdl.CreateRenderer(window, -1, 0xA) // acceleratet and texture
	if err != nil {
		panic(err)
	}
	defer rend.Destroy()

	texter := visual.NewTextDrawer(rend)

	label := visual.TextLabel{
		Text:     "Test text",
		Position: vector.Vector{X: 200, Y: 200},
		Scale:    vector.Vector{X: 2, Y: 2},
		Spacing:  vector.Vector{X: 0, Y: 0},
	}

	rect := sdl.Rect{X: 0, Y: 0, W: 200, H: 200}

	for ins.Quit == false && ins.Keys[sdl.SCANCODE_ESCAPE] == false {
		ins.Update()
		rend.SetDrawColor(128, 128, 128, 255)
		rend.Clear()
		rend.SetDrawColor(255, 0, 255, 255)
		rend.FillRect(&rect)
		texter.Draw(&label)

		rend.Present()
		sdl.Delay(20)
	}
}
