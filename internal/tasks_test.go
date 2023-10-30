package internal

import (
	"errors"
	"fmt"
	"testing"

	"github.com/kdsama/gansible/internal/modules"
)

func TestParseTaskErrors(t *testing.T) {

	testMap := map[string]struct {
		input []map[string]interface{}
		want  error
	}{

		"Path Not Found for filepermissions": {
			input: []map[string]interface{}{
				{
					"file": map[string]interface{}{
						"group": "kd",
					},
				},
			},
			want: modules.ErrNotFound,
		},
		"line Not Found for lineinfile": {
			input: []map[string]interface{}{
				{
					"lineinfile": map[string]interface{}{
						"group": "kd",
					},
				},
			},
			want: modules.ErrNotFound,
		},
		"shell error (cmd not found)with rest good cases": {
			input: []map[string]interface{}{
				{
					"lineinfile": map[string]interface{}{
						"group": "kd",
						"line":  "some line",
						"path":  "/some/good/path",
					},
				},
				{
					"file": map[string]interface{}{
						"owner": "kd",
						"path":  "/this/is/the/file",
						"group": "kd1,kd2",
					},
				},
				{
					"shell": map[string]interface{}{
						"owner": "kd",
						"path":  "/this/is/the/file",
						"group": "kd1,kd2",
					},
				},
			},
			want: modules.ErrNotFound,
		},
	}
	for name, obj := range testMap {
		t.Run(name, func(t *testing.T) {
			for _, inputs := range obj.input {
				_, got := parseTask(inputs)
				if got == nil {
					continue
				}
				if !errors.Is(got, obj.want) {
					t.Errorf("want %v but got %v", obj.want, got)
				}
			}

		})

	}
}

func TestParseTask(t *testing.T) {
	dummyPath, dummyLine := "/dummy/Path", "Hi this is dummyLine"
	testMap := map[string]struct {
		input []map[string]interface{}
		want  []*Task
	}{
		"lineinfile + user": {
			input: []map[string]interface{}{
				{
					"lineinfile": map[string]interface{}{
						"path": dummyPath,
						"line": dummyLine,
					},
				},
				{
					"user": map[string]interface{}{
						"name":        "kd",
						"state":       "present",
						"create_home": true,
					},
				},
			},
			want: []*Task{
				{cmds: []string{fmt.Sprintf("echo \"%s\" >> %s", dummyLine, dummyPath)}, os: "any"},
				{cmds: []string{fmt.Sprintf("useradd %s", "kd"), fmt.Sprintf("passwd -u %s", "kd"), fmt.Sprintf("mkdir /home/%s", "kd"), fmt.Sprintf("chown %s:%s /home/%s", "kd", "kd", "kd")}, os: "any"},
			},
		},
	}
	for name, obj := range testMap {
		t.Run(name, func(t *testing.T) {
			for j := 0; j < len(obj.input); j++ {
				got, err := parseTask(obj.input[j])
				if err != nil {
					t.Errorf("Expected nil but got error %s", err.Error())
				}
				for i := range got.cmds {
					if got.cmds[i] != obj.want[j].cmds[i] {
						t.Errorf("wanted %v but got %v", obj.want[j].cmds[i], got.cmds[i])
					}
				}
			}

		})

	}
}
