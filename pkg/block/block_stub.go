//go:build !linux && !darwin && !windows
// +build !linux,!darwin,!windows

// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package block

import (
	"fmt"
	"runtime"
)

func (i *Info) load() error {
	return fmt.Errorf("blockFillInfo not implemented on " + runtime.GOOS)
}
