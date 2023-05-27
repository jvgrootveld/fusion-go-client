package datamodel

type DataModel struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// The name of the index pipeline used for the data model.
	IndexPipeline string `json:"indexPipeline"`
	// The name of the query pipeline used for the data model.
	QueryPipeline string  `json:"queryPipeline,omitempty"`
	Fields        []Field `json:"fields,omitempty"`
}

type Field struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Required    bool           `json:"required,omitempty"`
	Mappings    []FieldMapping `json:"mappings"`
}

type FieldMapping struct {
	SolrField   string `json:"solrField"`
	QueryField  bool   `json:"queryField,omitempty"`
	PhraseMatch bool   `json:"phraseMatch,omitempty"`
	// The amount of boost to give to the query.
	// If this is a query object, apply this boost to matches
	BoostValue float32 `json:"boostValue,omitempty"`
	// The amount of boost to give to the query, if it matches as an exact phrase.
	// If this is a phrase match object, apply this boost to matches
	PhraseBoost float32 `json:"phraseBoost,omitempty"`
}
