package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

const batchSize = 1 << 14
const rectCounter = 1 << 18
const framerate = 600
const frameCount = 100 // Number of frames to average

func main() {
	f, err := os.Create("cpu.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	fmt.Println("CPU Profile Started")
	defer fmt.Println("CPU Profile Stopped")

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	ok := sdl.SetHint(sdl.HINT_RENDER_DRIVER, "gpu,software")
	if !ok {
		panic(sdl.GetError())
	}
	must(ttf.Init())
	defer ttf.Quit()
	must(sdl.Init(sdl.INIT_EVERYTHING))
	defer sdl.Quit()

	window, renderer, err := sdl.CreateWindowAndRenderer(1280, 720, sdl.WINDOW_RESIZABLE)
	must(err)
	defer window.Destroy()
	defer renderer.Destroy()

	font, err := ttf.OpenFont("Roboto-SemiBold.ttf", 14)
	must(err)
	defer font.Close()

	var renderTicker *time.Ticker
	if framerate > 0 {
		renderTicker = time.NewTicker(time.Second / time.Duration(framerate))
		defer renderTicker.Stop()
	}

	var dt time.Duration = time.Second
	var fps, avgFps float64
	updatedAt := time.Now()

	rects := make([]sdl.FRect, rectCounter)
	for i := range rects {
		rects[i] = sdl.FRect{X: 300, Y: 300, W: 10, H: 10}
	}

	frameTimes := make([]float64, 0, frameCount)

	err = initLetterAtlas(renderer, font)
	must(err)
	defer cleanupLetterAtlas()

Outer:
	for {
		if renderTicker != nil {
			<-renderTicker.C
		}
		fps = time.Second.Seconds() / dt.Seconds()
		dt = time.Since(updatedAt)
		updatedAt = time.Now()

		frameTimes = append(frameTimes, fps)
		if len(frameTimes) > frameCount {
			frameTimes = frameTimes[1:]
		}

		var totalFps float64
		for _, frameTime := range frameTimes {
			totalFps += frameTime
		}
		avgFps = totalFps / float64(len(frameTimes))

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				break Outer
			case *sdl.KeyboardEvent:
				if event.(*sdl.KeyboardEvent).Keysym.Scancode == sdl.SCANCODE_ESCAPE {
					break Outer
				}
			}
		}

		must(renderer.SetDrawColor(0, 0, 0, 255))
		must(renderer.Clear())

		must(renderer.SetDrawColor(255, 255, 255, 255))

		for i := 0; i < len(rects); i += batchSize {
			must(renderer.FillRectsF(rects[i : i+batchSize]))
		}

		must(drawText(renderer, 10, 10, fmt.Sprintf("FPS %f", fps)))
		must(drawText(renderer, 10, 30, fmt.Sprintf("Avg FPS %f", avgFps)))
		must(drawText(renderer, 10, 50, fmt.Sprintf("dt %s", dt.String())))

		renderer.Present()
	}
}

func must(err error) {
	if err != nil {
		println(err.Error())
		panic(err)
	}
}

// Single texture atlas approach
var letterAtlasTexture *sdl.Texture
var letterPositions []sdl.Rect

func initLetterAtlas(renderer *sdl.Renderer, font *ttf.Font) error {
	// Create an array sized to fit all printable ASCII characters
	letterPositions = make([]sdl.Rect, '~'-' '+1)

	// Define the characters to preload (ASCII printable range)
	chars := []rune{}
	for r := ' '; r <= '~'; r++ {
		chars = append(chars, r)
	}

	// First pass: render each character to measure sizes
	var totalWidth, maxHeight int32 = 0, 0
	tempSurfaces := make(map[rune]*sdl.Surface)

	for _, char := range chars {
		surface, err := font.RenderUTF8Blended(string(char), sdl.Color{R: 255, G: 255, B: 255, A: 255})
		if err != nil {
			return err
		}

		tempSurfaces[char] = surface
		totalWidth += surface.W
		if surface.H > maxHeight {
			maxHeight = surface.H
		}
	}

	// Create a single large surface for our atlas
	atlasSurface, err := sdl.CreateRGBSurface(0, totalWidth, maxHeight, 32,
		0x00FF0000, 0x0000FF00, 0x000000FF, 0xFF000000)
	if err != nil {
		return err
	}
	defer atlasSurface.Free()

	// Place each character in the atlas
	var xOffset int32 = 0
	for _, char := range chars {
		surface := tempSurfaces[char]

		// Record position in the atlas
		letterPositions[char-' '] = sdl.Rect{
			X: xOffset,
			Y: 0,
			W: surface.W,
			H: surface.H,
		}

		// Copy character to atlas
		srcRect := sdl.Rect{X: 0, Y: 0, W: surface.W, H: surface.H}
		dstRect := sdl.Rect{X: xOffset, Y: 0, W: surface.W, H: surface.H}

		err = surface.Blit(&srcRect, atlasSurface, &dstRect)
		if err != nil {
			return err
		}

		xOffset += surface.W
	}

	// Create texture from the atlas
	letterAtlasTexture, err = renderer.CreateTextureFromSurface(atlasSurface)
	if err != nil {
		return err
	}

	// Free temporary surfaces
	for _, surface := range tempSurfaces {
		surface.Free()
	}

	return nil
}

func drawText(renderer *sdl.Renderer, x, y int32, text string) error {
	curX := x

	for _, char := range text {
		// Check if character is in our range
		charIndex := int(char - ' ')
		if charIndex < 0 || charIndex >= len(letterPositions) {
			curX += 8 // Skip with a default width
			continue
		}

		pos := letterPositions[charIndex]

		src := &sdl.Rect{X: pos.X, Y: pos.Y, W: pos.W, H: pos.H}
		dst := &sdl.Rect{X: curX, Y: y, W: pos.W, H: pos.H}

		// Render the character from the atlas
		if err := renderer.Copy(letterAtlasTexture, src, dst); err != nil {
			return err
		}

		// Advance cursor position
		curX += pos.W
	}

	return nil
}

func cleanupLetterAtlas() {
	if letterAtlasTexture != nil {
		letterAtlasTexture.Destroy()
	}
}
