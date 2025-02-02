package cmd

import (
	"cmx/pkg/config"
	"cmx/pkg/config/definition"
	"cmx/pkg/util"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// 消息名称确认流程
var createMsgGetNameSteps = []*survey.Question{
	{
		Name:   "message",
		Prompt: &survey.Input{Message: "message name: "},
	},
}

var createMsgQuestions = []*survey.Question{
	// 选择类型
	{
		Name: "select_message_type",
		Prompt: &survey.Select{
			Message: "select message type",
			Options: []string{
				string(definition.RefTypeMessageField),
				string(definition.RefTypeMessage),
			},
			Default: string(definition.RefTypeMessageField),
		},
	},
	{
		Name:   "ref",
		Prompt: &survey.Input{Message: "引用Ref: "},
	},
	{
		Name:   "is continued",
		Prompt: &survey.Confirm{Message: "是否继续添加字段"},
	},
	{
		Name:   "out path",
		Prompt: &survey.Input{Message: "输出路径: "},
		// Prompt: &survey.Confirm{Message: "是否继续添加字段"},
	},
}

type createMsgAnswers struct {
	Message     string
	Type        string `survey:"select_message_type"`
	Ref         string
	Tips        []definition.MessageField
	Field       []definition.MessageField
	Interaction bool
	OutPath     string
}

var createMsgCmd = &cobra.Command{
	Use: "create-msg",
	Run: func(cmd *cobra.Command, args []string) {
		answers := createMsgAnswers{}
		// 确认消息名称
		survey.AskOne(createMsgGetNameSteps[0].Prompt, &answers.Message)
		msgData := map[string][]definition.MessageField{}
		msgData[answers.Message] = []definition.MessageField{}
		// 添加字段选择
		answers.Interaction = true
		for answers.Interaction {
			// 确定引用
			q1 := []*survey.Question{}
			survey.Ask(append(q1,
				createMsgQuestions[0],
				createMsgQuestions[1]),
				&answers)
			switch definition.MessageReferenceType(answers.Type) {
			case definition.RefTypeMessageField:
				answers.TipMessagesField(answers.Ref)
			case definition.RefTypeMessage:
				answers.TipMessages(answers.Ref)
			default:
				panic(fmt.Errorf("not support type: %s", answers.Type))
			}
			// 确定是否退出
			survey.AskOne(createMsgQuestions[2].Prompt, &answers.Interaction)
			if !answers.Interaction {
				// 设置输出路径
				survey.AskOne(createMsgQuestions[3].Prompt, &answers.OutPath)
			}
		}
		msgData[answers.Message] = answers.Field
		text, _ := yaml.Marshal(msgData)
		if answers.OutPath == "" {
			fmt.Println(string(text))
		} else {
			out := util.MustSucc(
				os.OpenFile(answers.OutPath, os.O_CREATE|os.O_WRONLY, 0666),
			)
			// defer out.Close()
			out.Truncate(0)
			out.WriteString(string(text))
			log.Println("输出完毕.", out.Name())
		}
	},
}

func (cm *createMsgAnswers) TipMessages(ref string) (success bool) {
	df := config.GetDefaultConfig().GetDefinition()
	reference := config.NewReferenceInformation(ref)
	values, is := df.GetMessagesBySpecify(ref)
	if !is || len(values) == 0 {
		return
	}
	content := []string{}

	for _, item := range values {
		item := fmt.Sprintf(
			"col: %s array: %t one_of: %v desc %s %s",
			item.ColumnName, item.Array, item.OneOf.Select, item.Comment, item.DetailComment)
		content = append(content, item)
	}
	tip := &survey.Confirm{
		Message: "确认对象: \n" + strings.Join(content, "\n"),
	}
	cf := false
	survey.AskOne(tip, &cf)
	if cf {
		cm.Field = append(cm.Field, definition.MessageField{
			ColumnName: reference.Field,
			Ref: definition.MessageReference{
				Type: definition.RefTypeMessage,
				Ref:  ref,
			},
		})
	}
	return true
}

func (cm *createMsgAnswers) TipMessagesField(
	ref string) (success bool) {
	df := config.GetDefaultConfig().GetDefinition()
	values, is := df.GetMessagesBySpecify(ref)
	if !is || len(values) == 0 {
		return
	}
	content := map[string]definition.MessageField{}
	lo.ForEach(values, func(item definition.MessageField, index int) {
		in := fmt.Sprintf(
			"col: %s array: %t one_of: %v desc %s %s",
			item.ColumnName, item.Array, item.OneOf.Select, item.Comment, item.DetailComment)
		content[in] = item
	})
	tip := &survey.MultiSelect{
		Message: "选择多个属性",
		Options: lo.Keys(content),
	}
	selects := []string{}
	survey.AskOne(tip, &selects)
	lo.ForEach(selects, func(item string, index int) {
		cm.Field = append(cm.Field, definition.MessageField{
			ColumnName: content[item].ColumnName,
			Array:      content[item].Array,
			Ref: definition.MessageReference{
				Type:   definition.RefTypeMessageField,
				Ref:    ref,
				Select: []string{content[item].ColumnName},
			},
		})
	})
	return true
}
