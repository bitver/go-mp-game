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
	"time"
)

type CreateSpaceShipManagers struct {
	EntityManager *ecs.EntityManager
	Positions     *stdcomponents.PositionComponentManager
	Rotations     *stdcomponents.RotationComponentManager
	Scales        *stdcomponents.ScaleComponentManager
	Velocities    *stdcomponents.VelocityComponentManager
	//Sprites       *stdcomponents.SpriteComponentManager
	SDLSprites   *stdcomponents.SDLSpriteComponentManager
	BoxColliders *stdcomponents.BoxColliderComponentManager
	RigidBodies  *stdcomponents.RigidBodyComponentManager

	PlayerTags       *components.PlayerTagComponentManager
	Hps              *components.HpComponentManager
	Weapons          *components.WeaponComponentManager
	SpaceshipIntents *components.SpaceshipIntentComponentManager
	SoundEffects     *components.SoundEffectsComponentManager
	KeyboardInput    *components.KeyboardInputComponentManager
}

func CreateSpaceShip(
	props CreateSpaceShipManagers,
	posX, posY float32,
	angle float64,
) ecs.Entity {
	spaceShip := props.EntityManager.Create()

	props.Positions.Create(spaceShip, stdcomponents.Position{
		XY: vectors.Vec2{
			X: posX,
			Y: posY,
		},
	})

	props.Rotations.Create(spaceShip, stdcomponents.Rotation{}.SetFromDegrees(angle))

	props.Scales.Create(spaceShip, stdcomponents.Scale{
		XY: vectors.Vec2{
			X: 1,
			Y: 1,
		},
	})

	props.Velocities.Create(spaceShip, stdcomponents.Velocity{
		X: 0,
		Y: 0,
	})

	//props.Sprites.Create(spaceShip, stdcomponents.Sprite{
	//	Texture: assets.Textures.Get("ship_E.png"),
	//	Origin:  rl.Vector2{X: 32, Y: 40},
	//	Frame:   rl.Rectangle{0, 0, 64, 64},
	//	Tint:    color.RGBA{255, 255, 255, 255},
	//})
	props.SDLSprites.Create(spaceShip, stdcomponents.SDLSprite{
		Surface: assets.SDLSpritesGet("ship_E.png"),
		//TODO: why 40?
		Origin: vectors.Vec2{X: 32, Y: 40},
		Src:    sdl.FRect{X: 0, Y: 0, W: 64, H: 64},
		Dst:    sdl.FRect{X: 0, Y: 0, W: 64, H: 64},
	})
	props.BoxColliders.Create(spaceShip, stdcomponents.BoxCollider{
		WH: vectors.Vec2{
			X: 32,
			Y: 32,
		},
		Offset: vectors.Vec2{
			X: 16,
			Y: 16,
		},
		Layer: config.PlayerCollisionLayer,
		Mask:  1<<config.EnemyCollisionLayer | 1<<config.WallCollisionLayer | 1<<config.BulletCollisionLayer,
	})

	props.RigidBodies.Create(spaceShip, stdcomponents.RigidBody{
		IsStatic: false,
		Mass:     2,
	})
	props.PlayerTags.Create(spaceShip, components.PlayerTag{})
	props.KeyboardInput.Create(spaceShip, components.KeyboardInput{})
	props.Hps.Create(spaceShip, components.Hp{
		Hp:    3,
		MaxHp: 3,
	})

	props.Weapons.Create(spaceShip, components.Weapon{
		Damage:       1,
		Cooldown:     time.Millisecond * 100,
		CooldownLeft: 0,
	})

	props.SpaceshipIntents.Create(spaceShip, components.SpaceshipIntent{})
	props.SoundEffects.Create(spaceShip, components.SoundEffect{
		Clip:      assets.Audio.Get("fly_sound.wav"),
		IsPlaying: false,
		IsLooping: true,
		Pitch:     1.0,
		Volume:    1.0,
		Pan:       0.5,
	})

	return spaceShip
}
