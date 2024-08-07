package repo

import (
	"context"
	"encoding/json"
	"github.com/spf13/cast"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"gorm.io/plugin/soft_delete"
	"time"
)

var _ soft_delete.DeletedAt
var _ = time.ANSIC
var _ = cast.StringToDate
var _ = json.Marshal

// ----- oneof definition -----

// ----- model definition  -----
type GeoMap struct {
	Id      string `gorm:"column:id;type:varchar(50);primaryKey;not null" json:"id,omitempty" yaml:"id,omitempty"`
	Pid     string `gorm:"column:pid;type:varchar(250);not null" json:"pid,omitempty" yaml:"pid,omitempty"`
	Deep    string `gorm:"column:deep;type:varchar(250);not null" json:"deep,omitempty" yaml:"deep,omitempty"`
	Name    string `gorm:"column:name;type:varchar(250);not null" json:"name,omitempty" yaml:"name,omitempty"`
	ExtPath string `gorm:"column:ext_path;type:varchar(255);not null" json:"ext_path,omitempty" yaml:"ext_path,omitempty"`
	Geo     string `gorm:"column:geo;type:geometry;not null" json:"geo,omitempty" yaml:"geo,omitempty"`
	Polygon string `gorm:"column:polygon;type:geometry;not null" json:"polygon,omitempty" yaml:"polygon,omitempty"`
}

func (m *GeoMap) TableName() string {
	return "geo_map"
}

// ----- repo definition -----

type GeoMapRepo struct {
	db *gorm.DB
	geoMap
}

func NewGeoMapRepo(db *gorm.DB) *GeoMapRepo {
	return &GeoMapRepo{
		db:     db,
		geoMap: newGeoMap(db),
	}
}

func NewTableGeoMapRepo(db *gorm.DB) *GeoMapRepo {
	return &GeoMapRepo{
		db:     db,
		geoMap: *newGeoMap(db).Table("geo_map"),
	}
}

// ----- gen gorm -----
type geoMap struct {
	geoMapDo geoMapDo
	ALL      field.Asterisk
	Id       field.String
	Pid      field.String
	Deep     field.String
	Name     field.String
	ExtPath  field.String
	Geo      field.String
	Polygon  field.String

	fieldMap map[string]field.Expr
}

func newGeoMap(db *gorm.DB, opts ...gen.DOOption) geoMap {
	_geoMap := geoMap{}

	_geoMap.geoMapDo.UseDB(db, opts...)
	_geoMap.geoMapDo.UseModel(GeoMap{})

	tableName := _geoMap.geoMapDo.TableName()
	_geoMap.Id = field.NewString(tableName, "id")
	_geoMap.Pid = field.NewString(tableName, "pid")
	_geoMap.Deep = field.NewString(tableName, "deep")
	_geoMap.Name = field.NewString(tableName, "name")
	_geoMap.ExtPath = field.NewString(tableName, "ext_path")
	_geoMap.Geo = field.NewString(tableName, "geo")
	_geoMap.Polygon = field.NewString(tableName, "polygon")
	_geoMap.fillFieldMap()

	return _geoMap
}

func (c geoMap) Table(newTableName string) *geoMap {
	c.geoMapDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c geoMap) As(alias string) *geoMap {
	c.geoMapDo.DO = *(c.geoMapDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *geoMap) updateTableName(table string) *geoMap {
	c.ALL = field.NewAsterisk(table)
	c.Id = field.NewString(table, "id")
	c.Pid = field.NewString(table, "pid")
	c.Deep = field.NewString(table, "deep")
	c.Name = field.NewString(table, "name")
	c.ExtPath = field.NewString(table, "ext_path")
	c.Geo = field.NewString(table, "geo")
	c.Polygon = field.NewString(table, "polygon")
	c.fillFieldMap()
	return c
}

func (c *geoMap) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 7)
	c.fieldMap["id"] = c.Id
	c.fieldMap["pid"] = c.Pid
	c.fieldMap["deep"] = c.Deep
	c.fieldMap["name"] = c.Name
	c.fieldMap["ext_path"] = c.ExtPath
	c.fieldMap["geo"] = c.Geo
	c.fieldMap["polygon"] = c.Polygon
}

func (c *geoMap) WithContext(ctx context.Context) *geoMapDo { return c.geoMapDo.WithContext(ctx) }

func (c *geoMap) CallBackWithContext(ctx context.Context, call func(context.Context, gen.Dao) gen.Dao) *geoMapDo {
	return c.geoMapDo.withDO(call(ctx, &c.geoMapDo.WithContext(ctx).DO))
}

