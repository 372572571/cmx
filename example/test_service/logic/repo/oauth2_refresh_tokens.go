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
type Oauth2RefreshTokens struct {
	Id           uint64    `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement;not null" json:"id,omitempty" yaml:"id,omitempty"`
	ClientId     string    `gorm:"column:client_id;type:varchar(100);not null" json:"client_id,omitempty" yaml:"client_id,omitempty"`
	UserId       uint64    `gorm:"column:user_id;type:bigint unsigned;not null" json:"user_id,omitempty" yaml:"user_id,omitempty"`
	RefreshToken string    `gorm:"column:refresh_token;type:varchar(512);uniqueIndex:oauth2_refresh_tokens_refresh_token_unique;not null" json:"refresh_token,omitempty" yaml:"refresh_token,omitempty"`
	ExpiresIn    uint64    `gorm:"column:expires_in;type:bigint unsigned;not null" json:"expires_in,omitempty" yaml:"expires_in,omitempty"`
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime;not null" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
}

func (m *Oauth2RefreshTokens) TableName() string {
	return "oauth2_refresh_tokens"
}

// ----- repo definition -----

type Oauth2RefreshTokensRepo struct {
	db *gorm.DB
	oauth2RefreshTokens
}

func NewOauth2RefreshTokensRepo(db *gorm.DB) *Oauth2RefreshTokensRepo {
	return &Oauth2RefreshTokensRepo{
		db:                  db,
		oauth2RefreshTokens: newOauth2RefreshTokens(db),
	}
}

func NewTableOauth2RefreshTokensRepo(db *gorm.DB) *Oauth2RefreshTokensRepo {
	return &Oauth2RefreshTokensRepo{
		db:                  db,
		oauth2RefreshTokens: *newOauth2RefreshTokens(db).Table("oauth2_refresh_tokens"),
	}
}

// ----- gen gorm -----
type oauth2RefreshTokens struct {
	oauth2RefreshTokensDo oauth2RefreshTokensDo
	ALL                   field.Asterisk
	Id                    field.Uint64
	ClientId              field.String
	UserId                field.Uint64
	RefreshToken          field.String
	ExpiresIn             field.Uint64
	CreatedAt             field.Time
	UpdatedAt             field.Time

	fieldMap map[string]field.Expr
}

func newOauth2RefreshTokens(db *gorm.DB, opts ...gen.DOOption) oauth2RefreshTokens {
	_oauth2RefreshTokens := oauth2RefreshTokens{}

	_oauth2RefreshTokens.oauth2RefreshTokensDo.UseDB(db, opts...)
	_oauth2RefreshTokens.oauth2RefreshTokensDo.UseModel(Oauth2RefreshTokens{})

	tableName := _oauth2RefreshTokens.oauth2RefreshTokensDo.TableName()
	_oauth2RefreshTokens.Id = field.NewUint64(tableName, "id")
	_oauth2RefreshTokens.ClientId = field.NewString(tableName, "client_id")
	_oauth2RefreshTokens.UserId = field.NewUint64(tableName, "user_id")
	_oauth2RefreshTokens.RefreshToken = field.NewString(tableName, "refresh_token")
	_oauth2RefreshTokens.ExpiresIn = field.NewUint64(tableName, "expires_in")
	_oauth2RefreshTokens.CreatedAt = field.NewTime(tableName, "created_at")
	_oauth2RefreshTokens.UpdatedAt = field.NewTime(tableName, "updated_at")
	_oauth2RefreshTokens.fillFieldMap()

	return _oauth2RefreshTokens
}

func (c oauth2RefreshTokens) Table(newTableName string) *oauth2RefreshTokens {
	c.oauth2RefreshTokensDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c oauth2RefreshTokens) As(alias string) *oauth2RefreshTokens {
	c.oauth2RefreshTokensDo.DO = *(c.oauth2RefreshTokensDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *oauth2RefreshTokens) updateTableName(table string) *oauth2RefreshTokens {
	c.ALL = field.NewAsterisk(table)
	c.Id = field.NewUint64(table, "id")
	c.ClientId = field.NewString(table, "client_id")
	c.UserId = field.NewUint64(table, "user_id")
	c.RefreshToken = field.NewString(table, "refresh_token")
	c.ExpiresIn = field.NewUint64(table, "expires_in")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.fillFieldMap()
	return c
}

func (c *oauth2RefreshTokens) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 7)
	c.fieldMap["id"] = c.Id
	c.fieldMap["client_id"] = c.ClientId
	c.fieldMap["user_id"] = c.UserId
	c.fieldMap["refresh_token"] = c.RefreshToken
	c.fieldMap["expires_in"] = c.ExpiresIn
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
}

func (c *oauth2RefreshTokens) WithContext(ctx context.Context) *oauth2RefreshTokensDo {
	return c.oauth2RefreshTokensDo.WithContext(ctx)
}

func (c *oauth2RefreshTokens) CallBackWithContext(ctx context.Context, call func(context.Context, gen.Dao) gen.Dao) *oauth2RefreshTokensDo {
	return c.oauth2RefreshTokensDo.withDO(call(ctx, &c.oauth2RefreshTokensDo.WithContext(ctx).DO))
}

func (c oauth2RefreshTokens) TableName() string { return c.oauth2RefreshTokensDo.TableName() }

func (c oauth2RefreshTokens) Alias() string { return c.oauth2RefreshTokensDo.Alias() }

func (c oauth2RefreshTokens) Columns(cols ...field.Expr) gen.Columns {
	return c.oauth2RefreshTokensDo.Columns(cols...)
}

func (c *oauth2RefreshTokens) GetFieldsByName(fieldName []string) []field.OrderExpr {
	_f := []field.OrderExpr{}
	for _, v := range fieldName {
		_rf, ok := c.GetFieldByName(v)
		if ok {
			_f = append(_f, _rf)
		}
	}
	return _f
}

