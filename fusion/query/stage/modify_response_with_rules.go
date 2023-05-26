package stage

// ModifyResponseWithRules Modify the response from Solr using matching rules from the Apply Rules stage
type ModifyResponseWithRules struct {
	BaseStage
}

func NewModifyResponseWithRules(label string) ModifyResponseWithRules {
	return ModifyResponseWithRules{
		BaseStage: NewBaseStage("query-rules-augment-response", label),
	}
}
