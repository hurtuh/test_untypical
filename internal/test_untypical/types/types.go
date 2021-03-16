package types

type KeyValue struct {
	Key   interface{} `json:"key,omitempty"`
	Value interface{} `json:"value"`
}

type ListValues struct {
	Values []interface{} `json:"values"`
}
