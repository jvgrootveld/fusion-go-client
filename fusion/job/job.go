package job

import (
	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/generic"
	"strings"
	"time"
)

const ApiName = "jobs"

// API Contains all the builders required to access the Fusion Job
// Also see Fusion docs: https://doc.lucidworks.com/fusion/5.8/327/jobs-api
type API struct {
	connection  *connection.Connection
	application string
}

// NewApi api group with connection
func NewApi(con *connection.Connection, application string) *API {
	return &API{connection: con, application: application}
}

// Deleter new builder to delete Jobs
func (api *API) Deleter() *generic.Deleter {
	return generic.NewDeleter(api.connection, ApiName).
		ForApplication(api.application)
}

// Getter new builder to retrieve one or all Jobs
func (api *API) Getter() *generic.Getter[Job] {
	return generic.NewGetter[Job](api.connection, ApiName).
		ForApplication(api.application)
}

// Creator builder to create new Jobs
func (api *API) Creator() *generic.Creator[Job] {
	return generic.NewCreator[Job](api.connection, ApiName).
		ForApplication(api.application).
		WithModelValidator(createJobValidator)
}

// Spark new Api (sub) group for Spark Jobs
func (api *API) Spark() *SparkAPI {
	return NewSparkAPI(api.connection, api.application)
}

type Job struct {
	Id            string         `json:"resource"`
	Enabled       bool           `json:"enabled"`
	Status        string         `json:"status"`
	Extra         map[string]any `json:"extra"`
	LastStartTime time.Time      `json:"lastStartTime"`
	LastEndTime   time.Time      `json:"lastEndTime"`
	NextStartTime time.Time      `json:"nextStartTime"`
}

type Type string

const (
	TypeDatasource Type = "datasource"
	TypeSpark      Type = "spark"
	TypeTask       Type = "task"
	TypeUnknown    Type = "unknown"
)

var allTypes = []Type{TypeDatasource, TypeSpark, TypeTask}

func (j Job) Type() Type {
	for _, t := range allTypes {
		if strings.HasPrefix(j.Id, string(t)+":") {
			return t
		}
	}

	return TypeUnknown
}
