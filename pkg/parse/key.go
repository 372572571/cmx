package parse


type ModelKey struct {
	KeyType KeyType `json:"key_type" yaml:"key_type"`
	Key     string        `json:"key" yaml:"key"`
	// is complex key
	IsComplex bool     `json:"is_complex" yaml:"is_complex"`
	Field     []string `json:"field" yaml:"field"`
}

func NewModelKey() *ModelKey {
	return &ModelKey{
		Field: []string{},
	}
}
