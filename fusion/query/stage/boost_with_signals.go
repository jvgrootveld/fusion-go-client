package stage

// BoostWithSignals uses recommended items for search time boosting
type BoostWithSignals struct {
	BaseStage
}

func NewBoostWithSignals(label string) BoostWithSignals {
	return BoostWithSignals{
		BaseStage: NewBaseStage("recommendation", label),
	}
}
