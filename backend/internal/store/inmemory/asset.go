package inmemory

import (
	"bytes"
	"io"
	"io/fs"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
)

type assetServer server

var _ fs.FS = (*assetServer)(nil)

func (s *assetServer) Open(name string) (fs.File, error) {
	s.mu.RLock()
	b, ok := s.store.Assets[name]
	s.mu.RUnlock()

	if ok {
		return &byteFile{
			name,
			bytes.NewReader(b),
		}, nil
	}

	return nil, fs.ErrNotExist
}

type assetUploadServer authorizedServer

var _ foodtinder.AssetUploadServer = (*assetUploadServer)(nil)

func (s *assetUploadServer) Upload(r io.Reader) (string, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	id := atomic.AddUint64(&s.assetIx, 1)
	name := strconv.FormatUint(id, 10)

	s.mu.Lock()
	defer s.mu.Unlock()

	s.store.Assets[name] = b

	return name, nil
}

type byteFile struct {
	name   string
	reader *bytes.Reader
}

func (f *byteFile) Stat() (fs.FileInfo, error) {
	return (*fileInfo)(f), nil
}

func (f *byteFile) Read(b []byte) (int, error) {
	return f.reader.Read(b)
}

func (f *byteFile) Close() error {
	return nil
}

type fileInfo byteFile

func (n *fileInfo) Name() string       { return n.name }
func (n *fileInfo) Size() int64        { return n.reader.Size() }
func (n *fileInfo) Mode() fs.FileMode  { return 0444 } // -r--r--r--
func (n *fileInfo) ModTime() time.Time { return time.Now() }
func (n *fileInfo) IsDir() bool        { return false }
func (n *fileInfo) Sys() interface{}   { return nil }
