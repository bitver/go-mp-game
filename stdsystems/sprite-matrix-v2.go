/*
This Source Code Form is subject to the terms of the Mozilla
Public License, v. 2.0. If a copy of the MPL was not distributed
with this file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package stdsystems

import (
	"gomp/pkg/ecs"
	"gomp/stdcomponents"
)

func NewSpriteMatrixV2System() SpriteMatrixV2System {
	return SpriteMatrixV2System{}
}

// SpriteMatrixSystem is a system that prepares SpriteSheet to be rendered
type SpriteMatrixV2System struct {
	SpriteMatrixes  *stdcomponents.SpriteMatrixV2ComponentManager
	SDLTextures     *stdcomponents.SDLTextureComponentManager
	AnimationStates *stdcomponents.AnimationStateComponentManager
	Positions       *stdcomponents.PositionComponentManager
}

func (s *SpriteMatrixV2System) Init() {}
func (s *SpriteMatrixV2System) Run() {
	s.SpriteMatrixes.EachEntityParallel(func(entity ecs.Entity) bool {
		spriteMatrix := s.SpriteMatrixes.Get(entity)    //
		animationState := s.AnimationStates.Get(entity) //

		frame := spriteMatrix.Animations[*animationState].Frame

		tr := s.SDLTextures.Get(entity)
		if tr != nil {
			tr.Frame = frame
		}
		return true
	})
}
func (s *SpriteMatrixV2System) Destroy() {}
