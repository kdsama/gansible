package modules

import (
	"errors"
	"testing"
)

func TestNewLineInFileErrors(t *testing.T) {

	type testcase struct {
		input map[string]interface{}
		want  error
	}

	testMap := map[string]testcase{
		"Path not present": testcase{
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
		"Append not present": {
			input: map[string]interface{}{
				"path": "/some/line",
				"line": "some line",
			},
			want: ErrNotFound,
		},
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
