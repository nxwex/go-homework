package files

import (
	"os"
	"path/filepath"
)

type Files struct{}

func (f *Files) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (f *Files) IsJSON(path string) bool {
	return filepath.Ext(path) == ".json"
}
