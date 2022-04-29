package fileserver

import (
	"crypto/sha256"
	"encoding/base64"
	"image"
	"io"
	"io/fs"
	"os"

	_ "image/jpeg"
	_ "image/png"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
	"github.com/bbrks/go-blurhash"
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

// Blurhash creates the blur hash of the image asset with the given name. If the
// hash cannot be created for whatever reason, then an empty string is returned.
// The error may also be nil.
func Blurhash(fs foodtinder.FileServer, name string) (string, error) {
	f, err := fs.Open(name)
	if err != nil {
		return "", err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return "", nil
	}

	hash, _ := blurhash.Encode(5, 5, img)
	return hash, nil
}
