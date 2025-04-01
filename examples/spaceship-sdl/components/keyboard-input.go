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

package components

import "gomp/pkg/ecs"

type keyboardEvent int32

const (
	EventNone keyboardEvent = iota
	EventMoveUp
	EventMoveDown
	EventRotateLeft
	EventRotateRight
	EventFire
	EventDebug
)

type KeyboardInput struct {
	MoveUp      bool
	MoveDown    bool
	RotateLeft  bool
	RotateRight bool
	Fire        bool
	Debug       bool
	Pprof       bool
	Escape      bool
	Delete      bool
}

type KeyboardInputComponentManager = ecs.ComponentManager[KeyboardInput]

func NewKeyboardInputComponentManager() KeyboardInputComponentManager {
	return ecs.NewComponentManager[KeyboardInput](KeyboardInputComponentId)
}
