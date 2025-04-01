/*
This Source Code Form is subject to the terms of the Mozilla
Public License, v. 2.0. If a copy of the MPL was not distributed
with this file, You can obtain one at http://mozilla.org/MPL/2.0/.

===-===-===-===-===-===-===-===-===-===
Donations during this file development:
-===-===-===-===-===-===-===-===-===-===

none :)]

Thank you for your support!
*/

package systems

import (
	"gomp/examples/spaceship-sdl/components"
	"gomp/examples/spaceship-sdl/entities"
	"gomp/pkg/ecs"
	"gomp/stdcomponents"
	"time"
)

func NewAssteroddSystem() AssteroddSystem {
	return AssteroddSystem{}
}

type AssteroddSystem struct {
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

func (s *AssteroddSystem) Init() {
	//entities.CreateSpaceShip(entities.CreateSpaceShipManagers{
	//	EntityManager: s.EntityManager,
	//	Positions:     s.Positions,
	//	Rotations:     s.Rotations,
	//	Scales:        s.Scales,
	//	Velocities:    s.Velocities,
	//	//Sprites:          s.Sprites,
	//	SDLSprites:       s.SDLSprites,
	//	BoxColliders:     s.BoxColliders,
	//	RigidBodies:      s.RigidBodies,
	//	PlayerTags:       s.PlayerTags,
	//	KeyboardInput:    s.KeyboardInput,
	//	Hps:              s.Hps,
	//	Weapons:          s.Weapons,
	//	SpaceshipIntents: s.SpaceshipIntents,
	//	SoundEffects:     s.SoundEffects,
	//}, 300, 300, -44.9)
	entities.CreateSatellite(entities.CreateSatelliteManagers{
		EntityManager: s.EntityManager,
		Positions:     s.Positions,
		Rotations:     s.Rotations,
		Scales:        s.Scales,
		Velocities:    s.Velocities,
		//Sprites:       s.Sprites,
		SDLSprites:   s.SDLSprites,
		BoxColliders: s.BoxColliders,
		RigidBodies:  s.RigidBodies,
	}, 500, 500, 0)
	entities.CreateSpaceSpawner(entities.CreateSpaceSpawnerManagers{
		EntityManager: s.EntityManager,
		Positions:     s.Positions,
		Velocities:    s.Velocities,
		SpaceSpawners: s.SpaceSpawnerTags,
	}, 16, 100, 1000, time.Millisecond*200)

	wallManager := entities.CreateWallManagers{
		EntityManager: s.EntityManager,
		Positions:     s.Positions,
		Rotations:     s.Rotations,
		Scales:        s.Scales,
		BoxColliders:  s.BoxColliders,
		//Sprites:       s.Sprites,
		SDLSprites:  s.SDLSprites,
		WallTags:    s.WallTags,
		RigidBodies: s.RigidBodies,
	}
	entities.CreateWall(&wallManager, 0, -1000, 0, 5000, 1000)
	entities.CreateWall(&wallManager, 0, 5000, 0, 5000, 1000)
	entities.CreateWall(&wallManager, -1000, -1000, 0, 1000, 7000)
	entities.CreateWall(&wallManager, 5000, -1000, 0, 1000, 7000)

	manager := s.EntityManager.Create()
	s.SceneManager.Create(manager, components.AsteroidSceneManager{})
}
func (s *AssteroddSystem) Run(dt time.Duration) {
	s.PlayerTags.EachEntity(func(e ecs.Entity) bool {
		intents := s.SpaceshipIntents.Get(e)
		keyboardInput := s.KeyboardInput.Get(e)

		intents.MoveUp = keyboardInput.MoveUp
		intents.MoveDown = keyboardInput.MoveDown
		intents.RotateLeft = keyboardInput.RotateLeft
		intents.RotateRight = keyboardInput.RotateRight
		intents.Fire = keyboardInput.Fire
		return true
	})
	s.SceneManager.EachEntity(func(e ecs.Entity) bool {
		sceneManager := s.SceneManager.Get(e)
		s.PlayerTags.EachEntity(func(e ecs.Entity) bool {
			playerHp := s.Hps.Get(e)
			if playerHp == nil {
				return true
			}
			sceneManager.PlayerHp = playerHp.Hp
			return true
		})
		return true
	})
}
func (s *AssteroddSystem) Destroy() {}
