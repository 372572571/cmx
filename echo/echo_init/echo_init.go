package echo_init

import (
	"cmx/echo/echo_init/data_source"
	"cmx/pkg/config"
	"cmx/pkg/config/definition"
	"cmx/pkg/parse"
	"cmx/pkg/util"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/lo"
	"gopkg.in/yaml.v3"
)

// 初始化模型
// - 数据库配置
// - 获取数据库create statement
// - 判断是否存在已经生成的模型
//   - 存在则加载
//     - 根据数据库表结构变更,模型
//   - 不存在则创建
//     - 根据数据库表结构创建模型

func Generated(cfg config.Config, sourceType data_source.ESourceType) {
	// * 生成必要目录
	// cfg.ProjectPath  项目路径
	project := cfg.ProjectPath
	modelPath := filepath.Join(project, "model")

	// 根据配置生成api路径
	for k, _ := range cfg.Apis {
		apiPath := filepath.Join(project, k)
		os.MkdirAll(apiPath, os.ModePerm)

		apiConf := filepath.Join(apiPath, k+".conf")
		if !util.IsHaveFile(apiConf) {
			sb := strings.Builder{}
			sb.WriteString(getFileApiConf())
			os.WriteFile(apiConf, []byte(sb.String()), os.ModePerm)
		}
	}

	defaultPath := filepath.Join(project, "default")
	messagePath := filepath.Join(project, "message")
	localSqlPath := filepath.Join(project, "local")

	os.MkdirAll(modelPath, os.ModePerm)
	os.MkdirAll(defaultPath, os.ModePerm)
	os.MkdirAll(messagePath, os.ModePerm)
	os.MkdirAll(localSqlPath, os.ModePerm)

	linkModel := []string{}
	var sourceData = []*data_source.Create{}
	switch sourceType {
	case data_source.SourceTypeMysql:
		sourceData = data_source.NewMysqlData(cfg.DBConfig).Source()
	case data_source.SourceTypeLocal:
		sourceData = data_source.NewLocalData(localSqlPath).Source()
	default:
		panic("not support source type")
	}
	for _, cv := range sourceData {
		// create tables
		modelFile := filepath.Join(modelPath, cv.Table+".yaml")
		linkModel = append(linkModel, cv.Table)

		if !util.IsHaveFile(modelFile) {
			newBase := newBaseModel(*cv)
			if config.GetDefaultConfig().EnableGormSerializer {
				newBase = applySerializer(newBase, cv.Table)
			}
			content := util.MustSucc(yaml.Marshal(newBase))
			err := os.WriteFile(
				modelFile,
				content,
				os.ModePerm,
			)
			if err != nil {
				panic(err)
			}
			fmt.Printf("新增数据表 %s \n %s \n", cv.Table, cv.Create)
		} else {
			//  说明模型是存在的那么在不破坏原有模型的配置下输入新的字段,删除不存在的字段,修改字段类型
			oldModel := parseBaseModel(modelFile)
			newModel := newBaseModel(*cv)
			oldModel.TableDefinition[cv.Table] = updateField(oldModel.TableDefinition[cv.Table],
				newModel.TableDefinition[cv.Table], cv.Table)
			// get inhibit col
			inhibitCol := []string{}
			for _, v := range oldModel.TableDefinition[cv.Table] {
				if v.Inhibit == definition.Inhibit {
					inhibitCol = append(inhibitCol, v.ColumnName)
				}
			}
			oldModel.MessageDefinition[cv.Table] = updateMessageField(oldModel.MessageDefinition[cv.Table],
				newModel.MessageDefinition[cv.Table])
			// filter inhibit col
			oldModel.MessageDefinition[cv.Table] = lo.Filter(oldModel.MessageDefinition[cv.Table], func(v definition.MessageField, idx int) bool {
				return !lo.Contains(inhibitCol, v.ColumnName)
			})
			oldModel.StatementDefinition[cv.Table] = newModel.StatementDefinition[cv.Table]
			// get serializer
			serializerMap := map[string]string{}
			if config.GetDefaultConfig().EnableGormSerializer {
				for _, v := range oldModel.TableDefinition[cv.Table] {
					switch v.Type {
					case "time.Time":
						serializerMap[v.ColumnName] = "unixtime"
					}
				}
			}
			// apply serializer
			for k, v := range serializerMap {
				for i, field := range oldModel.MessageDefinition[cv.Table] {
					if field.ColumnName == k {
						oldModel.MessageDefinition[cv.Table][i].Serializer = v
					}
				}
			}

			if config.GetDefaultConfig().EnableGormSerializer {
				oldModel = applySerializer(oldModel, cv.Table)
			}

			content := util.MustSucc(yaml.Marshal(oldModel))
			err := os.WriteFile(
				modelFile,
				content,
				os.ModePerm,
			)
			if err != nil {
				panic(err)
			}
		}
	}
	setLinkModel(linkModel)
}

