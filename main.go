package main

import (
	"runtime"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/laranc/monorepo/engine/graphics3d"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	title  = "Template"
	width  = 680
	height = 460
)

var (
	renderer *graphics3d.Renderer3D
)

func init() {
	runtime.LockOSThread()
}

func main() {
	var err error

	renderer, err = graphics3d.NewRenderer3D(title, width, height, 4, 6)
	if err != nil {
		panic(err)
	}
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	run()
}

func run() {
	running := true
	for running {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e.(type) {
			case *sdl.QuitEvent:
				running = false
			default:
				break
			}
		}
		renderer.RenderBegin()
		renderer.RenderEnd()
	}
}
