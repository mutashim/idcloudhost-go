package idcloudhost

import "encoding/json"

type ParameterItem struct {
	Parameter   string          `json:"parameter"`
	Description string          `json:"description"`
	Mandatory   bool            `json:"mandatory"`
	Type        string          `json:"type"`
	Constraint  string          `json:"constraint"`
	IgnoreFor   string          `json:"ignore_for"`
	Min         int64           `json:"min"`
	Max         int64           `json:"max"`
	LimitedBy   string          `json:"limited_by"`
	Limits      json.RawMessage `json:"limits"`
	Values      []string        `json:"values"`
	Expression  string          `json:"expression"`
}

type ParameterLimitsOSVersion struct {
	OSName string   `json:"os_name"`
	Values []string `json:"values"`
}