func setLinkModel(link []string) {
	linkConf := filepath.Join(config.GetDefaultConfig().ProjectPath, "link.conf")
	out := linkConf
	buf := util.MustSucc(yaml.Marshal(&link))
	util.NoError(os.WriteFile(out, buf, os.ModePerm))
}

func parseBaseModel(file string) definition.BaseModel {
	newBase := definition.BaseModel{}
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(content, &newBase)
	if err != nil {
		panic(err)
	}
	return newBase
}

func newBaseModel(cv data_source.Create) definition.BaseModel {
	newBase := definition.BaseModel{}
	// set create statement
	newBase.StatementDefinition = make(map[string]definition.CreateStatement)
	newBase.TableDefinition = make(map[string][]definition.Field)
	newBase.MessageDefinition = make(map[string][]definition.MessageField)
	newBase.StatementDefinition[cv.Table] = definition.CreateStatement{
		Statement: cv.Create,
	}
	mod := parse.ParseSqlToModel(cv.Create,
		parse.Options{config.GetDefaultConfig().EnableProtoBigIntToString})

	for _, field := range mod.SortFields() {
		if _, ok := newBase.TableDefinition[cv.Table]; !ok {
			newBase.TableDefinition[cv.Table] = []definition.Field{}
			newBase.MessageDefinition[cv.Table] = []definition.MessageField{}
		}
		Field := definition.Field{
			ColumnName: field.FieldName,
			Type:       string(field.GoSchema.Type),
			Comment:    field.SqlSchema.Comment,
		}
		newBase.TableDefinition[cv.Table] = append(newBase.TableDefinition[cv.Table], Field)
		messageField := definition.MessageField{
			ColumnName: field.FieldName,
			Type:       string(field.ProtoSchema.Type),
			Comment:    field.ProtoSchema.Comment,
		}
		newBase.MessageDefinition[cv.Table] = append(newBase.MessageDefinition[cv.Table], messageField)
	}
	return newBase
}

func updateField(oldTable, newTable []definition.Field, table string) []definition.Field {
	// 移除旧字段
	upTable := lo.Filter(oldTable, func(oldField definition.Field, idx int) bool {
		if lo.ContainsBy(newTable, func(i definition.Field) bool {
			return i.ColumnName == oldField.ColumnName
		}) {
			return true
		} else {
			fmt.Printf("删除字段 %s.%s \n", table, oldField.ColumnName)
			return false
		}
	})

	// 更新字段
	for _, v := range newTable {
		if !lo.ContainsBy(upTable, func(i definition.Field) bool {
			return i.ColumnName == v.ColumnName
		}) {
			// 新增字段
			upTable = append(upTable, v)
			fmt.Printf("字段新增 %s.%s %s %s \n", table, v.ColumnName, v.Type, v.Comment)
			continue
		}
		// 更新字段
		for i, old := range upTable {
			if v.ColumnName == old.ColumnName {

				if upTable[i].Comment != v.Comment ||
					upTable[i].Type != v.Type || upTable[i].Fromat != v.Fromat {

					upTable[i].Comment = v.Comment
					upTable[i].Type = v.Type
					upTable[i].Fromat = v.Fromat
					fmt.Printf("字段更新 %s.%s %s %s \n", table, v.ColumnName, v.Type, v.Comment)
				}
				break
			}
		}
	}
	return upTable
}

func updateMessageField(oldMessage, newMessage []definition.MessageField) []definition.MessageField {
	// 移除旧字段
	upMessage := lo.Filter(oldMessage, func(oldField definition.MessageField, idx int) bool {
		if lo.ContainsBy(newMessage, func(i definition.MessageField) bool {
			return i.ColumnName == oldField.ColumnName
		}) {
			return true
		} else {
			return false
		}
	})
	// 更新字段
	for _, v := range newMessage {
		if !lo.ContainsBy(upMessage, func(i definition.MessageField) bool {
			return i.ColumnName == v.ColumnName
		}) {
			// 新增字段
			upMessage = append(upMessage, v)
			continue
		}
		// 更新字段
		for i, up := range upMessage {
			if v.ColumnName == up.ColumnName {
				upMessage[i].Comment = v.Comment
				upMessage[i].Type = v.Type
				break
			}
		}
	}
	return upMessage
}

func applySerializer(baseModel definition.BaseModel, tableName string) definition.BaseModel {
	// get serializer
	serializerMap := map[string]string{}
	for _, v := range baseModel.TableDefinition[tableName] {
		switch v.Type {
		case "time.Time":
			serializerMap[v.ColumnName] = "unixtime"
		}
	}
	// apply serializer
	for k, v := range serializerMap {
		for i, field := range baseModel.MessageDefinition[tableName] {
			if field.ColumnName == k {
				baseModel.MessageDefinition[tableName][i].Serializer = v
			}
		}
	}
	return baseModel
}

func getFileApiConf() string {
	return `
# 填写你需要关联的api模块
# 例如:
# - order_management
# - user
# - product
	`
}
