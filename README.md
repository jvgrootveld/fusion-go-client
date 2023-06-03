<a href="https://www.buymeacoffee.com/jvgrootveld" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" style="height: 50px !important;"></a>

# Fusion Go Client

Fusion Client library for Go.

> :warning: This library is far from complete. I will add when I need stuff. Any tips and feedback are always welcome!

## Example usage

An example creating an Index Pipeline with 3 stages.

```golang
package example

import (
	"context"
	"fusion-go-client/fusion"
	"fusion-go-client/fusion/auth"
	"fusion-go-client/fusion/indexpipeline/stage"
)

func createPipelineTest(username, password string) error {
	client, err := fusion.NewClient(fusion.Config{
		Host:        "acme-dev.b.lucidworks.cloud",
		Scheme:      "https",
		Application: "application_name",
		AuthConfig:  auth.NewBasicAuthConfig(username, password),
	})

	if err != nil {
		return err
	}

	return client.IndexPipeline().
		Creator().
		WithID("pipeline-id").
		WithStages(
			stage.NewFieldMapping("Field Mapping"),
			stage.NewSolrDynamicFieldNameMapping("Solr Dynamic Field Mapping"),
			stage.NewSolrIndex("Solr Index"),
		).
		Do(context.Background())
}
```
