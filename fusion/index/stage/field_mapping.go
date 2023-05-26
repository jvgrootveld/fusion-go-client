package stage

// FieldMapping Keep, delete, add, set, copy, or move fields on a document.
type FieldMapping struct {
	BaseStage
}

func NewFieldMapping(label string) FieldMapping {
	return FieldMapping{
		BaseStage: NewBaseStage("field-mapping", label),
	}
}
