// +build !linux,!windows,!darwin,!freebsd,!netbsd
package version

import "runtime"

const (
	OperatingSystem = runtime.GOOS
)
