package stage

// TextTagger queries a Solr text tagger request handler to perform spell correction, phrase boosting, and synonym expansion.
type TextTagger struct {
	BaseStage
}

func NewTextTagger(label string) TextTagger {
	return TextTagger{
		BaseStage: NewBaseStage("text-tagger", label),
	}
}
