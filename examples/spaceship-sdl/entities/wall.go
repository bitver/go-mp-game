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

package entities

import (
	"github.com/jupiterrider/purego-sdl3/sdl"
	"gomp/examples/spaceship-sdl/assets"
	"gomp/examples/spaceship-sdl/components"
	"gomp/examples/spaceship-sdl/config"
	"gomp/pkg/ecs"
	"gomp/stdcomponents"
	"gomp/vectors"
	"math"
)

type CreateWallManagers struct {
	EntityManager *ecs.EntityManager
	Positions     *stdcomponents.PositionComponentManager
	Rotations     *stdcomponents.RotationComponentManager
	Scales        *stdcomponents.ScaleComponentManager
	BoxColliders  *stdcomponents.BoxColliderComponentManager
	//Sprites       *stdcomponents.SpriteComponentManager
	SDLSprites  *stdcomponents.SDLSpriteComponentManager
	WallTags    *components.WallTagComponentManager
	RigidBodies *stdcomponents.RigidBodyComponentManager
}

func CreateWall(
	props *CreateWallManagers,
	posX, posY float32,
	angle float64,
	width, height float32,
) ecs.Entity {
	entity := props.EntityManager.Create()
	props.Positions.Create(entity, stdcomponents.Position{
		XY: vectors.Vec2{
			X: posX,
			Y: posY,
		},
	})
	props.Rotations.Create(entity, stdcomponents.Rotation{}.SetFromDegrees(angle))
	props.Scales.Create(entity, stdcomponents.Scale{
		XY: vectors.Vec2{
			X: 1,
			Y: 1,
		},
	})
	props.BoxColliders.Create(entity, stdcomponents.BoxCollider{
		WH: vectors.Vec2{
			X: width,
			Y: height,
		},
		Offset: vectors.Vec2{
			X: 0,
			Y: 0,
		},
		Layer: config.WallCollisionLayer,
		Mask:  0,
	})
	props.RigidBodies.Create(entity, stdcomponents.RigidBody{
		IsStatic: true,
		Mass:     math.MaxFloat32,
	})
	//props.Sprites.Create(entity, stdcomponents.Sprite{
	//	Texture: assets.Textures.Get("wall.png"),
	//	Frame: rl.Rectangle{
	//		X:      0,
	//		Y:      0,
	//		Width:  width,
	//		Height: height,
	//	},
	//	Origin: rl.Vector2{
	//		X: 0,
	//		Y: 0,
	//	},
	//	Tint: color.RGBA{
	//		R: 255,
	//		G: 255,
	//		B: 255,
	//		A: 255,
	//	},
	//})
	surface := assets.SDLSpritesGet("wall.png")
	props.SDLSprites.Create(entity, stdcomponents.SDLSprite{
		Surface: surface,
		Origin:  vectors.Vec2{X: 0, Y: 0},
		Src: sdl.FRect{
			W: float32(surface.W), H: float32(surface.H),
		},
		Dst: sdl.FRect{
			W: width, H: height,
		},
		Tiled: true,
	})
	props.WallTags.Create(entity, components.Wall{})

	return entity
}
