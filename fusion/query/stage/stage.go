package stage

// Stage type for Index and Query pipeline stages
type Stage interface{}

// BaseStage contains properties that every stage has
type BaseStage struct {
	Skip      bool   `json:"skip"`
	Type      string `json:"type"`
	Label     string `json:"label,omitempty"`
	Condition string `json:"condition,omitempty"`
}

// NewBaseStage creates a new BaseStage with widely used fields typeName and label
func NewBaseStage(typeName, label string) BaseStage {
	return BaseStage{
		Type:  typeName,
		Label: label,
	}
}
