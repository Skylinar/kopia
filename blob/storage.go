package blob

import (
	"bytes"
	"io"
	"time"
)

// PutOptions modify the behavior of Storage.PutBlock().
type PutOptions int

// Possible values of PutOptions
const (
	PutOptionsDefault   PutOptions = 0
	PutOptionsOverwrite PutOptions = 1
)

// Storage encapsulates API for connecting to blob storage
type Storage interface {
	io.Closer

	PutBlock(id string, data ReaderWithLength, options PutOptions) error
	DeleteBlock(id string) error
	Flush() error
	BlockExists(id string) (bool, error)
	GetBlock(id string) ([]byte, error)
	ListBlocks(prefix string) chan (BlockMetadata)
	Configuration() StorageConfiguration
}

// ReaderWithLength supports reading from a block and returns its length.
type ReaderWithLength interface {
	io.ReadCloser
	Len() int
}

type bytesReaderWithLength struct {
	*bytes.Buffer
}

// NewReader wraps the provided buffer and returns a ReaderWithLength.
func NewReader(b *bytes.Buffer) ReaderWithLength {
	return &bytesReaderWithLength{b}
}

func (bbr *bytesReaderWithLength) Close() error {
	return nil
}

// BlockMetadata represents metadata about a single block in a blob.
// If Error field is set, no other field values should be assumed to be correct.
type BlockMetadata struct {
	BlockID   string
	Length    uint64
	TimeStamp time.Time
	Error     error
}
