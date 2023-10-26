package modules

import "fmt"

func modifyFileOwnership(owner, group string, file string) string {
	var str string

	if owner != "" {
		str += owner
	}
	if group != "" {
		str += fmt.Sprintf(":%s", group)
	}

	return fmt.Sprintf("chown %s %s", str, file)
}
func modifyFileMode(mod string, file string) string {
	return fmt.Sprintf("chmod %s %s", mod, file)
}

func NewFilePermissions(fileMap map[string]interface{}) ([]string, error) {
	var (
		path   string
		owner  string
		group  string
		result []string
	)
	if _, ok := fileMap[pathParam]; !ok {
		return result, fmt.Errorf("path %w", ErrNotFound)
	}
	path = fileMap[pathParam].(string)
	if _, ok := fileMap[ownerParam]; ok {
		owner = fileMap[ownerParam].(string)
	}
	if _, ok := fileMap[groupParam]; ok {
		group = fileMap[groupParam].(string)
	}
	if owner != "" || group != "" {
		result = append(result, modifyFileOwnership(owner, group, path))
	}

	if _, ok := fileMap[modeParam]; ok {
		result = append(result, modifyFileMode(fileMap[modeParam].(string), path))
	}
	return result, nil
}
