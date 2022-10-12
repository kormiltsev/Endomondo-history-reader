package file_searcher

import (
	"fmt"
	"path/filepath"
	"strings"
)

// like : "./originaldata/*.csv"

func FileSearcher(keyword string, dir string) []string {
	filenames := make([]string, 0)

	m, err := filepath.Glob(dir)
	if err != nil {
		panic(err)
	}
	for _, filename := range m {
		if strings.Contains(filename, keyword) {
			filenames = append(filenames, filename)
		}
	}
	return filenames
}
