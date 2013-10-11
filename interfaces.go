// Copyright (C) 2013 Ryan Chew. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

package cbl

type IPreUpdatable interface {
	PreUpdate(gt GameTime)
}

type IUpdatable interface {
	Update(gt GameTime)
}

type IPostUpdatable interface {
	PostUpdate(gt GameTime)
}

type IPreDrawable interface {
	PreDraw(gt GameTime)
}

type IDrawable interface {
	Draw(gt GameTime)
}

type IPostDrawable interface {
	PostDraw(gt GameTime)
}

type IComponent interface {
	Initialise()
	Shutdown()
}
