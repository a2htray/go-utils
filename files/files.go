package files

import (
	"bufio"
	"os"
)

// ReadAllLines is to read all lines from a given file path
func ReadAllLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, bufio.MaxScanTokenSize)
	// https://stackoverflow.com/questions/21124327/how-to-read-a-text-file-line-by-line-in-go-when-some-lines-are-long-enough-to-ca
	scanner.Buffer(buf, 1024 * 1024)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Remove files from the disk
func Remove(paths... string) error {
	for _, path := range paths {
		err := os.Remove(path)
		if err != nil {
			return err
		}
	}
	return nil
}