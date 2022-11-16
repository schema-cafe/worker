package golang

import (
	"os"
	"strings"

	"github.com/schema-cafe/go-types/filesystem"
)

func ReadDir(path string) (filesystem.Folder, error) {
	folder := filesystem.Folder{}
	entries, err := os.ReadDir(path)
	if err != nil {
		return folder, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			if !strings.HasPrefix(entry.Name(), ".") {
				folder.Contents = append(folder.Contents, filesystem.FolderItem{
					IsFolder: true,
					Name:     entry.Name(),
				})
			}
		} else {
			if strings.HasSuffix(entry.Name(), ".go") {
				folder.Contents = append(folder.Contents, filesystem.FolderItem{
					IsFolder: false,
					Name:     strings.TrimSuffix(entry.Name(), ".go"),
				})
			}
		}
	}
	return folder, nil
}
