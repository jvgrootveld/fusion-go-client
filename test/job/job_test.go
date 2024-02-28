package job

import (
	"github.com/jvgrootveld/fusion-go-client/fusion/job"
	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
	"testing"
)

// buildJobUrl is a convenience function to create a Job url
func buildJobUrl(id string) string {
	return testsuit.CreateApplicationUrl(job.ApiName, id)
}

// buildSparkJobUrl is a convenience function to create a Job url
func buildSparkJobUrl(id string) string {
	return testsuit.CreateApplicationUrl(job.SparkApiName, id)
}

func TestTypes(t *testing.T) {

	tests := map[string]struct {
		id       string
		expected job.Type
	}{
		"Type datasource": {
			id:       "datasource:an-id",
			expected: job.TypeDatasource,
		},
		"Type spark": {
			id:       "spark:an-id",
			expected: job.TypeSpark,
		},
		"Type task": {
			id:       "task:an-id",
			expected: job.TypeTask,
		},
		"Type unknown for unmatched value": {
			id:       "something:an-id",
			expected: job.TypeUnknown,
		},
		"Type unknown for empty value": {
			id:       "",
			expected: job.TypeUnknown,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			model := job.Job{Id: test.id}
			actual := model.Type()

			if actual != test.expected {
				t.Errorf("types don't match: expected %q, got %q", test.expected, actual)
			}
		})
	}
}
