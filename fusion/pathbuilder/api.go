package pathbuilder

import (
	"strings"
)

const (
	apiPathPrefix  = "/api"
	appsPathPrefix = "/api/apps"
)

// ApiPath builds the api url path with Components
// When application is set it's interpreted as an application api e.g. `/api/apps/acme/index-pipelines`
// When application is empty it's interpreted as a generic api e.g. `/api/collections'
func ApiPath(comp Components) string {
	var parts []string

	if comp.Application == "" {
		parts = []string{
			apiPathPrefix,
		}
	} else {
		parts = []string{
			appsPathPrefix,
			comp.Application, // E.g. acme
		}
	}

	parts = append(parts, comp.ApiName) // E.g. index-pipelines

	// Adds ObjectId only if set
	if comp.ObjectId != "" {
		parts = append(parts, comp.ObjectId)
	}

	return strings.Join(parts, "/")
}
