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
type Oauth2AuthorizationCodes struct {
	Id                uint64                `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement;not null;comment:授权码ID" json:"id,omitempty" yaml:"id,omitempty"`
	AuthorizationCode string                `gorm:"column:authorization_code;type:varchar(255);not null;comment:授权码" json:"authorization_code,omitempty" yaml:"authorization_code,omitempty"`
	ClientId          string                `gorm:"column:client_id;type:varchar(100);uniqueIndex:client_id;not null;comment:客户端ID" json:"client_id,omitempty" yaml:"client_id,omitempty"`
	UserId            uint64                `gorm:"column:user_id;type:bigint unsigned;index:user_id;uniqueIndex:client_id;not null;comment:用户ID" json:"user_id,omitempty" yaml:"user_id,omitempty"`
	ExpiresAt         int64                 `gorm:"column:expires_at;type:bigint;not null;comment:过期时间" json:"expires_at,omitempty" yaml:"expires_at,omitempty"`
	CreatedAt         time.Time             `gorm:"column:created_at;type:datetime;not null;comment:创建时间" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt         time.Time             `gorm:"column:updated_at;type:datetime;not null;comment:更新时间" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	DeletedAt         soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint;default:0;comment:删除时间" json:"deleted_at,omitempty" yaml:"deleted_at,omitempty"`
}

func (m *Oauth2AuthorizationCodes) TableName() string {
	return "oauth2_authorization_codes"
}

// ----- repo definition -----

type Oauth2AuthorizationCodesRepo struct {
	db *gorm.DB
	oauth2AuthorizationCodes
}

func NewOauth2AuthorizationCodesRepo(db *gorm.DB) *Oauth2AuthorizationCodesRepo {
	return &Oauth2AuthorizationCodesRepo{
		db:                       db,
		oauth2AuthorizationCodes: newOauth2AuthorizationCodes(db),
	}
}

func NewTableOauth2AuthorizationCodesRepo(db *gorm.DB) *Oauth2AuthorizationCodesRepo {
	return &Oauth2AuthorizationCodesRepo{
		db:                       db,
		oauth2AuthorizationCodes: *newOauth2AuthorizationCodes(db).Table("oauth2_authorization_codes"),
	}
}

// ----- gen gorm -----
type oauth2AuthorizationCodes struct {
	oauth2AuthorizationCodesDo oauth2AuthorizationCodesDo
	ALL                        field.Asterisk
	Id                         field.Uint64
	AuthorizationCode          field.String
	ClientId                   field.String
	UserId                     field.Uint64
	ExpiresAt                  field.Int64
	CreatedAt                  field.Time
	UpdatedAt                  field.Time
	DeletedAt                  field.Field

	fieldMap map[string]field.Expr
}

func newOauth2AuthorizationCodes(db *gorm.DB, opts ...gen.DOOption) oauth2AuthorizationCodes {
	_oauth2AuthorizationCodes := oauth2AuthorizationCodes{}

	_oauth2AuthorizationCodes.oauth2AuthorizationCodesDo.UseDB(db, opts...)
	_oauth2AuthorizationCodes.oauth2AuthorizationCodesDo.UseModel(Oauth2AuthorizationCodes{})

	tableName := _oauth2AuthorizationCodes.oauth2AuthorizationCodesDo.TableName()
	_oauth2AuthorizationCodes.Id = field.NewUint64(tableName, "id")
	_oauth2AuthorizationCodes.AuthorizationCode = field.NewString(tableName, "authorization_code")
	_oauth2AuthorizationCodes.ClientId = field.NewString(tableName, "client_id")
	_oauth2AuthorizationCodes.UserId = field.NewUint64(tableName, "user_id")
	_oauth2AuthorizationCodes.ExpiresAt = field.NewInt64(tableName, "expires_at")
	_oauth2AuthorizationCodes.CreatedAt = field.NewTime(tableName, "created_at")
	_oauth2AuthorizationCodes.UpdatedAt = field.NewTime(tableName, "updated_at")
	_oauth2AuthorizationCodes.DeletedAt = field.NewField(tableName, "deleted_at")
	_oauth2AuthorizationCodes.fillFieldMap()

	return _oauth2AuthorizationCodes
}

