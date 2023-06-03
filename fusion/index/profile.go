package index

import (
	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/generic"
)

const ProfileApiName = "index-profiles"

// ProfileAPI Contains all the builders required to access the Fusion Index Profile
// Also see Fusion docs: https://doc.lucidworks.com/fusion/5.8/znnx16/index-profiles-api
type ProfileAPI struct {
	connection  *connection.Connection
	application string
}

// NewProfileApi api group with connection
func NewProfileApi(con *connection.Connection, application string) *ProfileAPI {
	return &ProfileAPI{connection: con, application: application}
}

// Deleter new builder to delete Index Profile's
func (api *ProfileAPI) Deleter() *generic.Deleter {
	return generic.NewDeleter(api.connection, ProfileApiName).
		ForApplication(api.application)
}

// Getter new builder to retrieve one or all Index Profiles
func (api *ProfileAPI) Getter() *generic.Getter[Profile] {
	return generic.NewGetter[Profile](api.connection, ProfileApiName).
		ForApplication(api.application)
}

// Creator builder to create new Index Profiles
func (api *ProfileAPI) Creator() *generic.Creator[Profile] {
	return generic.NewCreator[Profile](api.connection, ProfileApiName).
		ForApplication(api.application).
		WithModelValidator(createProfileValidator)
}

type Profile struct {
	Id                   string            `json:"id"`
	IndexPipeline        string            `json:"indexPipeline"`
	Collection           string            `json:"collection"`
	Parser               string            `json:"parser,omitempty"`
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty"`
}
