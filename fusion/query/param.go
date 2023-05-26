package query

type PolicyType string

const (
	Append  PolicyType = "append"
	Default PolicyType = "default"
	Remove  PolicyType = "remove"
	Replace PolicyType = "replace"
)

type Param struct {
	Key    string     `json:"key"`
	Value  string     `json:"value"`
	Policy PolicyType `json:"policy"`
}
