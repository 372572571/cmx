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
type Oauth2AccessTokens struct {
	Id          uint64    `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement;not null" json:"id,omitempty" yaml:"id,omitempty"`
	ClientId    string    `gorm:"column:client_id;type:varchar(100);uniqueIndex:oauth2_access_tokens_client_id_user_id_unique;not null" json:"client_id,omitempty" yaml:"client_id,omitempty"`
	UserId      uint64    `gorm:"column:user_id;type:bigint unsigned;uniqueIndex:oauth2_access_tokens_client_id_user_id_unique;index:user_id;not null" json:"user_id,omitempty" yaml:"user_id,omitempty"`
	AccessToken string    `gorm:"column:access_token;type:varchar(512);uniqueIndex:oauth2_access_tokens_access_token_unique;not null;comment:访问token" json:"access_token,omitempty" yaml:"access_token,omitempty"`
	ExpiresIn   uint64    `gorm:"column:expires_in;type:bigint unsigned;not null" json:"expires_in,omitempty" yaml:"expires_in,omitempty"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime;not null" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
}

func (m *Oauth2AccessTokens) TableName() string {
	return "oauth2_access_tokens"
}

// ----- repo definition -----

type Oauth2AccessTokensRepo struct {
	db *gorm.DB
	oauth2AccessTokens
}

func NewOauth2AccessTokensRepo(db *gorm.DB) *Oauth2AccessTokensRepo {
	return &Oauth2AccessTokensRepo{
		db:                 db,
		oauth2AccessTokens: newOauth2AccessTokens(db),
	}
}

func NewTableOauth2AccessTokensRepo(db *gorm.DB) *Oauth2AccessTokensRepo {
	return &Oauth2AccessTokensRepo{
		db:                 db,
		oauth2AccessTokens: *newOauth2AccessTokens(db).Table("oauth2_access_tokens"),
	}
}

// ----- gen gorm -----
type oauth2AccessTokens struct {
	oauth2AccessTokensDo oauth2AccessTokensDo
	ALL                  field.Asterisk
	Id                   field.Uint64
	ClientId             field.String
	UserId               field.Uint64
	AccessToken          field.String
	ExpiresIn            field.Uint64
	CreatedAt            field.Time
	UpdatedAt            field.Time

	fieldMap map[string]field.Expr
}

func newOauth2AccessTokens(db *gorm.DB, opts ...gen.DOOption) oauth2AccessTokens {
	_oauth2AccessTokens := oauth2AccessTokens{}

	_oauth2AccessTokens.oauth2AccessTokensDo.UseDB(db, opts...)
	_oauth2AccessTokens.oauth2AccessTokensDo.UseModel(Oauth2AccessTokens{})

	tableName := _oauth2AccessTokens.oauth2AccessTokensDo.TableName()
	_oauth2AccessTokens.Id = field.NewUint64(tableName, "id")
	_oauth2AccessTokens.ClientId = field.NewString(tableName, "client_id")
	_oauth2AccessTokens.UserId = field.NewUint64(tableName, "user_id")
	_oauth2AccessTokens.AccessToken = field.NewString(tableName, "access_token")
	_oauth2AccessTokens.ExpiresIn = field.NewUint64(tableName, "expires_in")
	_oauth2AccessTokens.CreatedAt = field.NewTime(tableName, "created_at")
	_oauth2AccessTokens.UpdatedAt = field.NewTime(tableName, "updated_at")
	_oauth2AccessTokens.fillFieldMap()

	return _oauth2AccessTokens
}

func (c oauth2AccessTokens) Table(newTableName string) *oauth2AccessTokens {
	c.oauth2AccessTokensDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c oauth2AccessTokens) As(alias string) *oauth2AccessTokens {
	c.oauth2AccessTokensDo.DO = *(c.oauth2AccessTokensDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *oauth2AccessTokens) updateTableName(table string) *oauth2AccessTokens {
	c.ALL = field.NewAsterisk(table)
	c.Id = field.NewUint64(table, "id")
	c.ClientId = field.NewString(table, "client_id")
	c.UserId = field.NewUint64(table, "user_id")
	c.AccessToken = field.NewString(table, "access_token")
	c.ExpiresIn = field.NewUint64(table, "expires_in")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.fillFieldMap()
	return c
}

func (c *oauth2AccessTokens) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 7)
	c.fieldMap["id"] = c.Id
	c.fieldMap["client_id"] = c.ClientId
	c.fieldMap["user_id"] = c.UserId
	c.fieldMap["access_token"] = c.AccessToken
	c.fieldMap["expires_in"] = c.ExpiresIn
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
}

func (c *oauth2AccessTokens) WithContext(ctx context.Context) *oauth2AccessTokensDo {
	return c.oauth2AccessTokensDo.WithContext(ctx)
}

func (c *oauth2AccessTokens) CallBackWithContext(ctx context.Context, call func(context.Context, gen.Dao) gen.Dao) *oauth2AccessTokensDo {
	return c.oauth2AccessTokensDo.withDO(call(ctx, &c.oauth2AccessTokensDo.WithContext(ctx).DO))
}

func (c oauth2AccessTokens) TableName() string { return c.oauth2AccessTokensDo.TableName() }

func (c oauth2AccessTokens) Alias() string { return c.oauth2AccessTokensDo.Alias() }

func (c oauth2AccessTokens) Columns(cols ...field.Expr) gen.Columns {
	return c.oauth2AccessTokensDo.Columns(cols...)
}

func (c *oauth2AccessTokens) GetFieldsByName(fieldName []string) []field.OrderExpr {
	_f := []field.OrderExpr{}
	for _, v := range fieldName {
		_rf, ok := c.GetFieldByName(v)
		if ok {
			_f = append(_f, _rf)
		}
	}
	return _f
}

func (c *oauth2AccessTokens) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c oauth2AccessTokens) clone(db *gorm.DB) oauth2AccessTokens {
	c.oauth2AccessTokensDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c oauth2AccessTokens) replaceDB(db *gorm.DB) oauth2AccessTokens {
	c.oauth2AccessTokensDo.ReplaceDB(db)
	return c
}

// ----- DO -----
type oauth2AccessTokensDo struct{ gen.DO }

func (c oauth2AccessTokensDo) Debug() *oauth2AccessTokensDo {
	return c.withDO(c.DO.Debug())
}

func (c oauth2AccessTokensDo) WithContext(ctx context.Context) *oauth2AccessTokensDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c oauth2AccessTokensDo) ReadDB() *oauth2AccessTokensDo {
	return c.Clauses(dbresolver.Read)
}

func (c oauth2AccessTokensDo) WriteDB() *oauth2AccessTokensDo {
	return c.Clauses(dbresolver.Write)
}

func (c oauth2AccessTokensDo) Session(config *gorm.Session) *oauth2AccessTokensDo {
	return c.withDO(c.DO.Session(config))
}

func (c oauth2AccessTokensDo) Clauses(conds ...clause.Expression) *oauth2AccessTokensDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c oauth2AccessTokensDo) Returning(value interface{}, columns ...string) *oauth2AccessTokensDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c oauth2AccessTokensDo) Not(conds ...gen.Condition) *oauth2AccessTokensDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c oauth2AccessTokensDo) Or(conds ...gen.Condition) *oauth2AccessTokensDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c oauth2AccessTokensDo) Select(conds ...field.Expr) *oauth2AccessTokensDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c oauth2AccessTokensDo) Where(conds ...gen.Condition) *oauth2AccessTokensDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c oauth2AccessTokensDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *oauth2AccessTokensDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c oauth2AccessTokensDo) Order(conds ...field.Expr) *oauth2AccessTokensDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c oauth2AccessTokensDo) Distinct(cols ...field.Expr) *oauth2AccessTokensDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c oauth2AccessTokensDo) Omit(cols ...field.Expr) *oauth2AccessTokensDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c oauth2AccessTokensDo) Join(table schema.Tabler, on ...field.Expr) *oauth2AccessTokensDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c oauth2AccessTokensDo) LeftJoin(table schema.Tabler, on ...field.Expr) *oauth2AccessTokensDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c oauth2AccessTokensDo) RightJoin(table schema.Tabler, on ...field.Expr) *oauth2AccessTokensDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c oauth2AccessTokensDo) Group(cols ...field.Expr) *oauth2AccessTokensDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c oauth2AccessTokensDo) Having(conds ...gen.Condition) *oauth2AccessTokensDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c oauth2AccessTokensDo) Limit(limit int) *oauth2AccessTokensDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c oauth2AccessTokensDo) Offset(offset int) *oauth2AccessTokensDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c oauth2AccessTokensDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *oauth2AccessTokensDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c oauth2AccessTokensDo) Unscoped() *oauth2AccessTokensDo {
	return c.withDO(c.DO.Unscoped())
}

func (c oauth2AccessTokensDo) Create(values ...*Oauth2AccessTokens) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c oauth2AccessTokensDo) CreateInBatches(values []*Oauth2AccessTokens, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c oauth2AccessTokensDo) Save(values ...*Oauth2AccessTokens) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c oauth2AccessTokensDo) First() (*Oauth2AccessTokens, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*Oauth2AccessTokens), nil
	}
}

func (c oauth2AccessTokensDo) Take() (*Oauth2AccessTokens, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*Oauth2AccessTokens), nil
	}
}

func (c oauth2AccessTokensDo) Last() (*Oauth2AccessTokens, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*Oauth2AccessTokens), nil
	}
}

func (c oauth2AccessTokensDo) Find() ([]*Oauth2AccessTokens, error) {
	result, err := c.DO.Find()
	return result.([]*Oauth2AccessTokens), err
}

func (c oauth2AccessTokensDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*Oauth2AccessTokens, err error) {
	buf := make([]*Oauth2AccessTokens, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c oauth2AccessTokensDo) FindInBatches(result *[]*Oauth2AccessTokens, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c oauth2AccessTokensDo) Attrs(attrs ...field.AssignExpr) *oauth2AccessTokensDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c oauth2AccessTokensDo) Assign(attrs ...field.AssignExpr) *oauth2AccessTokensDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c oauth2AccessTokensDo) Joins(fields ...field.RelationField) *oauth2AccessTokensDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c oauth2AccessTokensDo) Preload(fields ...field.RelationField) *oauth2AccessTokensDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c oauth2AccessTokensDo) FirstOrInit() (*Oauth2AccessTokens, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*Oauth2AccessTokens), nil
	}
}

func (c oauth2AccessTokensDo) FirstOrCreate() (*Oauth2AccessTokens, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*Oauth2AccessTokens), nil
	}
}

func (c oauth2AccessTokensDo) FindByPage(offset int, limit int) (result []*Oauth2AccessTokens, count int64, err error) {
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

func (c oauth2AccessTokensDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c oauth2AccessTokensDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c oauth2AccessTokensDo) Delete(models ...*Oauth2AccessTokens) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *oauth2AccessTokensDo) withDO(do gen.Dao) *oauth2AccessTokensDo {
	c.DO = *do.(*gen.DO)
	return c
}
