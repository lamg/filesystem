package filesystem

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// FileSystem abstracts the OS file system
type FileSystem interface {
	Open(string) (File, error)
	Stat(string) (os.FileInfo, error)
	Create(string) (File, error)
	Rename(string, string) error
}

// File abstracts a file system file
type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
	Stat() (os.FileInfo, error)
}

// OSFS implements FileSystem using the OS file system
type OSFS struct {
}

// Open opens a file
func (fs *OSFS) Open(name string) (f File, e error) {
	f, e = os.Open(name)
	return
}

// Rename renames a file
func (fs *OSFS) Rename(old, new string) (e error) {
	e = os.Rename(old, new)
	return
}

// Stat stats a file
func (fs *OSFS) Stat(name string) (f os.FileInfo, e error) {
	f, e = os.Stat(name)
	return
}

// Create creates a file
func (fs *OSFS) Create(name string) (f File, e error) {
	f, e = os.Create(name)
	return
}

// BufferFS implements FileSystem using in memory buffers
type BufferFS struct {
	bfs map[string]*BFile
}

// NewBufferFS creates a new BufferFS
func NewBufferFS() (b *BufferFS) {
	b = &BufferFS{make(map[string]*BFile)}
	return
}

// Open creates a new file in memory
func (b *BufferFS) Open(name string) (f File, e error) {
	var ok bool
	f, ok = b.bfs[name]
	if !ok {
		e = fmt.Errorf("Not found file %s", name)
	}
	return
}

// GetBuffer gets the underlying
func (b *BufferFS) GetBuffer(n string) (bf *bytes.Buffer,
	ok bool) {
	var f *BFile
	f, ok = b.bfs[n]
	if ok {
		bf = f.Buffer
	}
	return
}

// Create creates a new file in memory
func (b *BufferFS) Create(name string) (f File, e error) {
	b.bfs[name] = NewBFile()
	f = b.bfs[name]
	return
}

// Rename renames a file
func (b *BufferFS) Rename(old, new string) (e error) {
	f, ok := b.bfs[old]
	if ok {
		delete(b.bfs, old)
		b.bfs[new] = f
	} else {
		e = fmt.Errorf("File %s doesn't exists", old)
	}
	return
}

// Stat stats an in memory file
func (b *BufferFS) Stat(name string) (f os.FileInfo, e error) {
	return
}

// BFile is a file stored in memory as a *bytes.Buffer
type BFile struct {
	*bytes.Buffer
}

// NewBFile creates a new BFile
func NewBFile() (b *BFile) {
	b = &BFile{bytes.NewBufferString("")}
	return
}

// Close closes the BFile
func (b *BFile) Close() (e error) {
	// TODO? b.bf.Reset()
	e = fmt.Errorf("Not implemented")
	return
}

// Stat stats the BFile
func (b *BFile) Stat() (f os.FileInfo, e error) {
	// TODO
	e = fmt.Errorf("Not implemented")
	return
}

// ReadAt implementation of io.ReaderAt
func (b *BFile) ReadAt(p []byte, off int64) (n int, e error) {
	e = fmt.Errorf("Not implemented")
	return
}

// Seek implementation of io.Seeker
func (b *BFile) Seek(offset int64,
	whence int) (n int64, e error) {
	e = fmt.Errorf("Not implemented")
	return
}
