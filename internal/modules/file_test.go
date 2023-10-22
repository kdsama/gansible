package modules

import (
	"fmt"
	"testing"
)

func TestNewMode(t *testing.T) {

	want := "chmod 0644 /path/to/somewhere"
	got := modifyFileMode("0644", "/path/to/somewhere")
	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}

}

func TestNewFilePermissions(t *testing.T) {
	dummyPath := "/some/dummy/path"
	testMap := map[string]struct {
		input map[string]interface{}
		want  string
	}{
		"set only group": {
			input: map[string]interface{}{
				"path": dummyPath,

				"group": "kd",
			},
			want: fmt.Sprintf("chown :%s %s", "kd", dummyPath),
		},
		"set only owner": {
			input: map[string]interface{}{
				"path":  dummyPath,
				"owner": "kd",
			},
			want: fmt.Sprintf("chown %s %s", "kd", dummyPath),
		},
		"set owner and group": {
			input: map[string]interface{}{
				"path":  dummyPath,
				"owner": "kd",
				"group": "kd1",
			},
			want: fmt.Sprintf("chown %s:%s %s", "kd", "kd1", dummyPath),
		},
		"set mod ": {
			input: map[string]interface{}{
				"path": dummyPath,
				"mode": "0644",
			},
			want: fmt.Sprintf("chmod %s %s", "0644", dummyPath),
		},
		"set mod  and user ": {
			input: map[string]interface{}{
				"path": dummyPath,

				"mode":  "0644",
				"owner": "kd",
			},
			want: fmt.Sprintf("chown %s %s && chmod %s %s", "kd", dummyPath, "0644", dummyPath),
		},
		"set mod and group ": {
			input: map[string]interface{}{
				"path": dummyPath,

				"mode":  "0644",
				"group": "kd",
			},
			want: fmt.Sprintf("chown :%s %s && chmod %s %s", "kd", dummyPath, "0644", dummyPath),
		},
		"set mod group and user": {
			input: map[string]interface{}{
				"path": dummyPath,

				"mode":  "0644",
				"owner": "kd",
				"group": "kd1",
			},
			want: fmt.Sprintf("chown %s:%s %s && chmod %s %s", "kd", "kd1", dummyPath, "0644", dummyPath),
		},
	}

	for name, obj := range testMap {
		t.Run(name, func(t *testing.T) {
			got, err := NewFilePermissions(obj.input)
			if err != nil {
				t.Error("Expected nil but got error ", err)
			}
			if got != obj.want {
				t.Errorf("want %s but got %s", obj.want, got)
			}
		})
	}
}
