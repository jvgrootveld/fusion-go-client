package fusion

import (
	"net/http"

	"github.com/jvgrootveld/fusion-go-client/fusion/datamodel"

	"github.com/jvgrootveld/fusion-go-client/fusion/index"
	"github.com/jvgrootveld/fusion-go-client/fusion/query"

	"github.com/jvgrootveld/fusion-go-client/fusion/connection"

	"github.com/jvgrootveld/fusion-go-client/fusion/auth"
)

// Config of the client endpoint
type Config struct {
	// Host of the Fusion instance; this is a mandatory field. (e.g. acme-dev.b.lucidworks.cloud)
	Host string
	// Scheme of the Fusion instance. (default: https)
	Scheme string

	// Application name required for application specific endpoints
	Application string

	// ConnectionClient that will be used to execute http requests to the Fusion instance.
	// If omitted a default will be used.
	ConnectionClient *http.Client

	// AuthConfig for authentication. Either this option or ConnectionClient can be used.
	AuthConfig auth.Config

	// Headers added for every request
	Headers map[string]string
}

// Client implementing the Fusion ProfileAPI
type Client struct {
	connection    *connection.Connection
	indexPipeline *index.PipelineAPI
	indexProfile  *index.ProfileAPI
	queryProfile  *query.ProfileAPI
	queryPipeline *query.PipelineAPI
	dataModel     *datamodel.API
}

// NewClient from Config
// Every function represents one ProfileAPI group of Fusion and provides a set of functions and builders to interact with them.
func NewClient(config Config) (*Client, error) {
	schema := config.Scheme
	if schema == "" {
		schema = "https"
	}

	con := connection.NewConnection(config.Scheme, config.Host, config.ConnectionClient, config.AuthConfig, config.Headers)

	// TODO ping Fusion

	client := &Client{
		connection:    con,
		indexPipeline: index.NewPipelineApi(con, config.Application),
		indexProfile:  index.NewProfileApi(con, config.Application),
		queryProfile:  query.NewProfileApi(con, config.Application),
		queryPipeline: query.NewPipelineApi(con, config.Application),
		dataModel:     datamodel.NewDataModelApi(con, config.Application),
	}

	return client, nil
}

// IndexPipeline ProfileAPI group
func (c *Client) IndexPipeline() *index.PipelineAPI {
	return c.indexPipeline
}

// IndexProfile ProfileAPI group
func (c *Client) IndexProfile() *index.ProfileAPI {
	return c.indexProfile
}

// QueryPipeline PipelineAPI group
func (c *Client) QueryPipeline() *query.PipelineAPI {
	return c.queryPipeline
}

// QueryProfile ProfileAPI group
func (c *Client) QueryProfile() *query.ProfileAPI {
	return c.queryProfile
}

// DataModel DataModelApi group
func (c *Client) DataModel() *datamodel.API {
	return c.dataModel
}
