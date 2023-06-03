package pathbuilder

// Components of the url Path
type Components struct {
	Application string // Only for app paths
	ApiName     string // e.g. index-pipelines
	ObjectId    string // Optional
}
