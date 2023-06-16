package filesysytem

import (
	"fmt"
	"os"
	"strings"
)

func (f *fileSystemRepository) loadIndex() error {
	f.fm.Lock()
	defer f.fm.Unlock()

	if _, err := os.Stat(f.filePath); os.IsNotExist(err) {
		f.index = make(map[string]struct{})
		return nil
	}

	data, err := os.ReadFile(f.filePath)
	if err != nil {
		return fmt.Errorf("failed to read file by path: %s", f.filePath)
	}

	rows := strings.Split(string(data), "\n")

	f.im.Lock()
	defer f.im.Unlock()

	f.index = make(map[string]struct{}, len(rows))

	for _, row := range rows {
		if row == "" {
			continue
		}

		f.index[row] = struct{}{}
	}

	return nil
}
