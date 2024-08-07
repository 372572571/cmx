package data_source

import (
	"cmx/pkg/util"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type LocalData struct {
	localSqlPath string
	Paths        []string
}

func NewLocalData(local string) *LocalData {
	return &LocalData{Paths: util.LoadDirAllFile(local, []string{".sql"})}
}

func (l *LocalData) Source() []*Create {
	creates := []*Create{}
	for _, v := range l.Paths {
		sqlPath := path.Join(l.localSqlPath, v)
		item := &Create{}
		// 获取不带后缀的名称
		base := filepath.Base(sqlPath)
		item.Table = strings.TrimSuffix(base, filepath.Ext(base))
		bys, err := os.ReadFile(sqlPath)
		if err != nil {
			panic(err)
		}
		if len(bys) == 0 {
			continue
		}
		item.Create = string(bys)
		creates = append(creates, item)
	}
	return creates
}
