package modules

import (
	"errors"
	"fmt"
	"testing"
)

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

func TestNewUser(t *testing.T) {

	testMap := map[string]struct {
		input map[string]interface{}
		want  []string
	}{
		"remove user": {
			input: map[string]interface{}{
				"name":   "kd",
				"remove": true,
			},
			want: []string{fmt.Sprintf("userdel -r %s", "kd")},
		},
		"lock user": {
			input: map[string]interface{}{
				"name":  "kd",
				"state": "absent",
			},
			want: []string{fmt.Sprintf("passwd -l %s", "kd")},
		},
		"add user, without home directory": {
			input: map[string]interface{}{
				"name":  "kd",
				"state": "present",
			},
			want: []string{fmt.Sprintf("useradd %s", "kd"), fmt.Sprintf("passwd -u %s", "kd")},
		},
		"add user and assign it groups": {
			input: map[string]interface{}{
				"name":   "kd",
				"state":  "present",
				"groups": "group1,group2",
			},
			want: []string{fmt.Sprintf("useradd %s", "kd"), fmt.Sprintf("passwd -u %s", "kd"), "groupadd group1", "groupadd group2", fmt.Sprintf("usermod -aG %s %s", "group1,group2", "kd")},
		},
		"add user and home directory": {
			input: map[string]interface{}{
				"name":        "kd",
				"state":       "present",
				"create_home": true,
			},
			want: []string{fmt.Sprintf("useradd %s", "kd"), fmt.Sprintf("passwd -u %s", "kd"), fmt.Sprintf("mkdir /home/%s", "kd"), fmt.Sprintf("chown %s:%s /home/%s", "kd", "kd", "kd")},
		},
	}
	for name, obj := range testMap {
		t.Run(name, func(t *testing.T) {
			got, err := NewUser(obj.input)
			if err != nil {
				t.Errorf("expected nil but got error %v", err)
			}
			if len(obj.want) != len(got) {
				t.Errorf("want %v but got %v", obj.want, got)
			}
			for i := range got {
				if got[i] != obj.want[i] {
					t.Errorf("wanted %v but got %v", obj.want[i], got[i])
				}
			}
		})

	}
}

func TestModifyFileOwnership(t *testing.T) {

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
			got := modifyFileOwnership(obj.owner, obj.group, obj.file)
			if obj.want != got {
				t.Errorf("Wanted %v but got %v", obj.want, got)
			}
		})
	}
}
