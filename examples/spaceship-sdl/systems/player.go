/*
This Source Code Form is subject to the terms of the Mozilla
Public License, v. 2.0. If a copy of the MPL was not distributed
with this file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package systems

import (
	"gomp/examples/spaceship-sdl/components"
	"gomp/examples/spaceship-sdl/entities"
	"gomp/examples/spaceship-sdl/sprites"
	"gomp/pkg/ecs"
	"gomp/stdcomponents"
)

func NewPlayerSystem() PlayerSystem {
	return PlayerSystem{}
}

type PlayerSystem struct {
	EntityManager    *ecs.EntityManager
	SpriteMatrixV2es *stdcomponents.SpriteMatrixV2ComponentManager
	Positions        *stdcomponents.PositionComponentManager
	Rotations        *stdcomponents.RotationComponentManager
	Scales           *stdcomponents.ScaleComponentManager
	Velocities       *stdcomponents.VelocityComponentManager
	AnimationPlayers *stdcomponents.AnimationPlayerComponentManager
	AnimationStates  *stdcomponents.AnimationStateComponentManager
	Tints            *stdcomponents.TintComponentManager
	Flips            *stdcomponents.FlipComponentManager
	HP               *components.HpComponentManager
	Controllers      *components.ControllerComponentManager
	KeyboardInput    *components.KeyboardInputComponentManager
	Renderables      *stdcomponents.RenderableComponentManager
	YSorts           *stdcomponents.YSortComponentManager
	RenderOrders     *stdcomponents.RenderOrderComponentManager
	BoxColliders     *stdcomponents.BoxColliderComponentManager
	GenericCollider  *stdcomponents.GenericColliderComponentManager
	Players          *components.PlayerTagComponentManager
	SDLSprites       *stdcomponents.SDLSpriteComponentManager
	RigidBodies      *stdcomponents.RigidBodyComponentManager

	PlayerTags       *components.PlayerTagComponentManager
	AsteroidTags     *components.AsteroidComponentManager
	BulletTags       *components.BulletTagComponentManager
	Hps              *components.HpComponentManager
	Weapons          *components.WeaponComponentManager
	SpaceshipIntents *components.SpaceshipIntentComponentManager
	SpaceSpawnerTags *components.SpaceSpawnerComponentManager
	Collisions       *stdcomponents.CollisionComponentManager
	SceneManager     *components.AsteroidSceneManagerComponentManager
	WallTags         *components.WallTagComponentManager
	SoundEffects     *components.SoundEffectsComponentManager
}

func (s *PlayerSystem) Init() {
	s.SpriteMatrixV2es.Create(sprites.PlayerSpriteSharedComponentId, sprites.PlayerSpriteMatrixV2)

	player := entities.CreatePlayer(entities.CreatePlayerManagers{
		World:            s.EntityManager,
		SpriteMatrixV2es: s.SpriteMatrixV2es,
		Positions:        s.Positions,
		Rotations:        s.Rotations,
		Scales:           s.Scales,
		Velocities:       s.Velocities,
		AnimationPlayers: s.AnimationPlayers,
		AnimationStates:  s.AnimationStates,
		Tints:            s.Tints,
		Flips:            s.Flips,
		Renderables:      s.Renderables,
		YSorts:           s.YSorts,
		RenderOrders:     s.RenderOrders,
		BoxColliders:     s.BoxColliders,
		SdlSprites:       s.SDLSprites,
		RigidBodies:      s.RigidBodies,
		PlayerTags:       s.Players,
		Hps:              s.HP,
		Weapons:          s.Weapons,
		SpaceshipIntents: s.SpaceshipIntents,
		SoundEffects:     s.SoundEffects,
		KeyboardInput:    s.KeyboardInput,
	}, 300, 600, 0)

	s.Controllers.Create(player.Entity, components.Controller{})
}
func (s *PlayerSystem) Run() {

	var speed float32 = 300

	for e := range s.Controllers.EachEntity {
		velocity := s.Velocities.Get(e)
		flip := s.Flips.Get(e)
		animationState := s.AnimationStates.Get(e)
		inputs := s.KeyboardInput.Get(e)

		velocity.X = 0
		velocity.Y = 0

		if inputs.Fire {
			*animationState = entities.PlayerStateJump
		} else {
			*animationState = entities.PlayerStateIdle
			if inputs.RotateLeft {
				*animationState = entities.PlayerStateWalk
				velocity.X = speed
				flip.X = false
			}
			if inputs.RotateRight {
				*animationState = entities.PlayerStateWalk
				velocity.X = -speed
				flip.X = true
			}
			if inputs.MoveUp {
				*animationState = entities.PlayerStateWalk
				velocity.Y = -speed
			}
			if inputs.MoveDown {
				*animationState = entities.PlayerStateWalk
				velocity.Y = speed
			}
		}

		if inputs.Delete {
			s.EntityManager.Delete(e)
		}
	}

}
func (s *PlayerSystem) Destroy() {}
