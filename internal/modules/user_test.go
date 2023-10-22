package modules

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewFileOwner(t *testing.T) {

	type testcase struct {
		owner string
		group string
		file  string
		want  string
	}
	dummyFile := "/somefile/"
	testMap := map[string]testcase{
		"only owner": {
			owner: "kshitij",
			group: "",
			file:  dummyFile,
			want:  fmt.Sprintf("chown kshitij %s", dummyFile),
		},
		"only group": {
			owner: "",
			group: "kdGroup",
			file:  dummyFile,
			want:  fmt.Sprintf("chown :kdGroup %s", dummyFile),
		},
		"group and owner": {
			owner: "kshitij",
			group: "kdGroup",
			file:  dummyFile,
			want:  fmt.Sprintf("chown kshitij:kdGroup %s", dummyFile),
		},
	}
	for name, obj := range testMap {
		t.Run(name, func(t *testing.T) {
			got := NewFileOwner(obj.owner, obj.group, obj.file)
			if obj.want != got {
				t.Errorf("Wanted %v but got %v", obj.want, got)
			}
		})
	}
}

func TestNewUserErrors(t *testing.T) {
	testMap := map[string]struct {
		input map[string]interface{}
		want  error
	}{
		"name not present": {
			input: map[string]interface{}{
				"notpresent": "name",
			},
			want: ErrNotFound,
		},
	}

	for name, obj := range testMap {
		t.Run(name, func(t *testing.T) {
			_, got := NewUser(obj.input)
			if !errors.Is(got, obj.want) {
				t.Errorf("wanted %v but got %v", obj.want, got)
			}
		})

	}
}
