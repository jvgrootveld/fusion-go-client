package collection

import (
	"github.com/jvgrootveld/fusion-go-client/fusion/collection"
	"github.com/jvgrootveld/fusion-go-client/test/testsuit"
)

// buildUrl is a convenience function to create a Collection url
func buildUrl(id string) string {
	return testsuit.CreateApplicationUrl(collection.ApiName, id)
}
