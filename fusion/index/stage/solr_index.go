package stage

// SolrIndex sends documents to Solr
type SolrIndex struct {
	BaseStage
}

func NewSolrIndex(label string) SolrIndex {
	return SolrIndex{
		BaseStage: NewBaseStage("solr-index", label),
	}
}
