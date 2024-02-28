package job

import (
	"github.com/jvgrootveld/fusion-go-client/fusion/connection"
	"github.com/jvgrootveld/fusion-go-client/fusion/generic"
	"time"
)

const SparkApiName = "spark/configurations"

// SparkAPI Contains all the builders required to access the Fusion Spark Job
// Also see Fusion docs: https://doc.lucidworks.com/fusion/5.8/312/spark-jobs-api
type SparkAPI struct {
	connection  *connection.Connection
	application string
}

// NewSparkAPI api group with connection
func NewSparkAPI(con *connection.Connection, application string) *SparkAPI {
	return &SparkAPI{connection: con, application: application}
}

// Deleter new builder to delete Spark Jobs
func (api *SparkAPI) Deleter() *generic.Deleter {
	return generic.NewDeleter(api.connection, SparkApiName).
		ForApplication(api.application)
}

// Getter new builder to retrieve one or all Spark Jobs
func (api *SparkAPI) Getter() *SparkGetter {
	return NewSparkGetter(api.connection, SparkApiName).
		ForApplication(api.application)
}

// Creator builder to create new Spark Jobs
func (api *SparkAPI) Creator() *generic.Creator[Spark] {
	return generic.NewCreator[Spark](api.connection, SparkApiName).
		ForApplication(api.application).
		WithModelValidator(createSparkJobValidator)
}

const (
	SparkTypeParallelBulkLoader         = "parallel-bulk-loader"
	SparkTypeSynonymDetection           = "synonymDetection"
	SparkTypeHeadTailAnalysis           = "headTailAnalysis"
	SparkTypeSqlTemplate                = "sql_template"
	SparkTypeSip                        = "sip"
	SparkTypeArgoItemRecommenderUser    = "argo-item-recommender-user"
	SparkTypeArgoItemRecommenderContent = "argo-item-recommender-content"
	SparkTypeSimilarQueries             = "similar_queries"
	SparkTypeAlsRecommender             = "als_recommender"
	SparkTypeTokenPhraseSpellCorrection = "tokenPhraseSpellCorrection"
)

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SparkUpdate struct {
	UserId    string    `json:"userId"`
	Timestamp time.Time `json:"timestamp"`
}

type Spark interface {
	spark()
}

type SparkGeneric struct {
	Id      string        `json:"id"`
	Type    string        `json:"type"`
	Updates []SparkUpdate `json:"updates"`
}

func (SparkGeneric) spark() { return }

type SparkParallelBulkLoader struct {
	SparkGeneric
	Config                       []KeyValue `json:"sparkConfig"`
	Format                       string     `json:"format"`
	ReadOptions                  []KeyValue `json:"readOptions"`
	OutputCollection             string     `json:"outputCollection"`
	OutputIndexPipeline          string     `json:"outputIndexPipeline"`
	ClearDatasource              bool       `json:"clearDatasource"`
	DefineFieldsUsingInputSchema bool       `json:"defineFieldsUsingInputSchema"`
	AtomicUpdates                bool       `json:"atomicUpdates"`
	TransformScala               string     `json:"transformScala"`
	CacheAfterRead               bool       `json:"cacheAfterRead"`
	ContinueAfterFailure         bool       `json:"continueAfterFailure"`
}

func (SparkParallelBulkLoader) spark() { return }
