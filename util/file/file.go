package file

import (
	"os"
	"path/filepath"
)

func IsExist(paths ...string) bool {
	if _, err := os.Stat(filepath.Join(paths...)); err == nil {
		return true
	} else {
		return false
	}
}
