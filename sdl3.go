/*
This Source Code Form is subject to the terms of the Mozilla
Public License, v. 2.0. If a copy of the MPL was not distributed
with this file, You can obtain one at http://mozilla.org/MPL/2.0/.

===-===-===-===-===-===-===-===-===-===
Donations during this file development:
-===-===-===-===-===-===-===-===-===-===

none :)

Thank you for your support!
*/

package gomp

import (
	"github.com/jupiterrider/purego-sdl3/sdl"
	"github.com/jupiterrider/purego-sdl3/ttf"
	"time"
)

type SDLRender struct {
	Renderer    *sdl.Renderer
	Window      *sdl.Window
	windowTitle string
	width       int32
	height      int32
}

func NewSDLRenderer(windowTitle string, width, height int32) *SDLRender {
	return &SDLRender{
		windowTitle: windowTitle,
		width:       width,
		height:      height,
	}
}

func (s *SDLRender) Init() {
	if !sdl.Init(sdl.InitVideo | sdl.InitEvents) {
		panic("SDL_Init failed: " + sdl.GetError())
	}
	if !ttf.Init() {
		panic("TTF_Init failed")
	}
	//TODO: ability to resize from config menu on the fly
	if !sdl.CreateWindowAndRenderer(s.windowTitle, s.width, s.height, sdl.WindowResizable, &s.Window, &s.Renderer) {
		panic(sdl.GetError())
	}
	//sdl.SetRenderDrawColor(s.Renderer, 0, 0, 0, 255)
	//sdl.SetRenderDrawBlendMode(s.Renderer, sdl.BlendModeBlend)
	//sdl.RenderClear(s.Renderer)
	//sdl.RenderPresent(s.Renderer)
}

func (s *SDLRender) Prerender() bool {
	var event sdl.Event
	for sdl.PollEvent(&event) {
		switch event.Type() {
		case sdl.EventQuit:
			return false
		case sdl.EventKeyDown:
			switch event.Key().Scancode {
			case sdl.ScancodeEscape:
				return false
			}
		}
	}
	return true
}

func (s *SDLRender) Render(dt time.Duration) {
	sdl.RenderPresent(s.Renderer)
}

func (s *SDLRender) Destroy() {
	ttf.Quit()
	sdl.DestroyRenderer(s.Renderer)
	sdl.DestroyWindow(s.Window)
	sdl.Quit()
}
