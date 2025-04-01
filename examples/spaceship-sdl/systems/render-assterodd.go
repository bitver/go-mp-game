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

package systems

import (
	"fmt"
	"github.com/jupiterrider/purego-sdl3/sdl"
	"github.com/jupiterrider/purego-sdl3/ttf"
	"github.com/negrel/assert"
	"gomp"
	"gomp/examples/spaceship-sdl/components"
	"gomp/pkg/ecs"
	"gomp/stdcomponents"
	"gomp/vectors"
	"log"
	"math"
	"slices"
	"sync"
	"time"
)

func NewRenderAssteroddSystem() RenderAssteroddSystem {
	return RenderAssteroddSystem{
		instanceData: make([]stdcomponents.RLTexturePro, 0, 8192),
	}
}

type RenderAssteroddSystem struct {
	EntityManager                      *ecs.EntityManager
	RlTexturePros                      *stdcomponents.RLTextureProComponentManager
	Positions                          *stdcomponents.PositionComponentManager
	Rotations                          *stdcomponents.RotationComponentManager
	Scales                             *stdcomponents.ScaleComponentManager
	AnimationPlayers                   *stdcomponents.AnimationPlayerComponentManager
	Tints                              *stdcomponents.TintComponentManager
	Flips                              *stdcomponents.FlipComponentManager
	AnimationStates                    *stdcomponents.AnimationStateComponentManager
	Sprites                            *stdcomponents.SpriteComponentManager
	SpriteMatrixV2es                   *stdcomponents.SpriteMatrixV2ComponentManager
	RenderOrders                       *stdcomponents.RenderOrderComponentManager
	BoxColliders                       *stdcomponents.BoxColliderComponentManager
	CircleColliders                    *stdcomponents.CircleColliderComponentManager
	AABBs                              *stdcomponents.AABBComponentManager
	Collisions                         *stdcomponents.CollisionComponentManager
	ColliderSleepStateComponentManager *stdcomponents.ColliderSleepStateComponentManager
	renderList                         []sdlRenderEntry
	instanceData                       []stdcomponents.RLTexturePro
	//TODO: camera system? camera helper? something with camera!
	camera        sdl.FRect
	SceneManager  *components.AsteroidSceneManagerComponentManager
	KeyboardInput *components.KeyboardInputComponentManager
	SDLTextures   *stdcomponents.SDLTextureComponentManager
	SDLSprites    *stdcomponents.SDLSpriteComponentManager

	monitorWidth  int32
	monitorHeight int32

	PlayerTag    *components.PlayerTagComponentManager
	debug        bool
	font         *ttf.Font
	textures     map[*sdl.Surface]*sdl.Texture
	renderer     *sdl.Renderer
	window       *sdl.Window
	renderSystem *gomp.SDLRender
	textTexture  *sdl.Texture
}

type sdlRenderEntry struct {
	Texture  *stdcomponents.SDLTexture
	Position sdl.FPoint
	Scale    vectors.Vec2
	ZIndex   float32
	Rotation float64
}

// Init
// TODO: Complete rewrite of hierarchy of the game
// temporary workaround
func (s *RenderAssteroddSystem) Init() {
	s.renderer = s.renderSystem.Renderer
	s.window = s.renderSystem.Window
	s.textures = make(map[*sdl.Surface]*sdl.Texture)

	var x, y int32
	sdl.GetWindowSizeInPixels(s.window, &x, &y)
	s.camera = sdl.FRect{
		X: float32(x / 2),
		Y: float32(y / 2),
		W: float32(x),
		H: float32(y),
	}

	//TODO: asset system
	s.font = ttf.OpenFont("./Roboto-SemiBold.ttf", 96)
	if s.font == nil {
		log.Fatal("Error loading font")
	}

	var text = "Game Over"
	textSurface := ttf.RenderTextBlended(s.font, text, uint64(len(text)), sdl.Color{R: 255, G: 0, B: 0, A: 255})
	s.textTexture = sdl.CreateTextureFromSurface(s.renderer, textSurface)
	s.monitorWidth = x
	s.monitorHeight = y
}