func (c geoMap) TableName() string { return c.geoMapDo.TableName() }

func (c geoMap) Alias() string { return c.geoMapDo.Alias() }

func (c geoMap) Columns(cols ...field.Expr) gen.Columns { return c.geoMapDo.Columns(cols...) }

func (c *geoMap) GetFieldsByName(fieldName []string) []field.OrderExpr {
	_f := []field.OrderExpr{}
	for _, v := range fieldName {
		_rf, ok := c.GetFieldByName(v)
		if ok {
			_f = append(_f, _rf)
		}
	}
	return _f
}

func (c *geoMap) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c geoMap) clone(db *gorm.DB) geoMap {
	c.geoMapDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c geoMap) replaceDB(db *gorm.DB) geoMap {
	c.geoMapDo.ReplaceDB(db)
	return c
}

// ----- DO -----
type geoMapDo struct{ gen.DO }

func (c geoMapDo) Debug() *geoMapDo {
	return c.withDO(c.DO.Debug())
}

func (c geoMapDo) WithContext(ctx context.Context) *geoMapDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c geoMapDo) ReadDB() *geoMapDo {
	return c.Clauses(dbresolver.Read)
}

func (c geoMapDo) WriteDB() *geoMapDo {
	return c.Clauses(dbresolver.Write)
}

func (c geoMapDo) Session(config *gorm.Session) *geoMapDo {
	return c.withDO(c.DO.Session(config))
}

func (c geoMapDo) Clauses(conds ...clause.Expression) *geoMapDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c geoMapDo) Returning(value interface{}, columns ...string) *geoMapDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c geoMapDo) Not(conds ...gen.Condition) *geoMapDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c geoMapDo) Or(conds ...gen.Condition) *geoMapDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c geoMapDo) Select(conds ...field.Expr) *geoMapDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c geoMapDo) Where(conds ...gen.Condition) *geoMapDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c geoMapDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *geoMapDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c geoMapDo) Order(conds ...field.Expr) *geoMapDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c geoMapDo) Distinct(cols ...field.Expr) *geoMapDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c geoMapDo) Omit(cols ...field.Expr) *geoMapDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c geoMapDo) Join(table schema.Tabler, on ...field.Expr) *geoMapDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c geoMapDo) LeftJoin(table schema.Tabler, on ...field.Expr) *geoMapDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c geoMapDo) RightJoin(table schema.Tabler, on ...field.Expr) *geoMapDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c geoMapDo) Group(cols ...field.Expr) *geoMapDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c geoMapDo) Having(conds ...gen.Condition) *geoMapDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c geoMapDo) Limit(limit int) *geoMapDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c geoMapDo) Offset(offset int) *geoMapDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c geoMapDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *geoMapDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c geoMapDo) Unscoped() *geoMapDo {
	return c.withDO(c.DO.Unscoped())
}

func (c geoMapDo) Create(values ...*GeoMap) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c geoMapDo) CreateInBatches(values []*GeoMap, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c geoMapDo) Save(values ...*GeoMap) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c geoMapDo) First() (*GeoMap, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*GeoMap), nil
	}
}

func (c geoMapDo) Take() (*GeoMap, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*GeoMap), nil
	}
}

func (c geoMapDo) Last() (*GeoMap, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*GeoMap), nil
	}
}

func (c geoMapDo) Find() ([]*GeoMap, error) {
	result, err := c.DO.Find()
	return result.([]*GeoMap), err
}

func (c geoMapDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*GeoMap, err error) {
	buf := make([]*GeoMap, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c geoMapDo) FindInBatches(result *[]*GeoMap, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c geoMapDo) Attrs(attrs ...field.AssignExpr) *geoMapDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c geoMapDo) Assign(attrs ...field.AssignExpr) *geoMapDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c geoMapDo) Joins(fields ...field.RelationField) *geoMapDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c geoMapDo) Preload(fields ...field.RelationField) *geoMapDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c geoMapDo) FirstOrInit() (*GeoMap, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*GeoMap), nil
	}
}

func (c geoMapDo) FirstOrCreate() (*GeoMap, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*GeoMap), nil
	}
}

func (c geoMapDo) FindByPage(offset int, limit int) (result []*GeoMap, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c geoMapDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c geoMapDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c geoMapDo) Delete(models ...*GeoMap) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *geoMapDo) withDO(do gen.Dao) *geoMapDo {
	c.DO = *do.(*gen.DO)
	return c
}