func (c oauth2AuthorizationCodes) Table(newTableName string) *oauth2AuthorizationCodes {
	c.oauth2AuthorizationCodesDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c oauth2AuthorizationCodes) As(alias string) *oauth2AuthorizationCodes {
	c.oauth2AuthorizationCodesDo.DO = *(c.oauth2AuthorizationCodesDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *oauth2AuthorizationCodes) updateTableName(table string) *oauth2AuthorizationCodes {
	c.ALL = field.NewAsterisk(table)
	c.Id = field.NewUint64(table, "id")
	c.AuthorizationCode = field.NewString(table, "authorization_code")
	c.ClientId = field.NewString(table, "client_id")
	c.UserId = field.NewUint64(table, "user_id")
	c.ExpiresAt = field.NewInt64(table, "expires_at")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.fillFieldMap()
	return c
}

func (c *oauth2AuthorizationCodes) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 8)
	c.fieldMap["id"] = c.Id
	c.fieldMap["authorization_code"] = c.AuthorizationCode
	c.fieldMap["client_id"] = c.ClientId
	c.fieldMap["user_id"] = c.UserId
	c.fieldMap["expires_at"] = c.ExpiresAt
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
}

func (c *oauth2AuthorizationCodes) WithContext(ctx context.Context) *oauth2AuthorizationCodesDo {
	return c.oauth2AuthorizationCodesDo.WithContext(ctx)
}

func (c *oauth2AuthorizationCodes) CallBackWithContext(ctx context.Context, call func(context.Context, gen.Dao) gen.Dao) *oauth2AuthorizationCodesDo {
	return c.oauth2AuthorizationCodesDo.withDO(call(ctx, &c.oauth2AuthorizationCodesDo.WithContext(ctx).DO))
}

func (c oauth2AuthorizationCodes) TableName() string { return c.oauth2AuthorizationCodesDo.TableName() }

func (c oauth2AuthorizationCodes) Alias() string { return c.oauth2AuthorizationCodesDo.Alias() }

func (c oauth2AuthorizationCodes) Columns(cols ...field.Expr) gen.Columns {
	return c.oauth2AuthorizationCodesDo.Columns(cols...)
}

func (c *oauth2AuthorizationCodes) GetFieldsByName(fieldName []string) []field.OrderExpr {
	_f := []field.OrderExpr{}
	for _, v := range fieldName {
		_rf, ok := c.GetFieldByName(v)
		if ok {
			_f = append(_f, _rf)
		}
	}
	return _f
}

func (c *oauth2AuthorizationCodes) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c oauth2AuthorizationCodes) clone(db *gorm.DB) oauth2AuthorizationCodes {
	c.oauth2AuthorizationCodesDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c oauth2AuthorizationCodes) replaceDB(db *gorm.DB) oauth2AuthorizationCodes {
	c.oauth2AuthorizationCodesDo.ReplaceDB(db)
	return c
}

// ----- DO -----
type oauth2AuthorizationCodesDo struct{ gen.DO }

func (c oauth2AuthorizationCodesDo) Debug() *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.Debug())
}

func (c oauth2AuthorizationCodesDo) WithContext(ctx context.Context) *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c oauth2AuthorizationCodesDo) ReadDB() *oauth2AuthorizationCodesDo {
	return c.Clauses(dbresolver.Read)
}

func (c oauth2AuthorizationCodesDo) WriteDB() *oauth2AuthorizationCodesDo {
	return c.Clauses(dbresolver.Write)
}

func (c oauth2AuthorizationCodesDo) Session(config *gorm.Session) *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.Session(config))
}

