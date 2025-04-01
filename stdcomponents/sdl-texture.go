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

type SDLTexture struct {
	Texture  *sdl.Texture
	Frame    sdl.FRect
	Origin   vectors.Vec2
	Tint     color.RGBA
	Dest     sdl.FRect
	Rotation float64
	Tiled    bool
}

type SDLTextureComponentManager = ecs.ComponentManager[SDLTexture]

func NewSDLTextureComponentManager() SDLTextureComponentManager {
	return ecs.NewComponentManager[SDLTexture](SDLTextureComponentId)
}
