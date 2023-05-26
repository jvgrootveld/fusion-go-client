package index

import (
	"context"
	"net/http"

	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/except"
	"github.com/jvgrootveld/fusion-go-client/fusion/pathbuilder"
)

// ProfileGetter builder to get an Index profile
type ProfileGetter struct {
	connection  *connection.Connection
	application string
	id          string
}

// NewProfileGetter with connection
func NewProfileGetter(connection *connection.Connection) *ProfileGetter {
	return &ProfileGetter{
		connection: connection,
	}
}

// ForApplication specifies the application the profile is configured.
// Note: This overrides the default in the Connection
func (getter *ProfileGetter) ForApplication(application string) *ProfileGetter {
	getter.application = application
	return getter
}

// WithID specifies a specific Index Profile to get.
// When empty, all are returned.
func (getter *ProfileGetter) WithID(id string) *ProfileGetter {
	getter.id = id
	return getter
}

// Do get one or all Index profile's
func (getter *ProfileGetter) Do(ctx context.Context) ([]Profile, error) {
	path := pathbuilder.AppsPath(pathbuilder.Components{
		Application: getter.application,
		ApiName:     ProfileApiName,
		ObjectId:    getter.id,
	})
	responseData, err := getter.connection.RunREST(ctx, path, http.MethodGet, nil)
	err = except.CheckResponseDataErrorAndStatusCode(responseData, err, 200)
	if err != nil {
		return nil, err
	}

	// If id is set, unmarshal to a single instance
	if getter.id != "" {
		var resultProfile Profile
		parseErr := responseData.DecodeBodyIntoTarget(&resultProfile)
		return []Profile{resultProfile}, parseErr
	}

	var resultProfiles []Profile
	parseErr := responseData.DecodeBodyIntoTarget(&resultProfiles)
	return resultProfiles, parseErr
}