func (c *oauth2RefreshTokens) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c oauth2RefreshTokens) clone(db *gorm.DB) oauth2RefreshTokens {
	c.oauth2RefreshTokensDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c oauth2RefreshTokens) replaceDB(db *gorm.DB) oauth2RefreshTokens {
	c.oauth2RefreshTokensDo.ReplaceDB(db)
	return c
}

// ----- DO -----
type oauth2RefreshTokensDo struct{ gen.DO }

func (c oauth2RefreshTokensDo) Debug() *oauth2RefreshTokensDo {
	return c.withDO(c.DO.Debug())
}

func (c oauth2RefreshTokensDo) WithContext(ctx context.Context) *oauth2RefreshTokensDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c oauth2RefreshTokensDo) ReadDB() *oauth2RefreshTokensDo {
	return c.Clauses(dbresolver.Read)
}

func (c oauth2RefreshTokensDo) WriteDB() *oauth2RefreshTokensDo {
	return c.Clauses(dbresolver.Write)
}

func (c oauth2RefreshTokensDo) Session(config *gorm.Session) *oauth2RefreshTokensDo {
	return c.withDO(c.DO.Session(config))
}

func (c oauth2RefreshTokensDo) Clauses(conds ...clause.Expression) *oauth2RefreshTokensDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c oauth2RefreshTokensDo) Returning(value interface{}, columns ...string) *oauth2RefreshTokensDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c oauth2RefreshTokensDo) Not(conds ...gen.Condition) *oauth2RefreshTokensDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c oauth2RefreshTokensDo) Or(conds ...gen.Condition) *oauth2RefreshTokensDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c oauth2RefreshTokensDo) Select(conds ...field.Expr) *oauth2RefreshTokensDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c oauth2RefreshTokensDo) Where(conds ...gen.Condition) *oauth2RefreshTokensDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c oauth2RefreshTokensDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *oauth2RefreshTokensDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c oauth2RefreshTokensDo) Order(conds ...field.Expr) *oauth2RefreshTokensDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c oauth2RefreshTokensDo) Distinct(cols ...field.Expr) *oauth2RefreshTokensDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c oauth2RefreshTokensDo) Omit(cols ...field.Expr) *oauth2RefreshTokensDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c oauth2RefreshTokensDo) Join(table schema.Tabler, on ...field.Expr) *oauth2RefreshTokensDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c oauth2RefreshTokensDo) LeftJoin(table schema.Tabler, on ...field.Expr) *oauth2RefreshTokensDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c oauth2RefreshTokensDo) RightJoin(table schema.Tabler, on ...field.Expr) *oauth2RefreshTokensDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c oauth2RefreshTokensDo) Group(cols ...field.Expr) *oauth2RefreshTokensDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c oauth2RefreshTokensDo) Having(conds ...gen.Condition) *oauth2RefreshTokensDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c oauth2RefreshTokensDo) Limit(limit int) *oauth2RefreshTokensDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c oauth2RefreshTokensDo) Offset(offset int) *oauth2RefreshTokensDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c oauth2RefreshTokensDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *oauth2RefreshTokensDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c oauth2RefreshTokensDo) Unscoped() *oauth2RefreshTokensDo {
	return c.withDO(c.DO.Unscoped())
}

func (c oauth2RefreshTokensDo) Create(values ...*Oauth2RefreshTokens) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c oauth2RefreshTokensDo) CreateInBatches(values []*Oauth2RefreshTokens, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c oauth2RefreshTokensDo) Save(values ...*Oauth2RefreshTokens) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c oauth2RefreshTokensDo) First() (*Oauth2RefreshTokens, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*Oauth2RefreshTokens), nil
	}
}

func (c oauth2RefreshTokensDo) Take() (*Oauth2RefreshTokens, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*Oauth2RefreshTokens), nil
	}
}

func (c oauth2RefreshTokensDo) Last() (*Oauth2RefreshTokens, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*Oauth2RefreshTokens), nil
	}
}

func (c oauth2RefreshTokensDo) Find() ([]*Oauth2RefreshTokens, error) {
	result, err := c.DO.Find()
	return result.([]*Oauth2RefreshTokens), err
}

func (c oauth2RefreshTokensDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*Oauth2RefreshTokens, err error) {
	buf := make([]*Oauth2RefreshTokens, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c oauth2RefreshTokensDo) FindInBatches(result *[]*Oauth2RefreshTokens, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c oauth2RefreshTokensDo) Attrs(attrs ...field.AssignExpr) *oauth2RefreshTokensDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c oauth2RefreshTokensDo) Assign(attrs ...field.AssignExpr) *oauth2RefreshTokensDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c oauth2RefreshTokensDo) Joins(fields ...field.RelationField) *oauth2RefreshTokensDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c oauth2RefreshTokensDo) Preload(fields ...field.RelationField) *oauth2RefreshTokensDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c oauth2RefreshTokensDo) FirstOrInit() (*Oauth2RefreshTokens, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*Oauth2RefreshTokens), nil
	}
}

func (c oauth2RefreshTokensDo) FirstOrCreate() (*Oauth2RefreshTokens, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*Oauth2RefreshTokens), nil
	}
}

func (c oauth2RefreshTokensDo) FindByPage(offset int, limit int) (result []*Oauth2RefreshTokens, count int64, err error) {
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

func (c oauth2RefreshTokensDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c oauth2RefreshTokensDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c oauth2RefreshTokensDo) Delete(models ...*Oauth2RefreshTokens) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *oauth2RefreshTokensDo) withDO(do gen.Dao) *oauth2RefreshTokensDo {
	c.DO = *do.(*gen.DO)
	return c
}
