// SPDX-FileCopyrightText: 2021 Softbear, Inc.
// SPDX-License-Identifier: AGPL-3.0-or-later

package sector

import (
	"github.com/SoftbearStudios/mk48/server/world"
	"testing"
)

func BenchmarkSectorWorld(b *testing.B) {
	world.Bench(b, func(radius int) world.World {
		return New(float32(radius))
	}, 65536)
}
