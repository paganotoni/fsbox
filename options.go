package fsbox

// Options for the box.
type Options string

const (
	//OptionFSIgnoreGoEnv allways uses the Open method of the filesystem
	OptionFSIgnoreGoEnv = Options("FS_IGNORE_GO_ENV")
)
