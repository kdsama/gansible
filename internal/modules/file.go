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

func NewFilePermissions(fileMap map[string]interface{}) (string, error) {
	var (
		path  string
		owner string
		group string
		str   string
	)
	if _, ok := fileMap["path"]; !ok {
		return "", fmt.Errorf("path %w", ErrNotFound)
	}
	path = fileMap["path"].(string)
	if _, ok := fileMap["owner"]; ok {
		owner = fileMap["owner"].(string)
	}
	if _, ok := fileMap["group"]; ok {
		group = fileMap["group"].(string)
	}
	if owner != "" || group != "" {
		str = modifyFileOwnership(owner, group, path)
	}

	if _, ok := fileMap["mode"]; ok {
		if str != "" {
			str = fmt.Sprintf("%s && %s", str, modifyFileMode(fileMap["mode"].(string), path))
		} else {
			str = fmt.Sprint(modifyFileMode(fileMap["mode"].(string), path))
		}

	}
	return str, nil
}
