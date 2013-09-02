// Copyright (C) 2013 Ryan Chew. All rights reserved.
// Use of this source code is governed by a Apache 2.0
// license that can be found in the LICENSE file.

package cbl

import (
	"math"
)

const (
	EPSILON = 0.0001
)

func failNear(lhs, rhs float32) bool {
	return math.Abs(float64(lhs - rhs)) > EPSILON
}