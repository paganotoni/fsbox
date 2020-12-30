package fsbox

import (
	"io/fs"
	"os"
)

// filesystem allows us to wrap embedded file system search to
// use the file directory in development mode and with that allow
// changes on templates and other non-compiled resources to reflect
// in our application.
type filesystem struct {
	fs.FS
}

// Open wraps the fs.Open method and uses os.Open in case application
// is not running in production mode (GOENV=production).
func (fsystem filesystem) Open(path string) (fs.File, error) {
	if os.Getenv("GO_ENV") == "production" {
		return fsystem.FS.Open(path)
	}

	return os.Open(path)
}
