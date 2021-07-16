package fsbox

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/paganotoni/fsbox/internal/env"
)

// filesystem allows us to wrap embedded file system search to
// use the file directory in development mode and with that allow
// changes on templates and other non-compiled resources to reflect
// in our application.
type filesystem struct {
	fs.FS

	options []Options
}

// Checks to see if the FS has been set with a given option.
func (fsystem filesystem) hasOption(opt Options) bool {
	for _, op := range fsystem.options {
		if op != opt {
			continue
		}

		return true
	}

	return false
}

// Open wraps the fs.Open method and uses os.Open in case application
// is running in development mode (GOENV=development).
func (fsystem filesystem) Open(path string) (fs.File, error) {
	// If OptionFSIgnoreGoEnv then we should go straight
	// to fs.Open without considering ENV.
	if fsystem.hasOption(OptionFSIgnoreGoEnv) {
		return fsystem.FS.Open(path)
	}

	if env.Current() != env.Development {
		return fsystem.FS.Open(path)
	}

	return os.Open(filepath.Join(path))

}
