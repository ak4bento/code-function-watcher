package utils

import (
	"bufio"
	"os"
	"strings"
)

func LoadIgnoreList(path string) (map[string]struct{}, error) {
	ignoreSet := make(map[string]struct{})

	file, err := os.Open(path)
	if err != nil {
		// File tidak wajib ada
		return ignoreSet, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			ignoreSet[line] = struct{}{}
		}
	}

	return ignoreSet, scanner.Err()
}
