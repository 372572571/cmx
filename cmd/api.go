package cmd

import (
	"cmx/echo/echo_api"
	"cmx/echo/echo_stores"
	"cmx/pkg/config"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "create api",
	Long:  `create api`,
	Run: func(cmd *cobra.Command, args []string) {
		answers := struct {
			is bool
		}{}
		err := survey.AskOne(&survey.Confirm{Message: "Confirm api init"}, &answers.is)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if !answers.is {
			fmt.Println("no init api")
			return
		}
		cfg := config.GetDefaultConfig()
		groups := echo_stores.GetSelectApiConf(cfg.SelectApi, cfg)
		if cfg.StoresConfig.IsEnable {
			stores(cfg)
		}
		if cfg.Apis[cfg.SelectApi].IsJoin {
			ApiJoin(groups, cfg)
		} else {
			apiAlone(groups, cfg)
		}
	},
}

func apiAlone(groups []string, cfg config.Config) {
	for _, v := range groups {
		bus, err := echo_api.Generated(v)
		if err != nil {
			panic(err)
		}
		outPath := strings.ReplaceAll(cfg.Apis[cfg.SelectApi].ProtoConfig.OutputPath, "${group}", v)
		os.MkdirAll(outPath, os.ModePerm)
		// alone echo
		for kk, vv := range bus {
			err = os.WriteFile(
				path.Join(outPath, fmt.Sprintf("%s.proto", kk+"_alone")),
				[]byte(vv), os.ModePerm,
			)
			fmt.Println(path.Join(outPath, fmt.Sprintf("%s.proto", kk+"_alone")))
			if err != nil {
				panic(err)
			}
		}
	}
}

func ApiJoin(groups []string, cfg config.Config) {
	for _, v := range groups {
		bus, err := echo_api.GeneratedGroup(v)
		if err != nil {
			panic(err)
		}
		outPath := strings.ReplaceAll(cfg.Apis[cfg.SelectApi].ProtoConfig.OutputPath, "${group}", v)
		os.MkdirAll(outPath, os.ModePerm)
		// alone echo
		err = os.WriteFile(
			path.Join(outPath, fmt.Sprintf("%s.proto", v+"_join")),
			[]byte(bus), os.ModePerm,
		)
		fmt.Println(path.Join(outPath, fmt.Sprintf("%s.proto", v+"_join")))
		if err != nil {
			panic(err)
		}
	}
}

// stores 输出被引用的模型
func stores(cfg config.Config) {
	str, err := echo_stores.Generated(cfg)
	if err != nil {
		panic(err)
	}
	storesConfig := cfg.StoresConfig
	outPath := path.Join(storesConfig.ProtoConfig.OutputPath)
	os.MkdirAll(outPath, os.ModePerm)
	err = os.WriteFile(
		path.Join(outPath, strings.Split(storesConfig.StoresName, ".")[0]+".proto"),
		[]byte(str), os.ModePerm,
	)
	if err != nil {
		panic(err)
	}
}
