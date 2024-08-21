package data_source

import (
	"fmt"
	config "cmx/v1/logic/aggregate/build_config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlData struct {
	dataBaseConfig config.DataBaseConfig
}

func NewMysqlData(dataBaseConfig config.DataBaseConfig) IDataSource {
	return &MysqlData{dataBaseConfig: dataBaseConfig}
}

func (m *MysqlData) Source() []*Create {
	// database init dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.dataBaseConfig.User,
		m.dataBaseConfig.Pwd,
		m.dataBaseConfig.Host,
		m.dataBaseConfig.Port,
		m.dataBaseConfig.DbName,
	)
	// connect
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(db)
	}
	// get create statement
	tbs := []*Tables{}
	err = db.Raw("SELECT * FROM `information_schema`.`tables` where `table_schema` = ? ORDER BY CREATE_TIME desc",
		m.dataBaseConfig.DbName).Scan(&tbs).Error
	if err != nil {
		panic(err)
	}
	creates := []*Create{}
	for _, v := range tbs {
		cv := Create{}
		err = db.Raw("show create table " + v.TABLE_NAME).Scan(&cv).Error
		if err != nil {
			panic(err)
		}
		creates = append(creates, &cv)
	}
	return creates
}
