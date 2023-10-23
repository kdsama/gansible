package internal

import (
	"errors"
	"fmt"
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
	dummyPath, dummyLine := "/dummy/Path", "Hi this is dummyLine"
	testMap := map[string]struct {
		input map[string]interface{}
		want  []*Task
	}{
		"lineinfile + user": {
			input: map[string]interface{}{
				"lineinfile": map[string]interface{}{
					"path": dummyPath,
					"line": dummyLine,
				},
				"user": map[string]interface{}{
					"name":        "kd",
					"state":       "present",
					"create_home": true,
				},
			},
			want: []*Task{
				{[]string{fmt.Sprintf("echo \"%s\" >> %s", dummyLine, dummyPath)}},
				{[]string{fmt.Sprintf("useradd %s", "kd"), fmt.Sprintf("passwd -u %s", "kd"), fmt.Sprintf("mkdir /home/%s", "kd"), fmt.Sprintf("chown %s:%s /home/%s", "kd", "kd", "kd")}},
			},
		},
	}
	for name, obj := range testMap {
		t.Run(name, func(t *testing.T) {
			got, err := parseTask(obj.input)
			if err != nil {
				t.Errorf("Expected nil but got error %s", err.Error())
			}
			if len(got) != len(obj.want) {
				t.Errorf("wanted %v but got %v", obj.want, got)
			}
			for i := range got {
				for j := range got[i].cmds {
					if got[i].cmds[j] != obj.want[i].cmds[j] {
						t.Errorf("wanted %v but got %v", obj.want[i].cmds[j], got[i].cmds[j])
					}
				}
			}

		})

	}
}
