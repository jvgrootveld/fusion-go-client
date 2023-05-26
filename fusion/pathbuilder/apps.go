package pathbuilder

import (
	"strings"
)

const appsPathPrefix = "/api/apps"

// AppsPath creates an apps url path with Components
// Example "/api/apps/acme/index-pipelines/test-pipeline"
func AppsPath(comp Components) string {
	parts := []string{
		appsPathPrefix,
		comp.Application, // E.g. acme
		comp.ApiName,     // E.g. index-pipelines
	}

	// Adds ObjectId only if set
	if comp.ObjectId != "" {
		parts = append(parts, comp.ObjectId)
	}

	return strings.Join(parts, "/")
}
