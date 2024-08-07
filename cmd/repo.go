package cmd

import (
	"cmx/echo/echo_repo"
	"cmx/pkg/config"
	"cmx/pkg/util"
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
	util.NoError(yaml.Unmarshal(linkBuf, &link))

	// definition := config.GetDefaultConfig().GetDefinition()
	for _, v := range link {
		mr := fmt.Sprintf("model.%s.%s", v, v)
		repoInfo := util.MustSucc(echo_repo.Generated(mr))
		fmt.Printf("write path: %s model: %s.go\n", config.GetDefaultConfig().RepoConfig.OutputPath, v)
		err = os.WriteFile(
			path.Join(config.GetDefaultConfig().RepoConfig.OutputPath, v+".go"),
			repoInfo, os.ModePerm,
		)
		if err != nil {
			panic(err)
		}
	}
}
