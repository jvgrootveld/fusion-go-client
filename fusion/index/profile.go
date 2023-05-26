package index

import (
	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/generic"
)

const ProfileApiName = "index-profiles"

// ProfileAPI Contains all the builders required to access the Fusion Index Profile
type ProfileAPI struct {
	connection  *connection.Connection
	application string
}

// NewProfile api group with connection
func NewProfile(con *connection.Connection, application string) *ProfileAPI {
	return &ProfileAPI{connection: con, application: application}
}

// Deleter new builder to delete Index Profile's
func (api *ProfileAPI) Deleter() *generic.Deleter {
	return generic.NewDeleter(api.connection, ProfileApiName).
		ForApplication(api.application)
}

// Getter new builder to retrieve one or all Index Profile's
func (api *ProfileAPI) Getter() *ProfileGetter {
	return NewProfileGetter(api.connection).
		ForApplication(api.application)
}

// Creator builder to create new Index Profile's
func (api *ProfileAPI) Creator() *ProfileCreator {
	return NewProfileCreator(api.connection).
		ForApplication(api.application)
}

type Profile struct {
	Id                   string            `json:"id"`
	IndexPipeline        string            `json:"indexPipeline"`
	Collection           string            `json:"collection"`
	Parser               string            `json:"parser,omitempty"`
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty"`
}
