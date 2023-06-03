package collection

import "time"

type Collection struct {
	Id              string    `json:"id"` // Also the Name
	CreatedAt       time.Time `json:"createdAt"`
	SearchClusterId string    `json:"searchClusterId"`
	CommitWithin    int       `json:"commitWithin"`
	SolrParams      SolrParam `json:"solrParams"`
	Type            string    `json:"type"`
	Metadata        any       `json:"metadata"`
	Updates         []Update  `json:"updates"`
}

type SolrParam struct {
	Name              string `json:"name"`
	NumShards         int    `json:"numShards"`
	ReplicationFactor int    `json:"replicationFactor"`
}

type Update struct {
	UserId    string    `json:"userId"`
	Timestamp time.Time `json:"timestamp"`
}

type CreateCollection struct {
	SolrParams      SolrParam `json:"solrParams"`
	Id              string    `json:"id"`
	SearchClusterId string    `json:"searchClusterId"` // Optional. E.g. default
}
