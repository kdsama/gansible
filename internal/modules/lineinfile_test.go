package modules

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewLineInFile(t *testing.T) {
	t.Parallel()
	type testcase struct {
		input map[string]interface{}
		want  string
	}
	dummyPath, dummyLine := "/dummy/Path", "Hi this is dummyLine"
	testMap := map[string]testcase{
		"return true if the line is present": {
			input: map[string]interface{}{
				"path":  dummyPath,
				"line":  dummyLine,
				"state": "present",
			},
			want: fmt.Sprintf("if grep -q %s %s; then echo 'true'; else echo 'false'; fi", dummyLine, dummyPath),
		},
		"return true if the line is absent": {
			input: map[string]interface{}{
				"path":  dummyPath,
				"line":  dummyLine,
				"state": "absent",
			},
			want: fmt.Sprintf("if grep -q %s %s; then echo 'false'; else echo 'true'; fi", dummyLine, dummyPath),
		},
		"Command line to add the line at the end of the file": {
			input: map[string]interface{}{
				"path": dummyPath,
				"line": dummyLine,
			},
			want: fmt.Sprintf("echo \"%s\" >> %s", dummyLine, dummyPath),
		},
	}
	for testname, testObj := range testMap {
		t.Run(testname, func(t *testing.T) {
			got, err := NewLineInFile(testObj.input)
			if err != nil {
				t.Error("Wanted an error but got nil")
			}
			if testObj.want != got {
				t.Errorf("Wanted %v but got %v", testObj.want, got)
			}
		})
	}
}

func TestNewLineInFileErrors(t *testing.T) {
	t.Parallel()
	type testcase struct {
		input map[string]interface{}
		want  error
	}

	testMap := map[string]testcase{
		"Path not present": {
			input: map[string]interface{}{
				"line":   "/some/line",
				"append": "yes",
			},
			want: ErrNotFound,
		},
		"Line not present": {
			input: map[string]interface{}{
				"path":   "/some/line",
				"append": "yes",
			},
			want: ErrNotFound,
		},
		"Invalid State": {
			input: map[string]interface{}{
				"path":  "/some/line",
				"line":  "append this content",
				"state": "garbagestate",
			},
			want: ErrInvalidInput,
		},
		// "Append not present": {
		// 	input: map[string]interface{}{
		// 		"path": "/some/line",
		// 		"line": "some line",
		// 	},
		// 	want: ErrNotFound,
		// },
	}

	for testname, testObj := range testMap {
		t.Run(testname, func(t *testing.T) {
			_, got := NewLineInFile(testObj.input)
			if got == nil {
				t.Error("Wanted an error but got nil")
			}
			if errors.Is(testObj.want, got) {
				t.Errorf("Wanted %v but got %v", testObj.want, got)
			}
		})
	}

}