func (s *RenderAssteroddSystem) Run(dt time.Duration) bool {
	s.PlayerTag.EachEntity(func(entity ecs.Entity) bool {
		event := s.KeyboardInput.Get(entity)
		if event.Debug {
			s.debug = !s.debug
		}
		return true
	})

	s.loadTextures()
	s.prepareRender(dt)

	sdl.SetRenderDrawColor(s.renderer, 0, 0, 0, 255)
	sdl.RenderClear(s.renderer)

	s.render()

	//Render
	var fps = int(time.Second.Seconds() / dt.Seconds())
	sdl.SetRenderDrawColor(s.renderer, 200, 200, 200, 255)
	sdl.RenderDebugText(s.renderer, 10, 10, "FPS: "+fmt.Sprintf("%d", fps))
	sdl.RenderDebugText(s.renderer, 10, 30, fmt.Sprintf("%d entities", s.EntityManager.Size()))

	s.SceneManager.EachComponent(func(a *components.AsteroidSceneManager) bool {
		sdl.RenderDebugText(s.renderer, 10, 50, fmt.Sprintf("Player HP: %d", a.PlayerHp))
		sdl.RenderDebugText(s.renderer, 10, 70, fmt.Sprintf("Score: %d", a.PlayerScore))
		if a.PlayerHp <= 0 {

			var dstRect = sdl.FRect{
				X: float32(s.monitorWidth-s.textTexture.W) / 2,
				Y: float32(s.monitorHeight-s.textTexture.H) / 2,
				W: float32(s.textTexture.W),
				H: float32(s.textTexture.H),
			}
			sdl.RenderTexture(s.renderer, s.textTexture, nil, &dstRect)
		}
		return false
	})

	return true
}

func (s *RenderAssteroddSystem) Destroy() {
	sdl.DestroyTexture(s.textTexture)
	ttf.CloseFont(s.font)
}

