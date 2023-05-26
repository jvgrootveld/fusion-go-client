package stage

// ApplyRules look up and apply rules to the query
type ApplyRules struct {
	BaseStage
}

func NewApplyRules(label string) ApplyRules {
	return ApplyRules{
		BaseStage: NewBaseStage("query-rules", label),
	}
}