func (c oauth2AuthorizationCodesDo) Clauses(conds ...clause.Expression) *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c oauth2AuthorizationCodesDo) Returning(value interface{}, columns ...string) *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c oauth2AuthorizationCodesDo) Not(conds ...gen.Condition) *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c oauth2AuthorizationCodesDo) Or(conds ...gen.Condition) *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c oauth2AuthorizationCodesDo) Select(conds ...field.Expr) *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c oauth2AuthorizationCodesDo) Where(conds ...gen.Condition) *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c oauth2AuthorizationCodesDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *oauth2AuthorizationCodesDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c oauth2AuthorizationCodesDo) Order(conds ...field.Expr) *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c oauth2AuthorizationCodesDo) Distinct(cols ...field.Expr) *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c oauth2AuthorizationCodesDo) Omit(cols ...field.Expr) *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c oauth2AuthorizationCodesDo) Join(table schema.Tabler, on ...field.Expr) *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c oauth2AuthorizationCodesDo) LeftJoin(table schema.Tabler, on ...field.Expr) *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c oauth2AuthorizationCodesDo) RightJoin(table schema.Tabler, on ...field.Expr) *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c oauth2AuthorizationCodesDo) Group(cols ...field.Expr) *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c oauth2AuthorizationCodesDo) Having(conds ...gen.Condition) *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c oauth2AuthorizationCodesDo) Limit(limit int) *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c oauth2AuthorizationCodesDo) Offset(offset int) *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c oauth2AuthorizationCodesDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c oauth2AuthorizationCodesDo) Unscoped() *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.Unscoped())
}

func (c oauth2AuthorizationCodesDo) Create(values ...*Oauth2AuthorizationCodes) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c oauth2AuthorizationCodesDo) CreateInBatches(values []*Oauth2AuthorizationCodes, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c oauth2AuthorizationCodesDo) Save(values ...*Oauth2AuthorizationCodes) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c oauth2AuthorizationCodesDo) First() (*Oauth2AuthorizationCodes, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*Oauth2AuthorizationCodes), nil
	}
}

func (c oauth2AuthorizationCodesDo) Take() (*Oauth2AuthorizationCodes, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*Oauth2AuthorizationCodes), nil
	}
}

func (c oauth2AuthorizationCodesDo) Last() (*Oauth2AuthorizationCodes, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*Oauth2AuthorizationCodes), nil
	}
}

func (c oauth2AuthorizationCodesDo) Find() ([]*Oauth2AuthorizationCodes, error) {
	result, err := c.DO.Find()
	return result.([]*Oauth2AuthorizationCodes), err
}

func (c oauth2AuthorizationCodesDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*Oauth2AuthorizationCodes, err error) {
	buf := make([]*Oauth2AuthorizationCodes, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c oauth2AuthorizationCodesDo) FindInBatches(result *[]*Oauth2AuthorizationCodes, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c oauth2AuthorizationCodesDo) Attrs(attrs ...field.AssignExpr) *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c oauth2AuthorizationCodesDo) Assign(attrs ...field.AssignExpr) *oauth2AuthorizationCodesDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c oauth2AuthorizationCodesDo) Joins(fields ...field.RelationField) *oauth2AuthorizationCodesDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c oauth2AuthorizationCodesDo) Preload(fields ...field.RelationField) *oauth2AuthorizationCodesDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c oauth2AuthorizationCodesDo) FirstOrInit() (*Oauth2AuthorizationCodes, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*Oauth2AuthorizationCodes), nil
	}
}

func (c oauth2AuthorizationCodesDo) FirstOrCreate() (*Oauth2AuthorizationCodes, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*Oauth2AuthorizationCodes), nil
	}
}

func (c oauth2AuthorizationCodesDo) FindByPage(offset int, limit int) (result []*Oauth2AuthorizationCodes, count int64, err error) {
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

func (c oauth2AuthorizationCodesDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c oauth2AuthorizationCodesDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c oauth2AuthorizationCodesDo) Delete(models ...*Oauth2AuthorizationCodes) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *oauth2AuthorizationCodesDo) withDO(do gen.Dao) *oauth2AuthorizationCodesDo {
	c.DO = *do.(*gen.DO)
	return c
}