func (s *RenderAssteroddSystem) render() {
	/// ==========
	// DEBUG
	// ==========
	if s.debug {
		s.BoxColliders.EachEntity(func(e ecs.Entity) bool {
			col := s.BoxColliders.Get(e)
			scale := s.Scales.Get(e)
			pos := s.Positions.Get(e)
			//SDL can't draw rotated rectangles
			//rot := s.Rotations.Get(e)

			//rl.DrawRectanglePro(rl.Rectangle{
			//	X:      pos.XY.X,
			//	Y:      pos.XY.Y,
			//	Width:  col.WH.X * scale.XY.X,
			//	Height: col.WH.Y * scale.XY.Y,
			//}, rl.Vector2{
			//	X: col.Offset.X * scale.XY.X,
			//	Y: col.Offset.Y * scale.XY.Y,
			//}, float32(rot.Degrees()), rl.DarkGreen)
			sdl.SetRenderDrawColor(s.renderer, 0, 255, 0, 255)
			sdl.RenderRect(s.renderer, &sdl.FRect{
				X: pos.XY.X - s.camera.X,
				Y: pos.XY.Y - s.camera.Y,
				W: col.WH.X * scale.XY.X,
				H: col.WH.Y * scale.XY.Y,
			})
			return true
		})
		s.CircleColliders.EachEntity(func(e ecs.Entity) bool {
			col := s.CircleColliders.Get(e)
			scale := s.Scales.Get(e)
			pos := s.Positions.Get(e)

			posWithOffset := pos.XY.Add(col.Offset.Mul(scale.XY))
			//rl.DrawCircle(int32(posWithOffset.X), int32(posWithOffset.Y), col.Radius*scale.XY.X, rl.DarkGreen)
			//SDL doesn't have a circle, so we draw a rectangle instead
			sdl.SetRenderDrawColor(s.renderer, 0, 255, 0, 255)
			sdl.RenderRect(s.renderer, &sdl.FRect{
				X: posWithOffset.X - col.Radius*scale.XY.X - s.camera.X,
				Y: posWithOffset.Y - col.Radius*scale.XY.X - s.camera.Y,
				W: col.Radius * scale.XY.X * 2,
				H: col.Radius * scale.XY.X * 2,
			})
			return true
		})
	}

	// Extract and sort entities
	if cap(s.renderList) < s.SDLTextures.Len() {
		s.renderList = append(s.renderList, make([]sdlRenderEntry, 0, s.SDLTextures.Len()-cap(s.renderList))...)
	}
	s.SDLSprites.EachEntity(func(entity ecs.Entity) bool {
		renderOrder := s.RenderOrders.Get(entity)
		sprite := s.SDLSprites.Get(entity)
		texture := s.SDLTextures.Get(entity)
		scale := s.Scales.Get(entity)
		positions := s.Positions.Get(entity)
		rotation := s.Rotations.Get(entity)
		//animationMatrix := s.AnimationStates.Get(entity)
		//animationPlayer := s.AnimationPlayers.Get(entity)
		if sprite != nil {
			s.renderList = append(s.renderList, sdlRenderEntry{
				Position: sdl.FPoint(positions.XY),
				Texture:  texture,
				Scale:    scale.XY,
				Rotation: rotation.Degrees(),
				ZIndex:   renderOrder.CalculatedZ,
			})
			return true
		}

		panic("Unknown renderable type")
	})

	slices.SortStableFunc(s.renderList, func(a, b sdlRenderEntry) int {
		return int(math.Floor(float64(a.ZIndex - b.ZIndex)))
	})

	s.submitBatch(s.renderList)
	s.renderList = s.renderList[:0]

	// ==========
	// DEBUG
	// ==========
	if s.debug {
		s.AABBs.EachEntity(func(e ecs.Entity) bool {
			aabb := s.AABBs.Get(e)
			//rl.DrawRectangleLines(int32(aabb.Min.X), int32(aabb.Min.Y), int32(aabb.Max.X-aabb.Min.X), int32(aabb.Max.Y-aabb.Min.Y), rl.Green)
			frect := sdl.FRect{
				X: aabb.Min.X - s.camera.X,
				Y: aabb.Min.Y - s.camera.Y,
				W: aabb.Max.X - aabb.Min.X,
				H: aabb.Max.Y - aabb.Min.Y,
			}
			sdl.SetRenderDrawColor(s.renderer, 0, 255, 255, 255)
			sdl.RenderRect(s.renderer, &frect)
			return true
		})
		s.Collisions.EachEntity(func(entity ecs.Entity) bool {
			pos := s.Positions.Get(entity)
			//rl.DrawRectangle(int32(pos.XY.X-8), int32(pos.XY.Y-8), 16, 16, rl.Red)
			frect := sdl.FRect{
				X: pos.XY.X - s.camera.X,
				Y: pos.XY.Y - s.camera.Y,
				W: 16,
				H: 16,
			}
			sdl.SetRenderDrawColor(s.renderer, 255, 255, 0, 255)
			sdl.RenderRect(s.renderer, &frect)
			return true
		})
		s.SDLTextures.EachEntity(func(e ecs.Entity) bool {
			position := s.Positions.Get(e)
			//rl.DrawRectangle(int32(position.XY.X-2), int32(position.XY.Y-2), 4, 4, rl.Red)
			frect := sdl.FRect{
				X: position.XY.X - s.camera.X,
				Y: position.XY.Y - s.camera.Y,
				W: 4,
				H: 4,
			}
			sdl.SetRenderDrawColor(s.renderer, 255, 0, 0, 255)
			sdl.RenderRect(s.renderer, &frect)
			return true
		})
	}

}

