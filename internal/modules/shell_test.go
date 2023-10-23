package modules

import (
	"testing"
)

func TestNewShell(t *testing.T) {

	var testMap = map[string]struct {
		input map[string]interface{}
		want  []string
	}{
		"without change directory": {
			input: map[string]interface{}{
				"cmd": "echo HelloWorld",
			},
			want: []string{"echo HelloWorld"},
		},
		"with change directory": {
			input: map[string]interface{}{
				"cmd":   "echo HelloWorld",
				"chdir": "/new/folder",
			},
			want: []string{"cd /new/folder", "echo HelloWorld"},
		},
		"with different executable": {
			input: map[string]interface{}{
				"cmd":        "echo HelloWorld",
				"executable": "/bin/bash",
			},
			want: []string{"chsh -s /bin/bash", "echo HelloWorld"},
		},
		"with different executable and change directory": {
			input: map[string]interface{}{
				"cmd":        "echo HelloWorld",
				"chdir":      "/some/folder",
				"executable": "/bin/bash",
			},
			want: []string{"cd /some/folder", "chsh -s /bin/bash", "echo HelloWorld"},
		},
	}

	for name, obj := range testMap {
		t.Run(name, func(t *testing.T) {
			got, err := NewShell(obj.input)
			if err != nil {
				t.Errorf("Expected no error but got %v", err)
			}
			if len(obj.want) != len(got) {
				t.Errorf("Incorrect number of commands. Got %v but got %v", got, obj.want)

			}
			for i := range got {
				if obj.want[i] != got[i] {
					t.Errorf("Wanted %v but got %v", obj.want[i], got[i])
				}
			}
		})
	}

}
