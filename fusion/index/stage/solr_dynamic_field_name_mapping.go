package stage

// SolrDynamicFieldNameMapping map pipeline document fields to a Solr dynamic fields.
type SolrDynamicFieldNameMapping struct {
	BaseStage
}

func NewSolrDynamicFieldNameMapping(label string) SolrDynamicFieldNameMapping {
	return SolrDynamicFieldNameMapping{
		BaseStage: NewBaseStage("solr-dynamic-field-name-mapping", label),
	}
}
