package data_source

type Tables struct {
	TABLE_CATALOG   string `gorm:"column:TABLE_CATALOG"`
	TABLE_SCHEMA    string `gorm:"column:TABLE_SCHEMA"`
	TABLE_NAME      string `gorm:"column:TABLE_NAME"`
	TABLE_TYPE      string `gorm:"column:TABLE_TYPE"`
	ENGINE          string `gorm:"column:ENGINE"`
	VERSION         string `gorm:"column:VERSION"`
	ROW_FORMAT      string `gorm:"column:ROW_FORMAT"`
	TABLE_ROWS      string `gorm:"column:TABLE_ROWS"`
	AVG_ROW_LENGTH  string `gorm:"column:AVG_ROW_LENGTH"`
	DATA_LENGTH     string `gorm:"column:DATA_LENGTH"`
	MAX_DATA_LENGTH string `gorm:"column:MAX_DATA_LENGTH"`
	INDEX_LENGTH    string `gorm:"column:INDEX_LENGTH"`
	DATA_FREE       string `gorm:"column:DATA_FREE"`
	AUTO_INCREMENT  string `gorm:"column:AUTO_INCREMENT"`
	CREATE_TIME     string `gorm:"column:CREATE_TIME"`
	UPDATE_TIME     string `gorm:"column:UPDATE_TIME"`
	CHECK_TIME      string `gorm:"column:CHECK_TIME"`
	TABLE_COLLATION string `gorm:"column:TABLE_COLLATION"`
	CHECKSUM        string `gorm:"column:CHECKSUM"`
	CREATE_OPTIONS  string `gorm:"column:CREATE_OPTIONS"`
	TABLE_COMMENT   string `gorm:"column:TABLE_COMMENT"`
}

func (*Tables) TableName() string {
	return "TABLES"
}

type Create struct {
	Table  string `gorm:"column:Table"`
	Create string `gorm:"column:Create Table" json:"Create Table"`
}

func (*Create) TableName() string {
	return "Create"
}
