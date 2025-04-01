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

package stdcomponents

import (
	"github.com/jupiterrider/purego-sdl3/sdl"
	"gomp/pkg/ecs"
	"gomp/vectors"
)

type SpriteMatrixV2Animation struct {
	Name        string
	Frame       sdl.FRect
	NumOfFrames uint8
	Vertical    bool
	Loop        bool
}

type SpriteMatrixV2 struct {
	Origin     vectors.Vec2
	FPS        int32
	Animations []SpriteMatrixV2Animation
}

type SpriteMatrixV2ComponentManager = ecs.SharedComponentManager[SpriteMatrixV2]

func NewSpriteMatrixV2ComponentManager() SpriteMatrixV2ComponentManager {
	return ecs.NewSharedComponentManager[SpriteMatrixV2](SpriteMatrixV2ComponentId)
}
