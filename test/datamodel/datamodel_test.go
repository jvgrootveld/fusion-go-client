package datamodel

import (
	"github.com/jvgrootveld/fusion-go-client/fusion/datamodel"
	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
)

// buildUrl is a convenience function to create a Data Model url
func buildUrl(id string) string {
	return testsuit.CreateApplicationUrl(datamodel.ApiName, id)
}
