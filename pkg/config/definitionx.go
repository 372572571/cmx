package config

import (
	"cmx/pkg/config/definition"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Definition struct {
	messageTree *Tree

	enumsTree   *Tree
	enumsRoutes []string

	tableTree     *Tree
	statementTree *Tree
	apiTree       *Tree
	// baseModel     *Tree
	projectPath string
}

func NewDefinition(projectPath string) *Definition {
	df := &Definition{
		messageTree:   NewTree(),
		enumsTree:     NewTree(),
		enumsRoutes:   []string{},
		tableTree:     NewTree(),
		statementTree: NewTree(),
		apiTree:       NewTree(),
		projectPath:   projectPath,
	}
	return df
}

//! --- table ---

// load table definition besides cache
func (df *Definition) AddTable(path string) {
	tdf := definition.Tabledefinition{}
	router := df.createRouterFromPath(path)
	df.tableTree.Add(router, tdf.ParseFile(path))
}

func (df *Definition) GetTable(route string) (msg definition.Tabledefinition, is bool) {
	if value, found := df.tableTree.Search(route); found {
		return *value.Item.(*definition.Tabledefinition), found
	} else {
		return definition.Tabledefinition{}, false
	}
}

// func (df *Definition) GetTableObject(referenceString string) (msg []definition.Field, is bool) {
// 	reference := NewReferenceInformation(referenceString)
// 	if reference == nil {
// 		panic(fmt.Errorf("reference is nil"))
// 	}
// 	if value, found := df.tableTree.Search(fmt.Sprintf("/%s", reference.Route)); found {
// 		di := *value.Item.(*definition.Tabledefinition)
// 		fmt.Println(di.Definition)
// 		if _, ok := di.Definition[reference.Object]; ok {
// 			return di.Definition[reference.Object], true
// 		}
// 	}

// 	return []definition.Field{}, false
// }

func (df *Definition) GetTableField(referenceString string) (msg definition.Field, is bool) {
	reference := NewReferenceInformation(referenceString)
	if reference == nil {
		fmt.Println("reference is nil")
		return definition.Field{}, false
	}
	if value, found := df.tableTree.Search(fmt.Sprintf("/%s", reference.Route)); found {
		di := *value.Item.(*definition.Tabledefinition)
		for _, v := range di.Definition[reference.Object] {
			if v.ColumnName == reference.Field {
				return v, true
			}
		}
	}
	return definition.Field{}, false
}

//! --- create statement ---

func (df *Definition) AddStatement(path string) {
	sdf := definition.Statementdefinition{}
	router := df.createRouterFromPath(path)
	df.statementTree.Add(router, sdf.ParseFile(path))
}

func (df *Definition) GetStatement(route string) (msg definition.Statementdefinition, is bool) {
	if value, found := df.statementTree.Search(route); found {
		return *value.Item.(*definition.Statementdefinition), found
	} else {
		return definition.Statementdefinition{}, false
	}
}

func (df *Definition) GetStatementField(referenceString string) (msg definition.CreateStatement, is bool) {
	reference := NewReferenceInformation(referenceString)
	if reference == nil {
		fmt.Println("reference is nil")
		return definition.CreateStatement{}, false
	}

	if value, found := df.statementTree.Search(fmt.Sprintf("/%s", reference.Route)); found {
		di := *value.Item.(*definition.Statementdefinition)
		return di.Definition[reference.Field], true
	}

	return definition.CreateStatement{}, false
}

//! --- message ---

// load message definition besides cache
func (df *Definition) AddMessage(path string) {
	mdf := definition.MessageDefinition{}
	router := df.createRouterFromPath(path)
	df.messageTree.Add(router, mdf.ParseFile(path))
}

func (df *Definition) GetMessages(route string) (msg definition.MessageDefinition, is bool) {
	if value, found := df.messageTree.Search(route); found {
		return *value.Item.(*definition.MessageDefinition), found
	} else {
		return definition.MessageDefinition{}, false
	}
}

// GetMessagesBySpecify get message by specify
func (df *Definition) GetMessagesBySpecify(referenceString string) (msg []definition.MessageField, is bool) {
	reference := NewReferenceInformation(referenceString)
	value, found := df.GetMessages("/" + reference.Route)
	if !found {
		return msg, false
	}
	if _, ok := value.Definition[reference.Field]; ok {
		return value.Definition[reference.Field], true
	}
	fmt.Println("not2")
	return msg, false
}

func (df *Definition) GetMessageField(referenceString string) (msg []definition.MessageField, is bool) {
	reference := NewReferenceInformation(referenceString)
	if reference == nil {
		fmt.Println("reference is nil")
		return []definition.MessageField{}, false
	}

	if value, found := df.messageTree.Search(fmt.Sprintf("/%s", reference.Route)); found {
		di := *value.Item.(*definition.MessageDefinition)
		return di.Definition[reference.Field], true
	}

	return []definition.MessageField{}, false
}

//! --- enum ---

// load enum definition besides cache
func (df *Definition) AddEnums(path string) {
	edf := definition.Enumsdefinition{}
	router := df.createRouterFromPath(path)
	df.enumsRoutes = lo.Uniq(append(df.enumsRoutes, router))
	df.enumsTree.Add(router, edf.ParseFile(path))
}

func (df *Definition) GetEnum(route string) (msg definition.Enumsdefinition, is bool) {
	if value, found := df.enumsTree.Search(route); found {
		return *value.Item.(*definition.Enumsdefinition), found
	} else {
		return definition.Enumsdefinition{}, false
	}
}

func (df *Definition) GetEnumField(referenceString string) (msg []definition.Enumx, is bool) {
	reference := NewReferenceInformation(referenceString)
	if reference == nil {
		fmt.Println("reference is nil")
		return []definition.Enumx{}, false
	}

	if value, found := df.enumsTree.Search(fmt.Sprintf("/%s", reference.Route)); found {
		di := *value.Item.(*definition.Enumsdefinition)
		return di.Definition[reference.Field], true
	}

	return []definition.Enumx{}, false
}

func (df *Definition) SelectEnumField(foo definition.FieldOneOf) []definition.Enumx {
	var enums []definition.Enumx
	if foo.Ref == "" {
		return enums
	}
	if value, found := df.GetEnumField(foo.Ref); found {
		if len(foo.Select) == 1 && foo.Select[0] == "*" {
			return value
		}
		for _, v := range value {
			for _, s := range foo.Select {
				if v.Key == s {
					enums = append(enums, v)
				}
			}
		}
		return enums
	}
	return enums
}

func (df *Definition) GetEnumRoutes() []string {
	return df.enumsRoutes
}

// get enum join comment
func (df *Definition) GetEnumComment(foo definition.FieldOneOf) string {
	list := df.SelectEnumField(foo)
	if len(list) == 0 {
		return ""
	}
	comment := strings.Builder{}
	comment.WriteString("有效类型: ")
	if foo.IsKey {
		for _, v := range list {
			desc := v.Zh
			if desc == "" {
				desc = v.Desc
			}
			comment.WriteString(fmt.Sprintf("%s: %s ", v.Key, desc))
		}
	} else {
		for _, v := range list {
			desc := v.Zh
			if desc == "" {
				desc = v.Desc
			}
			comment.WriteString(fmt.Sprintf("%s: %s ", v.Value, desc))
		}
	}
	return strconv.Quote(comment.String())
}

// ! --- api---
// load enum definition besides cache
func (df *Definition) AddApi(path string) {
	adf := definition.Apidefinition{}
	router := df.createRouterFromPath(path)
	df.enumsTree.Add(router, adf.ParseFile(path))
}

func (df *Definition) GetApi(route string) (dapis definition.Apidefinition, is bool) {
	if value, found := df.enumsTree.Search(route); found {
		return *value.Item.(*definition.Apidefinition), found
	} else {
		return definition.Apidefinition{}, false
	}
}

func (df *Definition) GetGroup(groupName string) (apis []definition.Api, is bool) {
	adf := definition.Apidefinition{}
	result := adf.GetGroup(groupName)
	if len(result) == 0 {
		return
	}
	return result, true
}

//! --- util ---

// removeProjectPathPrefix remove project path prefix
func (df *Definition) removeProjectPathPrefix(filePath string) string {
	relPath, err := filepath.Rel(df.projectPath, filePath)
	if err != nil {
		return filePath
	}
	return relPath
}

// removeExt remove file extension
func (df *Definition) removeExt(filePath string) string {
	ext := filepath.Ext(filePath)
	if ext != "" {
		filePath = filePath[:len(filePath)-len(ext)]
	}
	return filePath
}

// createRouterFromPath create router from file path
func (df *Definition) createRouterFromPath(path string) string {
	file := df.removeProjectPathPrefix(path)
	return fmt.Sprintf("/%s", df.removeExt(file))
}
