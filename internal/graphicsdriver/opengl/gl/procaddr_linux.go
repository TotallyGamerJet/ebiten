// SPDX-License-Identifier: MIT

package gl

import (
	"github.com/ebitengine/purego"
)

var opengl = purego.Dlopen("libGL.so", purego.RTLD_GLOBAL)

func getProcAddress(name string) uintptr {
	return purego.Dlsym(opengl, name)
}
