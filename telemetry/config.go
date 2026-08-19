package telemetry

import (
	"runtime/debug"
	"sync"
)

var (
	Name               = "github.com/foomo/keel/telemetry"
	DefaultServiceName = "undefined"
)

// Version returns the keel module version for use as the OTel instrumentation
// scope version. Resolved once from build info; empty string if unavailable.
var Version = sync.OnceValue(func() string {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}

	if bi.Main.Path == Name {
		return bi.Main.Version
	}

	for _, d := range bi.Deps {
		if d.Path == Name {
			return d.Version
		}
	}

	return ""
})
