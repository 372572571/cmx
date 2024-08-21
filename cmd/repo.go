package cmd

import (
	config "cmx/v1/logic/aggregate/build_config"
	"cmx/v1/logic/inside/echo/echo_repo"
	"cmx/v1/logic/util"
	"cmx/v1/pkg/logger"
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "create repo model",
	Long:  `create repo model`,
	Run: func(cmd *cobra.Command, args []string) {
		repo()
	},
}

// repo generate gorm repo file
func repo() {
	// link model
	linkBuf, err := os.ReadFile(filepath.Join(config.GetDefaultConfig().ProjectPath, "link.conf"))
	if err != nil {
		panic(err)
	}
	link := []string{}
	err = yaml.Unmarshal(linkBuf, &link)
	if err != nil {
		logger.Fatalf(context.Background(),"unmarshal link.conf error: %s", err.Error())
	}
	// definition := config.GetDefaultConfig().GetDefinition()
	for _, v := range link {
		mr := fmt.Sprintf("model.%s.%s", v, v)
		logger.Infof(context.Background(),"generate model: %s\n", mr)
		repoInfo := util.MustSuccess(echo_repo.Generated(mr))
		logger.Infof(context.Background(),"write path: %s model: %s.go\n", config.GetDefaultConfig().RepoConfig.OutputPath, v)
		err = os.WriteFile(
			path.Join(config.GetDefaultConfig().RepoConfig.OutputPath, v+".go"),
			repoInfo, os.ModePerm,
		)
		if err != nil {
			logger.Fatalf(context.Background(), "write file error: %s", err.Error())
		}
	}
}
