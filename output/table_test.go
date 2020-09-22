package output

import (
	"bytes"
	"strings"
	"testing"
)

func TestTable(t *testing.T) {
	tests := []struct {
		name     string
		input    []CheckResult
		expected []string
	}{
		{
			name: "No warnings or errors",
			input: []CheckResult{
				{
					FileName: "examples/kubernetes/service.yaml",
				},
			},
			expected: []string{},
		},
		{
			name: "A warning and a failure",
			input: []CheckResult{
				{
					FileName: "examples/kubernetes/service.yaml",
					Warnings: []Result{{Message: "first warning"}},
					Failures: []Result{{Message: "first failure"}},
				},
			},
			expected: []string{
				`+---------+----------------------------------+---------------+`,
				`| RESULT  |               FILE               |    MESSAGE    |`,
				`+---------+----------------------------------+---------------+`,
				`| warning | examples/kubernetes/service.yaml | first warning |`,
				`| failure | examples/kubernetes/service.yaml | first failure |`,
				`+---------+----------------------------------+---------------+`,
				``,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expected := strings.Join(tt.expected, "\n")

			buf := new(bytes.Buffer)
			if err := NewTable(buf).Output(tt.input); err != nil {
				t.Fatal("output table:", err)
			}
			actual := buf.String()

			if expected != actual {
				t.Errorf("Unexpected output. expected %v actual %v", expected, actual)
			}
		})
	}
}
