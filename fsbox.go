package fsbox

import "io/fs"

// New box with the given filesystem
func New(fsys fs.FS, prefix string, options ...Options) *box {
	return &box{
		fsys:   filesystem{fsys, options},
		prefix: prefix,
	}
}
