package functional

import (
	"bufio"
	"errors"
	"io"
	"os"
)

func (suite *Suite) checkRowExistInFileDB(row string) (bool, error) {
	file, err := os.Open(testFilePath)
	if err != nil {
		return false, err
	}

	r := bufio.NewReader(file)

	for {
		line, _, err := r.ReadLine() //nolint: govet // shadow is unusable here
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return false, err
		}

		if string(line) == row {
			return true, nil
		}
	}

	return false, nil
}
