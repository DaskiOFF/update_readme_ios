package projectstruct

import (
	"path/filepath"
)

// projectPaths возвращает список путей, которые необходимо создать.
// Например: ["app/classes/models", "app/classes/flows"]
func (flow *ProjectStructFlow) projectPaths() []string {
	projectStructure := flow.projectConfig.GetCustomProjectStructDescription()
	return flow.projectStructToPaths(projectStructure)
}

// projectStructToPaths генерирует по описанию структуры проекта из файла конфигурации списка путей для создаваемый папок
func (flow *ProjectStructFlow) projectStructToPaths(in interface{}) []string {
	switch v := in.(type) {

	case string:
		return []string{v}

	case map[interface{}]interface{}:
		response := make([]string, 0)
		for s, b := range v {
			prefix, _ := s.(string)
			for _, path := range flow.projectStructToPaths(b) {
				response = append(response, filepath.Join(prefix, path))
			}
		}
		return response

	case []string:
		return v

	case []interface{}:
		response := make([]string, 0)
		for _, b := range v {
			for _, path := range flow.projectStructToPaths(b) {
				response = append(response, path)
			}
		}
		return response

	default:
		return make([]string, 0)
	}
}
