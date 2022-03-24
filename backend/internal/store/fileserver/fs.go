package fileserver

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
	"io/fs"
	"os"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
	"github.com/pkg/errors"
)

// OnDisk returns a new foodtinder.AssetServer that points to an on-disk
// directory.
func OnDisk(path string) foodtinder.FileServer {
	return diskAssets(path)
}

// diskAssets is a path that points to the directory to read assets from. It
// partially implements the foodtinder.Server interface.
type diskAssets string

func (s diskAssets) Open(name string) (fs.File, error) {
	return os.DirFS(string(s)).Open(name)
}

func (s diskAssets) Create(_ *foodtinder.Session, r io.Reader) (string, error) {
	f, err := os.CreateTemp(string(s), ".uploading.*")
	if err != nil {
		return "", errors.Wrap(err, "mktemp failed")
	}
	defer os.Remove(f.Name())
	defer f.Close()

	hasher := sha256.New()
	writer := io.MultiWriter(f, hasher)

	if _, err := io.Copy(writer, r); err != nil {
		return "", errors.Wrap(err, "cannot store file")
	}

	name := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	if err := os.Rename(f.Name(), name); err != nil {
		return "", errors.Wrap(err, "cannot commit stored file")
	}

	return name, nil
}
