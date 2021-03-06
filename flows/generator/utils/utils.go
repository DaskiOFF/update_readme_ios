package utils

import (
	"path/filepath"

	"github.com/daskioff/jessica/utils/files"
)

// SearchTemplates возвращает список папок, находящихся в папке `templatesRoot` содержащих файл описания шаблона `templateDescriptionFileName`
func SearchTemplates(templatesRoot string, templateDescriptionFileName string) []string {
	folders := files.Folders(templatesRoot)

	templatesFolders := make([]string, 0)
	for _, folder := range folders {
		path := filepath.Join(templatesRoot, folder, templateDescriptionFileName)
		if files.IsFileExist(path) {
			templatesFolders = append(templatesFolders, folder)
		}
	}

	return templatesFolders
}
