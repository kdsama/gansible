package internal

import (
	"errors"
	"testing"

	"github.com/kdsama/gansible/internal/modules"
)

func TestParseTaskErrors(t *testing.T) {

	testMap := map[string]struct {
		input map[string]interface{}
		want  error
	}{

		"Path Not Found for filepermissions": {
			input: map[string]interface{}{
				"file": map[string]interface{}{
					"group": "kd",
				},
			},
			want: modules.ErrNotFound,
		},
		"line Not Found for lineinfile": {
			input: map[string]interface{}{
				"lineinfile": map[string]interface{}{
					"group": "kd",
				},
			},
			want: modules.ErrNotFound,
		},
		"shell error (cmd not found)with rest good cases": {
			input: map[string]interface{}{
				"lineinfile": map[string]interface{}{
					"group": "kd",
					"line":  "some line",
					"path":  "/some/good/path",
				},
				"file": map[string]interface{}{
					"owner": "kd",
					"path":  "/this/is/the/file",
					"group": "kd1,kd2",
				},
				"shell": map[string]interface{}{
					"owner": "kd",
					"path":  "/this/is/the/file",
					"group": "kd1,kd2",
				},
			},
			want: modules.ErrNotFound,
		},
	}
	for name, obj := range testMap {
		t.Run(name, func(t *testing.T) {
			_, got := parseTask(obj.input)
			if !errors.Is(got, obj.want) {
				t.Errorf("want %v but got %v", obj.want, got)
			}
		})

	}
}

func TestParseTask(t *testing.T) {

}
