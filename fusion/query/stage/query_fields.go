package stage

// QueryFields configure query parameters for Solr search
type QueryFields struct {
	BaseStage
}

func NewQueryFields(label string) QueryFields {
	return QueryFields{
		BaseStage: NewBaseStage("search-fields", label),
	}
}
