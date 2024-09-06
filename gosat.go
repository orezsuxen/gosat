package main

import (
	"local/gosat/insignals"

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
	rend, err := sdl.CreateRenderer(window, -1, 2)
	if err != nil {
		panic(err)
	}
	defer rend.Destroy()

	rect := sdl.Rect{X: 0, Y: 0, W: 200, H: 200}

	for ins.Quit == false && ins.Keys[sdl.SCANCODE_ESCAPE] == false {
		ins.Update()
		rend.SetDrawColor(128, 128, 128, 255)
		rend.Clear()
		rend.SetDrawColor(255, 0, 255, 255)
		rend.FillRect(&rect)

		rend.Present()
	}
}