func (s *RenderAssteroddSystem) submitBatch(data []sdlRenderEntry) {
	for i := range data {
		//rl.DrawTexturePro(cameraTexture, data[i].Frame, data[i].Dest, data[i].Origin, data[i].Rotation, data[i].Tint)
		//TODO: camera system, camera wrapper, camera helper, leave as is?
		dest := &sdl.FRect{
			X: data[i].Position.X - data[i].Texture.Origin.X*data[i].Scale.X - s.camera.X,
			Y: data[i].Position.Y - data[i].Texture.Origin.Y*data[i].Scale.Y - s.camera.Y,
			W: (data[i].Texture.Dest.W) * data[i].Scale.X,
			H: (data[i].Texture.Dest.H) * data[i].Scale.Y,
		}

		if s.debug {
			sdl.SetRenderDrawColor(s.renderer, 12, 145, 234, 255)
			sdl.RenderRect(s.renderer, dest)
		}
		//Tiles
		//TODO: tiles system?
		//Vibee coded tile system
		if data[i].Texture.Tiled {
			var texW, texH float32
			sdl.GetTextureSize(data[i].Texture.Texture, &texW, &texH)

			// Use the 'dest' rectangle already defined above
			destRect := *dest
			horizontalRepeats := int32(destRect.W / texW)
			verticalRepeats := int32(destRect.H / texH)
			horizontalRemainder := int32(destRect.W) % int32(texW)
			verticalRemainder := int32(destRect.H) % int32(texH)

			tileRect := sdl.FRect{X: destRect.X, Y: destRect.Y, W: texW, H: texH}

			for y := int32(0); y < verticalRepeats; y++ {
				tileRect.Y = destRect.Y + float32(y)*texH
				for x := int32(0); x < horizontalRepeats; x++ {
					tileRect.X = destRect.X + float32(x)*texW
					sdl.RenderTexture(s.renderer, data[i].Texture.Texture, nil, &tileRect)
				}
				// Render the remainder on the x-axis
				tileRect.X = destRect.X + float32(horizontalRepeats)*texW
				tileRect.W = float32(horizontalRemainder)
				if horizontalRemainder > 0 {
					sdl.RenderTexture(s.renderer, data[i].Texture.Texture, nil, &tileRect)
				}
				tileRect.W = texW
			}

			// Render the remainder on the y-axis
			tileRect.Y = destRect.Y + float32(verticalRepeats)*texH
			tileRect.H = float32(verticalRemainder)
			for x := int32(0); x < horizontalRepeats; x++ {
				tileRect.X = destRect.X + float32(x)*texW
				sdl.RenderTexture(s.renderer, data[i].Texture.Texture, nil, &tileRect)
			}
			// Render the corner remainder
			tileRect.X = destRect.X + float32(horizontalRepeats)*texW
			tileRect.W = float32(horizontalRemainder)
			if horizontalRemainder > 0 && verticalRemainder > 0 {
				sdl.RenderTexture(s.renderer, data[i].Texture.Texture, nil, &tileRect)
			}
		} else {
			assert.True(data[i].Texture == nil && data[i].Texture.Texture == nil, "Texture is nil")
			sdl.RenderTextureRotated(s.renderer, data[i].Texture.Texture, &data[i].Texture.Frame, dest, data[i].Rotation, nil, 0)
		}
	}
}

func (s *RenderAssteroddSystem) getInstanceData(e ecs.Entity) stdcomponents.SDLTexture {
	return *s.SDLTextures.Get(e)
}

func (s *RenderAssteroddSystem) prepareRender(dt time.Duration) {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go s.prepareAnimations(wg)
	//wg.Add(1)
	//go s.prepareFlips(wg)
	wg.Add(1)
	go s.preparePositions(wg, dt)
	wg.Add(1)
	go s.prepareRotations(wg)
	//wg.Add(1)
	//go s.prepareScales(wg)
	wg.Add(1)
	go s.prepareTints(wg)
	wg.Wait()
}

func (s *RenderAssteroddSystem) prepareAnimations(wg *sync.WaitGroup) {
	defer wg.Done()
	s.SDLTextures.EachEntityParallel(func(entity ecs.Entity) bool {
		texture := s.SDLTextures.Get(entity)
		animation := s.AnimationPlayers.Get(entity)
		if animation == nil {
			return true
		}
		if animation.Vertical {
			texture.Frame.Y += texture.Frame.H * float32(animation.Current)
		} else {
			texture.Frame.X += texture.Frame.W * float32(animation.Current)
		}
		return true
	})
}

func (s *RenderAssteroddSystem) prepareFlips(wg *sync.WaitGroup) {
	defer wg.Done()
	s.SDLTextures.EachEntityParallel(func(entity ecs.Entity) bool {
		texturePro := s.SDLTextures.Get(entity)
		flipped := s.Flips.Get(entity)
		if flipped == nil {
			return true
		}
		if flipped.X {
			texturePro.Frame.W *= -1
		}
		if flipped.Y {
			texturePro.Frame.H *= -1
		}
		return true
	})
}

