// Copyright (C) 2013 Ryan Chew. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

package cbl

import (
	"time"
)

type GameTime struct {
	Elapsed time.Duration
	Total   time.Duration
	ElapsedReal time.Duration
	TotalReal time.Duration
	IsRunningSlowly bool
}
