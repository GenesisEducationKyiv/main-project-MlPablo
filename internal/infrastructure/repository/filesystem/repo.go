package filesystem

import (
	"sync"
)

// I have decided to use a memory index for get operations for.
// This was done, because simple get operation is too heavy
// where we need to read the whole file and then iterate over every email (O(n)).
// On the big amount of data this can lead to performance issues.
// So better if we will index the file on the startup of the program
// and then we will add new items in the file and index.
// Get All operation made with file, just because I want to show read operation with filesystem)
// Don't forget about locks.
type Repository struct {
	filePath string
	// file mutex
	fm sync.RWMutex
}

func NewFileSystemRepository(filePath string) (*Repository, error) {
	f := &Repository{
		filePath: filePath,
		fm:       sync.RWMutex{},
	}

	return f, nil
}
