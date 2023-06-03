package query

import (
	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/generic"
)

const ProfileApiName = "query-profiles"

// ProfileAPI Contains all the builders required to access the Fusion Query Profile
// Also see Fusion docs: https://doc.lucidworks.com/fusion/5.8/363/query-profiles-api
type ProfileAPI struct {
	connection  *connection.Connection
	application string
}

// NewProfileApi api group with connection
func NewProfileApi(con *connection.Connection, application string) *ProfileAPI {
	return &ProfileAPI{connection: con, application: application}
}

// Deleter new builder to delete Query Profiles
func (api *ProfileAPI) Deleter() *generic.Deleter {
	return generic.NewDeleter(api.connection, ProfileApiName).
		ForApplication(api.application)
}

// Getter new builder to retrieve one or all Query Profiles
func (api *ProfileAPI) Getter() *generic.Getter[Profile] {
	return generic.NewGetter[Profile](api.connection, ProfileApiName).
		ForApplication(api.application)
}

// Creator builder to create new Query Profiles
func (api *ProfileAPI) Creator() *generic.Creator[Profile] {
	return generic.NewCreator[Profile](api.connection, ProfileApiName).
		ForApplication(api.application).
		WithModelValidator(createProfileValidator)
}

type Profile struct {
	Id                   string            `json:"id"`
	QueryPipeline        string            `json:"queryPipeline"`
	Collection           string            `json:"collection"`
	SearchHandler        string            `json:"searchHandler,omitempty"`
	Params               []Param           `json:"params,omitempty"`
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty"`
}
