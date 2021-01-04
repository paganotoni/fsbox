package fsbox

import (
	"io/fs"
	"os"
	"path/filepath"
)

// filesystem allows us to wrap embedded file system search to
// use the file directory in development mode and with that allow
// changes on templates and other non-compiled resources to reflect
// in our application.
type filesystem struct {
	fs.FS
}

// Open wraps the fs.Open method and uses os.Open in case application
// is running in development mode (GOENV=development).
func (fsystem filesystem) Open(path string) (fs.File, error) {
	if os.Getenv("GO_ENV") == "development" {
		return os.Open(filepath.Join(path))
	}

	return fsystem.FS.Open(path)
}
