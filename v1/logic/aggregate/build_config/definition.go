package buildconfig

import (
	api_model "cmx/v1/logic/model/api"
	enum_model "cmx/v1/logic/model/enum"
	message_model "cmx/v1/logic/model/message"
	statement_model "cmx/v1/logic/model/statement"
	"cmx/v1/logic/util"
	"cmx/v1/pkg/logger"
	"cmx/v1/pkg/tree"
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Definition struct {
	messageTree *tree.Tree

	enumsTree   *tree.Tree
	enumsRoutes []string

	tableTree     *tree.Tree
	statementTree *tree.Tree
	apiTree       *tree.Tree
	// baseModel     *Tree
	projectPath string
}

func NewDefinition(projectPath string) *Definition {
	df := &Definition{
		messageTree:   tree.NewTree(),
		enumsTree:     tree.NewTree(),
		enumsRoutes:   []string{},
		tableTree:     tree.NewTree(),
		statementTree: tree.NewTree(),
		apiTree:       tree.NewTree(),
		projectPath:   projectPath,
	}
	return df
}

//! --- table ---

// load table definition besides cache
func (df *Definition) AddTable(path string) {
	tdf := message_model.Table{}
	router := df.createRouterFromPath(path)
	df.tableTree.Add(router, util.MustSuccess(tdf.ParseFile(path)))
}

func (df *Definition) GetTable(route string) (msg message_model.Table, is bool) {
	if value, found := df.tableTree.Search(route); found {
		return *value.Item.(*message_model.Table), found
	} else {
		return message_model.Table{}, false
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

func (df *Definition) GetTableField(referenceString string) (msg message_model.Field, is bool) {
	reference := NewReferenceInformation(referenceString)
	if reference == nil {
		logger.Warnf(context.Background(), "reference is nil: %s", referenceString)
		return message_model.Field{}, false
	}
	if value, found := df.tableTree.Search(fmt.Sprintf("/%s", reference.Route)); found {
		di := *value.Item.(*message_model.Table)
		for _, v := range di.Definition[reference.Object] {
			if v.ColumnName == reference.Field {
				return v, true
			}
		}
	}
	return message_model.Field{}, false
}

//! --- create statement ---

func (df *Definition) AddStatement(path string) {
	sdf := statement_model.Statement{}
	router := df.createRouterFromPath(path)
	df.statementTree.Add(router, util.MustSuccess(sdf.ParseFile(path)))
}

func (df *Definition) GetStatement(route string) (msg statement_model.Statement, is bool) {
	if value, found := df.statementTree.Search(route); found {
		return *value.Item.(*statement_model.Statement), found
	} else {
		return statement_model.Statement{}, false
	}
}

func (df *Definition) GetStatementField(referenceString string) (msg statement_model.CreateStatement, is bool) {
	reference := NewReferenceInformation(referenceString)
	if reference == nil {
		logger.Warnf(context.Background(), "reference is nil: %s", referenceString)
		return statement_model.CreateStatement{}, false
	}

	if value, found := df.statementTree.Search(fmt.Sprintf("/%s", reference.Route)); found {
		di := *value.Item.(*statement_model.Statement)
		return di.Definition[reference.Field], true
	}

	return statement_model.CreateStatement{}, false
}

//! --- message ---

// load message definition besides cache
func (df *Definition) AddMessage(path string) {
	mdf := message_model.Message{}
	router := df.createRouterFromPath(path)
	df.messageTree.Add(router, util.MustSuccess(mdf.ParseFile(path)))
}

func (df *Definition) GetMessages(route string) (msg message_model.Message, is bool) {
	if value, found := df.messageTree.Search(route); found {
		return *value.Item.(*message_model.Message), found
	} else {
		return message_model.Message{}, false
	}
}

// GetMessagesBySpecify get message by specify
func (df *Definition) GetMessagesBySpecify(referenceString string) (msg []message_model.MessageField, is bool) {
	reference := NewReferenceInformation(referenceString)
	value, found := df.GetMessages("/" + reference.Route)
	if !found {
		return msg, false
	}
	if _, ok := value.Definition[reference.Field]; ok {
		return value.Definition[reference.Field], true
	}
	return msg, false
}

func (df *Definition) GetMessageField(referenceString string) (msg []message_model.MessageField, is bool) {
	reference := NewReferenceInformation(referenceString)
	if reference == nil {
		logger.Warnf(context.Background(), "reference is nil: %s", referenceString)
		return []message_model.MessageField{}, false
	}

	if value, found := df.messageTree.Search(fmt.Sprintf("/%s", reference.Route)); found {
		di := *value.Item.(*message_model.Message)
		return di.Definition[reference.Field], true
	}

	return []message_model.MessageField{}, false
}

//! --- enum ---

// load enum definition besides cache
func (df *Definition) AddEnums(path string) {
	edf := enum_model.EnumsGroup{}
	router := df.createRouterFromPath(path)
	df.enumsRoutes = lo.Uniq(append(df.enumsRoutes, router))
	df.enumsTree.Add(router, util.MustSuccess(edf.ParseFile(path)))
}

func (df *Definition) GetEnum(route string) (msg enum_model.EnumsGroup, is bool) {
	if value, found := df.enumsTree.Search(route); found {
		return *value.Item.(*enum_model.EnumsGroup), found
	} else {
		return enum_model.EnumsGroup{}, false
	}
}

func (df *Definition) GetEnumField(referenceString string) (msg []enum_model.Enum, is bool) {
	reference := NewReferenceInformation(referenceString)
	if reference == nil {
		logger.Warnf(context.Background(), "reference is nil: %s", referenceString)
		return []enum_model.Enum{}, false
	}

	if value, found := df.enumsTree.Search(fmt.Sprintf("/%s", reference.Route)); found {
		di := *value.Item.(*enum_model.EnumsGroup)
		return di.Definition[reference.Field], true
	}

	return []enum_model.Enum{}, false
}

func (df *Definition) SelectEnumField(foo message_model.FieldOneOf) []enum_model.Enum {
	var enums []enum_model.Enum
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
func (df *Definition) GetEnumComment(foo message_model.FieldOneOf) string {
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
	adf := api_model.ApiDefinition{}
	router := df.createRouterFromPath(path)
	df.enumsTree.Add(router, util.MustSuccess(adf.ParseFile(path)))
}

func (df *Definition) GetApi(route string) (dapis api_model.ApiDefinition, is bool) {
	if value, found := df.enumsTree.Search(route); found {
		return *value.Item.(*api_model.ApiDefinition), found
	} else {
		return api_model.ApiDefinition{}, false
	}
}

func (df *Definition) GetGroup(groupName string) (apis []api_model.Api, is bool) {
	adf := api_model.ApiDefinition{}
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
