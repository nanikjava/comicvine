package models

type CreateRequest struct {
	Document  interface{} `json:"doc"`
	Operation string      `json:"op"`
	IsBatch   bool        `json:"isBatch"`
}

// SQLMetaData stores sql query information
type SQLMetaData struct {
	Col       string        `json:"col" structs:"col"`
	SQL       string        `json:"sql" structs:"sql"`
	DbAlias   string        `json:"db" structs:"db"`
	Args      []interface{} `json:"args" structs:"args"`
	QueryTime string        `json:"queryTime" structs:"queryTime"`
}

// PostProcess filters the schema
type PostProcess struct {
	PostProcessAction []PostProcessAction
}

// PostProcessAction is struct of Action Field Value
type PostProcessAction struct {
	Action string
	Field  string
	Value  interface{}
}

// JoinOption describes the way a join needs to be performed
type JoinOption struct {
	// Op can be either All or One
	// This field decides the way the result of join is returned
	// If op is all, the result is returned as an array
	// If op is one, the result is returned as an object
	Op    string                 `json:"Op" mapstructure:"Op"`
	Type  string                 `json:"type" mapstructure:"type"`
	Table string                 `json:"table" mapstructure:"table"`
	As    string                 `json:"as" mapstructure:"as"`
	On    map[string]interface{} `json:"on" mapstructure:"on"`
	Join  []*JoinOption          `json:"join" mapstructure:"join"`
}

type ReadOptions struct {
	// Debug field is used internally to show
	// _query meta data in the graphql
	Debug      bool             `json:"debug"`
	Select     map[string]int32 `json:"select"`
	Sort       []string         `json:"sort"`
	Skip       *int64           `json:"skip"`
	Limit      *int64           `json:"limit"`
	Distinct   *string          `json:"distinct"`
	Join       []*JoinOption    `json:"join"`
	ReturnType string           `json:"returnType"`
	HasOptions bool             `json:"hasOptions"` // used internally
}

// ReadCacheOptions describes the cache options in requests
type ReadCacheOptions struct {
	TTL               int64 `json:"ttl" yaml:"ttl" mapstructure:"ttl"` // here ttl is represented in seconds
	InstantInvalidate bool  `json:"instantInvalidate" yaml:"instantInvalidate" mapstructure:"instantInvalidate"`
}

type ReadRequest struct {
	GroupBy     []interface{}            `json:"group"`
	Aggregate   map[string][]string      `json:"aggregate"`
	Find        map[string]interface{}   `json:"find"`
	Operation   string                   `json:"op"`
	Options     *ReadOptions             `json:"options"`
	IsBatch     bool                     `json:"isBatch"`
	Extras      map[string]interface{}   `json:"extras"`
	PostProcess map[string]*PostProcess  `json:"postProcess"`
	MatchWhere  []map[string]interface{} `json:"matchWhere"`
	Cache       *ReadCacheOptions        `json:"cache"`
}
