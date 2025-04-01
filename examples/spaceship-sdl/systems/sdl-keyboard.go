/*
This Source Code Form is subject to the terms of the Mozilla
Public License, v. 2.0. If a copy of the MPL was not distributed
with this file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package systems

import (
	"github.com/jupiterrider/purego-sdl3/sdl"
	"gomp/examples/spaceship-sdl/components"
)

func NewSDLKeyboardSystem() SDLKeyboardSystem {
	return SDLKeyboardSystem{}
}

type SDLKeyboardSystem struct {
	KeyboardInput *components.KeyboardInputComponentManager
}

func (s *SDLKeyboardSystem) Init() {}

func (s *SDLKeyboardSystem) Prerender() {

}

// Run
// THIS IS WRONG
// TODO: After rewrite of hierarchy of game, this should use sdl.PollEvent() instead
func (s *SDLKeyboardSystem) Run() {
	var input components.KeyboardInput
	var state = sdl.GetKeyboardState()
	if state[sdl.ScancodeEscape] {
		input.Escape = true
	}
	if state[sdl.ScancodeW] {
		input.MoveUp = true
	}
	if state[sdl.ScancodeS] {
		input.MoveDown = true
	}
	if state[sdl.ScancodeA] {
		input.RotateLeft = true
	}
	if state[sdl.ScancodeD] {
		input.RotateRight = true
	}
	if state[sdl.ScancodeSpace] {
		input.Fire = true
	}
	if state[sdl.ScancodeF12] {
		input.Debug = true
	}
	if state[sdl.ScancodeK] {
		input.Delete = true
	}

	s.KeyboardInput.EachComponent(func(ki *components.KeyboardInput) bool {
		ki.MoveUp = input.MoveUp
		ki.MoveDown = input.MoveDown
		ki.RotateLeft = input.RotateLeft
		ki.RotateRight = input.RotateRight
		ki.Fire = input.Fire
		ki.Debug = input.Debug
		ki.Escape = input.Escape
		return true
	})
}

func (s *SDLKeyboardSystem) Destroy() {}
