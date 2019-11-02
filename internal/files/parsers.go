package files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ildarusmanov/go-up/internal/constants"
)

// GetFoldersPathsList returns list of files for given directory
func GetFoldersPathsList(dir string) ([]string, error) {
	var dirList []string

	err := filepath.Walk(dir, func(curPath string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("folder traversing error: %s", err)
		}

		if info.IsDir() {
			relDirPath := strings.ReplaceAll(curPath, dir, "")
			dirList = append(dirList, relDirPath)
		}

		return nil
	})

	return dirList, err
}

// GetTemplatesPathsList returns *.tmpl  files in provided dir */
func GetTemplatesPathsList(dir string) ([]string, error) {
	var filesList []string
	err := filepath.Walk(dir, func(curPath string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("folder traversing error: %s", err)
		}

		if !info.IsDir() && filepath.Ext(curPath) == constants.TemplateExt {
			filesList = append(filesList, curPath)
		}

		return nil
	})

	return filesList, err
}
