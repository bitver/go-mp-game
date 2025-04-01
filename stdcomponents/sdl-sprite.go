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
	"image/color"
)

//TODO: rework sprite and textures system

type SDLSprite struct {
	Surface  *sdl.Surface
	Src      sdl.FRect
	Origin   vectors.Vec2
	Tint     color.RGBA
	Dst      sdl.FRect
	Rotation float32
	Tiled    bool
}

type SDLSpriteComponentManager = ecs.ComponentManager[SDLSprite]

func NewSDLSpriteComponentManager() SDLSpriteComponentManager {
	return ecs.NewComponentManager[SDLSprite](SDLSpriteComponentId)
}
