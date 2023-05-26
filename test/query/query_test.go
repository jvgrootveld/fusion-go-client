package query

import (
	"github.com/jvgrootveld/fusion-go-client/fusion/query"
	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
)

// buildPipelineUrl is a convenience function to create a pipeline url
func buildPipelineUrl(id string) string {
	return testsuit.CreateApplicationUrl(query.PipelineApiName, id)
}

// buildProfileUrl is a convenience function to create a profile url
func buildProfileUrl(id string) string {
	return testsuit.CreateApplicationUrl(query.ProfileApiName, id)
}
