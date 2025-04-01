/*
This Source Code Form is subject to the terms of the Mozilla
Public License, v. 2.0. If a copy of the MPL was not distributed
with this file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package entities

import (
	"github.com/jupiterrider/purego-sdl3/sdl"
	"gomp/examples/spaceship-sdl/assets"
	"gomp/examples/spaceship-sdl/components"
	"gomp/examples/spaceship-sdl/config"
	"gomp/examples/spaceship-sdl/sprites"
	"gomp/pkg/ecs"
	"gomp/stdcomponents"
	"gomp/vectors"
	"time"
)

const (
	PlayerStateIdle stdcomponents.AnimationState = iota
	PlayerStateWalk
	PlayerStateJump
	PlayerStateFall
	PlayerStateAttack
	PlayerStateHurt
	PlayerStateDie
)

type Player struct {
	ecs.Entity
	Position        *stdcomponents.Position
	Rotation        *stdcomponents.Rotation
	Scale           *stdcomponents.Scale
	Velocity        *stdcomponents.Velocity
	SpriteMatrixV2  *stdcomponents.SpriteMatrixV2
	SDLSprite       *stdcomponents.SDLSprite
	Tint            *stdcomponents.Tint
	AnimationPlayer *stdcomponents.AnimationPlayer
	AnimationState  *stdcomponents.AnimationState
	Flip            *stdcomponents.Flip
	Renderable      *stdcomponents.Renderable
	YSort           *stdcomponents.YSort
	RenderOrder     *stdcomponents.RenderOrder
	ColliderBox     *stdcomponents.BoxCollider

	RigidBodies      *stdcomponents.RigidBodyComponentManager
	PlayerTags       *components.PlayerTagComponentManager
	AsteroidTags     *components.AsteroidComponentManager
	BulletTags       *components.BulletTagComponentManager
	Hps              *components.HpComponentManager
	Weapons          *components.WeaponComponentManager
	SpaceshipIntents *components.SpaceshipIntentComponentManager
	KeyboardInput    *components.KeyboardInputComponentManager
	SpaceSpawnerTags *components.SpaceSpawnerComponentManager
	Collisions       *stdcomponents.CollisionComponentManager
	SceneManager     *components.AsteroidSceneManagerComponentManager
	WallTags         *components.WallTagComponentManager
	SoundEffects     *components.SoundEffectsComponentManager
}

type CreatePlayerManagers struct {
	World            *ecs.EntityManager
	SpriteMatrixV2es *stdcomponents.SpriteMatrixV2ComponentManager
	Positions        *stdcomponents.PositionComponentManager
	Rotations        *stdcomponents.RotationComponentManager
	Scales           *stdcomponents.ScaleComponentManager
	Velocities       *stdcomponents.VelocityComponentManager
	AnimationPlayers *stdcomponents.AnimationPlayerComponentManager
	AnimationStates  *stdcomponents.AnimationStateComponentManager
	Tints            *stdcomponents.TintComponentManager
	Flips            *stdcomponents.FlipComponentManager
	Renderables      *stdcomponents.RenderableComponentManager
	YSorts           *stdcomponents.YSortComponentManager
	RenderOrders     *stdcomponents.RenderOrderComponentManager
	BoxColliders     *stdcomponents.BoxColliderComponentManager
	SdlSprites       *stdcomponents.SDLSpriteComponentManager

	RigidBodies      *stdcomponents.RigidBodyComponentManager
	PlayerTags       *components.PlayerTagComponentManager
	AsteroidTags     *components.AsteroidComponentManager
	BulletTags       *components.BulletTagComponentManager
	Hps              *components.HpComponentManager
	Weapons          *components.WeaponComponentManager
	SpaceshipIntents *components.SpaceshipIntentComponentManager
	KeyboardInput    *components.KeyboardInputComponentManager
	SpaceSpawnerTags *components.SpaceSpawnerComponentManager
	Collisions       *stdcomponents.CollisionComponentManager
	SceneManager     *components.AsteroidSceneManagerComponentManager
	WallTags         *components.WallTagComponentManager
	SoundEffects     *components.SoundEffectsComponentManager
}

func CreatePlayer(
	props CreatePlayerManagers,
	posX, posY float32,
	angle float64,
) (player Player) {
	// Creating new player
	entity := props.World.Create()
	player.Entity = entity

	// Adding position component
	t := stdcomponents.Position{
		XY: vectors.Vec2{
			X: posX,
			Y: posY,
		},
	}
	player.Position = props.Positions.Create(entity, t)

	// Adding rotation component
	rotation := stdcomponents.Rotation{
		Angle: angle,
	}
	player.Rotation = props.Rotations.Create(entity, rotation)

	// Adding scale component
	scale := stdcomponents.Scale{
		XY: vectors.Vec2{
			X: 0.5,
			Y: 0.5,
		},
	}
	player.Scale = props.Scales.Create(entity, scale)

	// Adding velocity component
	velocity := stdcomponents.Velocity{}
	player.Velocity = props.Velocities.Create(entity, velocity)

	// Adding Tint component
	tint := stdcomponents.Tint{R: 255, G: 255, B: 255, A: 255}
	player.Tint = props.Tints.Create(entity, tint)

	// Adding sprite matrix component
	player.SpriteMatrixV2 = props.SpriteMatrixV2es.Set(entity, sprites.PlayerSpriteSharedComponentId)

	props.RigidBodies.Create(entity, stdcomponents.RigidBody{
		IsStatic: false,
		Mass:     2,
	})
	var width, height float32 = 96, 128
	props.BoxColliders.Create(entity, stdcomponents.BoxCollider{
		WH: vectors.Vec2{
			X: width * 0.8,
			Y: height * 0.8,
		},
		Offset: vectors.Vec2{
			X: width * 0.8 / 2,
			Y: height * 0.8 / 2,
		},
		Layer: config.PlayerCollisionLayer,
		Mask:  1<<config.EnemyCollisionLayer | 1<<config.WallCollisionLayer | 1<<config.BulletCollisionLayer,
	})

	surface := assets.SDLSpritesGet("milansheet.png")
	src := sprites.PlayerSpriteMatrixV2.Animations[0].Frame
	widthCoef := src.W / width
	heightCoef := src.H / height
	src.W = src.W * widthCoef
	src.H = src.H * heightCoef
	player.SDLSprite = props.SdlSprites.Create(entity, stdcomponents.SDLSprite{
		Surface: surface,
		Origin: vectors.Vec2{
			X: sprites.PlayerSpriteMatrixV2.Animations[0].Frame.W * widthCoef / 2,
			Y: sprites.PlayerSpriteMatrixV2.Animations[0].Frame.H * heightCoef / 2,
		},
		Src: src,
		Dst: sdl.FRect{
			X: 0,
			Y: 0,
			W: sprites.PlayerSpriteMatrixV2.Animations[0].Frame.W * widthCoef,
			H: sprites.PlayerSpriteMatrixV2.Animations[0].Frame.H * heightCoef,
		},
	})
	// Adding animation player component
	player.AnimationPlayer = props.AnimationPlayers.Create(entity, stdcomponents.AnimationPlayer{})

	// Adding Animation state component
	player.AnimationState = props.AnimationStates.Create(entity, PlayerStateWalk)

	// Adding Flip component
	player.Flip = props.Flips.Create(entity, stdcomponents.Flip{})

	// Adding renderable component
	//player.Renderable = props.Renderables.Create(entity, stdcomponents.SpriteMatrixV2RenderableType)

	// Adding YSort component
	player.YSort = props.YSorts.Create(entity, stdcomponents.YSort{})

	// Adding RenderOrder component
	player.RenderOrder = props.RenderOrders.Create(entity, stdcomponents.RenderOrder{})

	props.PlayerTags.Create(entity, components.PlayerTag{})
	props.KeyboardInput.Create(entity, components.KeyboardInput{})
	props.Hps.Create(entity, components.Hp{
		Hp:    3,
		MaxHp: 3,
	})

	props.Weapons.Create(entity, components.Weapon{
		Damage:       1,
		Cooldown:     time.Millisecond * 100,
		CooldownLeft: 0,
	})

	props.SpaceshipIntents.Create(entity, components.SpaceshipIntent{})
	props.SoundEffects.Create(entity, components.SoundEffect{
		Clip:      assets.Audio.Get("fly_sound.wav"),
		IsPlaying: false,
		IsLooping: true,
		Pitch:     1.0,
		Volume:    0.5,
		Pan:       0.5,
	})

	return player
}