func (s *RenderAssteroddSystem) preparePositions(wg *sync.WaitGroup, dt time.Duration) {
	defer wg.Done()
	//dts := dt.Seconds()
	s.PlayerTag.EachEntity(func(e ecs.Entity) bool {
		player := s.Positions.Get(e)
		if player == nil {
			return true
		}
		//TODO: this action should be in global system reacting on resize event
		var w, h int32
		sdl.GetWindowSizeInPixels(s.window, &w, &h)

		s.camera.X = player.XY.X - float32(w)/2
		s.camera.Y = player.XY.Y - float32(h)/2
		return false
	})
	//s.SDLTextures.EachEntityParallel(func(entity ecs.Entity) bool {
	//	texture := s.SDLTextures.Get(entity)
	//	position := s.Positions.Get(entity)
	//	if position == nil {
	//		return true
	//	}
	//	decay := 40.0 // DECAY IS TICKRATE DEPENDENT
	//	x := float32(s.expDecay(float64(texture.Dest.X), float64(position.XY.X), decay, dts))
	//	y := float32(s.expDecay(float64(texture.Dest.Y), float64(position.XY.Y), decay, dts))
	//	texture.Dest.X = x
	//	texture.Dest.Y = y
	//	player := s.Player.Get(entity)
	//	if player != nil {
	//		s.camera.X = x
	//		s.camera.Y = y
	//	}
	//
	//	return true
	//})
}

func (s *RenderAssteroddSystem) prepareRotations(wg *sync.WaitGroup) {
	defer wg.Done()
	s.SDLTextures.EachEntityParallel(func(entity ecs.Entity) bool {
		texturePro := s.SDLTextures.Get(entity)
		rotation := s.Rotations.Get(entity)
		if rotation == nil {
			return true
		}
		texturePro.Rotation = rotation.Angle
		return true
	})
}

//func (s *RenderAssteroddSystem) prepareScales(wg *sync.WaitGroup) {
//	defer wg.Done()
//	s.SDLTextures.EachEntityParallel(func(entity ecs.Entity) bool {
//		texturePro := s.SDLTextures.Get(entity)
//		scale := s.Scales.Get(entity)
//		if scale == nil {
//			return true
//		}
//		texturePro.Dest.W *= scale.XY.X
//		texturePro.Dest.H *= scale.XY.Y
//		return true
//	})
//}

func (s *RenderAssteroddSystem) prepareTints(wg *sync.WaitGroup) {
	defer wg.Done()
	s.SDLTextures.EachEntityParallel(func(entity ecs.Entity) bool {
		tr := s.SDLTextures.Get(entity)
		tint := s.Tints.Get(entity)
		if tint == nil {
			return true
		}
		trTint := &tr.Tint
		trTint.A = tint.A
		trTint.R = tint.R
		trTint.G = tint.G
		trTint.B = tint.B
		return true
	})
}

func (s *RenderAssteroddSystem) expDecay(a, b, decay, dt float64) float64 {
	return b + (a-b)*(math.Exp(-decay*dt))
}

func (s *RenderAssteroddSystem) loadTextures() {
	s.SDLSprites.EachEntity(func(entity ecs.Entity) bool {
		sprite := s.SDLSprites.Get(entity)
		assert.True(sprite != nil, "Sprite is nil")

		texture := s.SDLTextures.Get(entity)
		if texture == nil {
			if s.textures[sprite.Surface] == nil {
				t := sdl.CreateTextureFromSurface(s.renderer, sprite.Surface)
				s.textures[sprite.Surface] = t
			}

			s.SDLTextures.Create(entity, stdcomponents.SDLTexture{
				Texture: s.textures[sprite.Surface],
				Origin:  sprite.Origin,
				Frame:   sprite.Src,
				Dest:    sprite.Dst,
				Tiled:   sprite.Tiled,
			})
			s.RenderOrders.Create(entity, stdcomponents.RenderOrder{})
		}
		return true
	})
}

func (s *RenderAssteroddSystem) SetRenderer(renderer *gomp.SDLRender) {
	s.renderSystem = renderer
}
