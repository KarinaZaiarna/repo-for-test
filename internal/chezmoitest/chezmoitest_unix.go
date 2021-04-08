// +build !windows

package chezmoitest

import (
	"os"
	"syscall"
)

// Umask is the umask used in tests.
//
// If you change this then you will need to update the testscripts in
// testdata/scripts where permissions after applying umask are hardcoded as
// strings. Pure Go tests should use this value to ensure that they pass,
// irrespective of what it is set to. Be aware that the process's umask is a
// process-level property and cannot be locally changed within individual tests.
const Umask = os.FileMode(0o022)

func init() {
	syscall.Umask(int(Umask))
}
