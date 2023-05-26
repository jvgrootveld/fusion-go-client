package stage

// Facets configure range or field facets.
// Facets indicate categories and aggregations of results according to values in the configured fields
type Facets struct {
	BaseStage
}

func NewFacets(label string) Facets {
	return Facets{
		BaseStage: NewBaseStage("facet", label),
	}
}
