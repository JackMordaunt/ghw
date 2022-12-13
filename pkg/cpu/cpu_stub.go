//go:build !linux && !windows
// +build !linux,!windows

// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package cpu

import (
	"fmt"
	"runtime"
)

func (i *Info) load() error {
	return fmt.Errorf("cpu.Info.load not implemented on " + runtime.GOOS)
}
