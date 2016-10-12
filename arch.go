// +build !386,!amd64,!arm
package version

import "runtime"

const (
	Architecture = runtime.GOARCH
)
