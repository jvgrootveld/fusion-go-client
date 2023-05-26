package stage

// SolrQuery sends the search request to Solr
type SolrQuery struct {
	BaseStage
}

func NewSolrQuery(label string) SolrQuery {
	return SolrQuery{
		BaseStage: NewBaseStage("solr-query", label),
	}
}
