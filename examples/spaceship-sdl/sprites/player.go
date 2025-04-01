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

package sprites

import (
	"github.com/jupiterrider/purego-sdl3/sdl"
	"gomp/stdcomponents"
	"gomp/vectors"
)

var PlayerSpriteMatrixV2 = stdcomponents.SpriteMatrixV2{
	Origin: vectors.Vec2{X: 0.5, Y: 0.5},
	FPS:    12,
	Animations: []stdcomponents.SpriteMatrixV2Animation{
		{
			Name:        "idle",
			Frame:       sdl.FRect{X: 0, Y: 0, W: 96, H: 128},
			NumOfFrames: 1,
			Vertical:    false,
			Loop:        true,
		},
		{
			Name:        "walk",
			Frame:       sdl.FRect{X: 0, Y: 512, W: 96, H: 128},
			NumOfFrames: 8,
			Vertical:    false,
			Loop:        true,
		},
		{
			Name:        "jump",
			Frame:       sdl.FRect{X: 96, Y: 0, W: 96, H: 128},
			NumOfFrames: 1,
			Vertical:    false,
			Loop:        false,
		},
	},
}
