package fsbox

import (
	"io/fs"
)

// New box with the given filesystem
func New(fsys fs.FS, prefix string) *box {
	return &box{
		fsys:   filesystem{fsys},
		prefix: prefix,
	}
}
