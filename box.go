// fsbox package contains the implementation of a gobuffalo/packd box
// that uses the io/fs package to embed files. It is dependent on Go 1.16.
package fsbox

import (
	"errors"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/packd"
)

// ErrCannotAdd is returned if you want to add to box
var ErrCannotAdd = errors.New("stdbox does not allow to add")

// Ensuring the box meets packd.Box and with that
// that we can use it to replace our packr.Box
var _ packd.Box = (*box)(nil)

// box implementation for buffalo's Packd, this box aims to replace Packr Boxes
// by using the new Go 1.16 fs.
type box struct {
	fsys   fs.FS
	prefix string
}

// New box with the given filesystem
func New(fsys fs.FS, prefix string) *box {
	return &box{
		fsys:   fsys,
		prefix: prefix,
	}
}

func (bx *box) Open(path string) (http.File, error) {
	f, err := bx.fsys.Open(bx.PathFor(path))
	if err != nil {
		return nil, err
	}

	return packd.NewFile(path, f)
}

// AddString is not allowed for box, only here to meet packd.Box
func (bx *box) AddString(path string, t string) error {
	return ErrCannotAdd
}

// AddBytes is not allowed for box, only here to meet packd.Box
func (bx *box) AddBytes(path string, t []byte) error {
	return ErrCannotAdd
}

func (bx *box) List() []string {
	result := []string{}
	_ = fs.WalkDir(bx.fsys, bx.prefix, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}

		result = append(result, path)
		return nil
	})

	return result
}

func (bx *box) Find(path string) ([]byte, error) {
	result, err := fs.ReadFile(bx.fsys, bx.PathFor(path))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (bx *box) FindString(name string) (string, error) {
	b, err := bx.Find(name)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (bx *box) Walk(wf packd.WalkFunc) error {
	return fs.WalkDir(bx.fsys, bx.prefix, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			dir, err := packd.NewDir(path)
			if err != nil {
				return err
			}

			return wf(path, dir)
		}

		f, err := bx.fsys.Open(path)
		if err != nil {
			return err
		}

		file, err := packd.NewFile(path, f)
		if err != nil {
			return err
		}

		return wf(path, file)
	})
}

func (bx *box) WalkPrefix(prefix string, wf packd.WalkFunc) error {
	return fs.WalkDir(bx.fsys, bx.prefix, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasPrefix(path, prefix) {
			return nil
		}

		if d.IsDir() {
			dir, err := packd.NewDir(path)
			if err != nil {
				return err
			}

			return wf(path, dir)
		}

		f, err := bx.fsys.Open(path)
		if err != nil {
			return err
		}

		file, err := packd.NewFile(path, f)
		if err != nil {
			return err
		}

		return wf(path, file)
	})
}

func (bx *box) Has(path string) bool {
	matches, err := fs.Glob(bx.fsys, bx.PathFor(path))
	if err != nil {
		return false
	}

	found := len(matches) > 0
	return found
}

func (bx *box) PathFor(base string) string {
	if strings.HasPrefix(base, bx.prefix) {
		return base
	}

	return filepath.Join(bx.prefix, base)
}
