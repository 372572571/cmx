package data_source

type ESourceType = int32

const (
	SourceTypeMysql ESourceType = iota
	SourceTypeLocal ESourceType = 1
)

type IDataSource interface {
	Source() []*Create
}
