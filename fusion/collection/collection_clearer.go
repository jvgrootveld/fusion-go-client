package collection

import (
	"context"
	"fmt"
	"github.com/jvgrootveld/fusion-go-client/fusion/fault"
	"net/http"
	"strings"

	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/except"
)

const QueryAll = "*:*"

// Clearer builder to clear a collection
type Clearer struct {
	connection *connection.Connection
	id         string
	query      string
}

// NewClearer with Connection and apiName of the configured object to create like index-pipeline or query-profile
func NewClearer(connection *connection.Connection) *Clearer {
	return &Clearer{
		connection: connection,
		query:      QueryAll,
	}
}

// WithId the collection to clear
func (clearer *Clearer) WithID(id string) *Clearer {
	clearer.id = id
	return clearer
}

// WithQuery specifies which items to delete from the Collection. Default: QueryAll
func (clearer *Clearer) WithQuery(query string) *Clearer {
	clearer.query = query
	return clearer
}

// Do delete the specified items in the Collection
func (clearer *Clearer) Do(ctx context.Context) error {
	err := clearer.Validate()
	if err != nil {
		return err
	}

	body := SolrUpdate{SolrUpdateDelete{Query: clearer.query}}

	path := strings.Join([]string{
		"/api/solr",
		clearer.id,
		"update",
	}, "/")

	path += "?commit=true"

	responseData, err := clearer.connection.RunREST(ctx, path, http.MethodPost, body)
	return except.CheckResponseDataErrorAndStatusCode(responseData, err, 200)
}

func (clearer *Clearer) Validate() error {
	typeName := fmt.Sprint(ApiName, "Clearer")

	if clearer.id == "" {
		return fault.NewRequiredError(typeName, "id")
	}

	return nil
}

type SolrUpdate struct {
	Delete SolrUpdateDelete `json:"delete"`
}

type SolrUpdateDelete struct {
	Query string `json:"query"`
}
