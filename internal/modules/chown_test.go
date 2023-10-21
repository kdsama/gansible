package modules

import (
	"fmt"
	"testing"
)

func TestNewOwner(t *testing.T) {

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
			got := NewOwner(obj.owner, obj.group, obj.file)
			if obj.want != got {
				t.Errorf("Wanted %v but got %v", obj.want, got)
			}
		})
	}
}
