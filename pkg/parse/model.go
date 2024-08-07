package parse

import (
	"cmp"

	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type Model struct {
	Reference           string `json:"reference" yaml:"reference"`
	TableName           string `json:"table_name" yaml:"table_name"`
	UpTableName         string `json:"up_table_name" yaml:"up_table_name"`
	FirstLowerTableName string `json:"first_lower_table_name" yaml:"first_lower_table_name"`
	PackageName         string `json:"package_name" yaml:"package_name"`
	Comment             string `json:"comment" yaml:"comment"`
	// TODO 主键不允许复合主键
	Key    map[KeyType][]*ModelKey `json:"key" yaml:"key"`
	Fields map[string]*ModelField  `json:"fields" yaml:"fields"`
}

func NewModel() *Model {
	return &Model{
		Key: map[KeyType][]*ModelKey{
			KeyTypePrimary: []*ModelKey{},
			KeyTypeUnique:  []*ModelKey{},
			KeyTypeIndex:   []*ModelKey{},
		},
		Fields: make(map[string]*ModelField),
	}
}

func (m Model) AddKey(keyType KeyType, key *ModelKey) {
	m.Key[keyType] = append(m.Key[keyType], key)
}

func (m Model) AddField(fieldName string, field *ModelField) {
	m.Fields[fieldName] = field
}

func (m Model) SearchFieldKey(field string) map[KeyType][]*ModelKey {
	result := map[KeyType][]*ModelKey{}

	for t, key := range m.Key {
		for _, k := range key {
			if slices.Contains(k.Field, field) {
				result[t] = append(result[t], k)
			}
		}
	}

	return result
}

func (m Model) SortFields() []ModelField {
	ary := lo.MapToSlice(m.Fields, func(k string, v *ModelField) ModelField {
		return *v
	})
	slices.SortFunc(ary, func(i ModelField, j ModelField) int {
		return cmp.Compare(i.Idx, j.Idx)
	})
	return ary
}
