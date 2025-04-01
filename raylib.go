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
	rl "github.com/gen2brain/raylib-go/raylib"
	"time"
)

func NewRlRenderSystem(windowTitle string, width, height int32) *SDLRender {
	return &SDLRender{
		windowTitle: windowTitle,
		width:       width,
		height:      height,
	}
}

type RlRenderSystem struct {
	windowTitle string
	width       int32
	height      int32
}

func (s *RlRenderSystem) Init() {
	//monitor := rl.GetCurrentMonitor()
	//width, height := rl.GetMonitorWidth(monitor), rl.GetMonitorHeight(monitor)
	rl.InitWindow(s.width, s.height, s.windowTitle)
	//rl.SetWindowState(rl.FlagFullscreenMode)
}

func (s *RlRenderSystem) Prerender() bool {
	if rl.WindowShouldClose() {
		return false
	}
	return true
}

func (s *RlRenderSystem) Render(dt time.Duration) {
}

func (s *RlRenderSystem) Destroy() {
	rl.CloseWindow()
}
