package file

import(
	"os"
	"path/filepath"
)

func ReadAllFiles(dirname string)(map[string]os.FileInfo, error) {
	list := make(map[string]os.FileInfo, 0)

	err := filepath.Walk(dirname, func(path string, f os.FileInfo, err error) error {
		list[path] = f
		return nil	
	})
	if err != nil {
		return nil, err
	}
	return list, nil
}