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

package gomp

import "time"

type GameSystem interface {
	Init()
	Run(dt time.Duration)
	Destroy()
}

type SceneSystem interface {
	Init()
	Run(dt time.Duration)
	Destroy()
}

type RenderSystem interface {
	Init()
	Prerender() bool
	Render(dt time.Duration)
	Destroy()
}
