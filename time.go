package main

import (
	"fmt"
	"path/filepath"
)

func produceDatePath(year int, month int, day int) string {
	return filepath.Join(
		BasePath,
		fmt.Sprintf("%04d", year),
		fmt.Sprintf("%02d", month),
		fmt.Sprintf("%02d", day),
	)
}
