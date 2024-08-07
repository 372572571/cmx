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
type Oauth2Clients struct {
	Id             uint64    `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement;not null" json:"id,omitempty" yaml:"id,omitempty"`
	ClientId       string    `gorm:"column:client_id;type:varchar(100);uniqueIndex:oauth2_clients_client_id_unique;not null" json:"client_id,omitempty" yaml:"client_id,omitempty"`
	ClientSecret   string    `gorm:"column:client_secret;type:varchar(100);not null" json:"client_secret,omitempty" yaml:"client_secret,omitempty"`
	RedirectUri    string    `gorm:"column:redirect_uri;type:text" json:"redirect_uri,omitempty" yaml:"redirect_uri,omitempty"`
	Privatekey     string    `gorm:"column:privatekey;type:text" json:"privatekey,omitempty" yaml:"privatekey,omitempty"`
	Publickey      string    `gorm:"column:publickey;type:text" json:"publickey,omitempty" yaml:"publickey,omitempty"`
	CreatedAt      time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:datetime;not null" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	InternalSecret string    `gorm:"column:internal_secret;type:varchar(255);not null;default:内部 secret" json:"internal_secret,omitempty" yaml:"internal_secret,omitempty"`
}

func (m *Oauth2Clients) TableName() string {
	return "oauth2_clients"
}

// ----- repo definition -----

type Oauth2ClientsRepo struct {
	db *gorm.DB
	oauth2Clients
}

func NewOauth2ClientsRepo(db *gorm.DB) *Oauth2ClientsRepo {
	return &Oauth2ClientsRepo{
		db:            db,
		oauth2Clients: newOauth2Clients(db),
	}
}

func NewTableOauth2ClientsRepo(db *gorm.DB) *Oauth2ClientsRepo {
	return &Oauth2ClientsRepo{
		db:            db,
		oauth2Clients: *newOauth2Clients(db).Table("oauth2_clients"),
	}
}

// ----- gen gorm -----
type oauth2Clients struct {
	oauth2ClientsDo oauth2ClientsDo
	ALL             field.Asterisk
	Id              field.Uint64
	ClientId        field.String
	ClientSecret    field.String
	RedirectUri     field.String
	Privatekey      field.String
	Publickey       field.String
	CreatedAt       field.Time
	UpdatedAt       field.Time
	InternalSecret  field.String

	fieldMap map[string]field.Expr
}

func newOauth2Clients(db *gorm.DB, opts ...gen.DOOption) oauth2Clients {
	_oauth2Clients := oauth2Clients{}

	_oauth2Clients.oauth2ClientsDo.UseDB(db, opts...)
	_oauth2Clients.oauth2ClientsDo.UseModel(Oauth2Clients{})

	tableName := _oauth2Clients.oauth2ClientsDo.TableName()
	_oauth2Clients.Id = field.NewUint64(tableName, "id")
	_oauth2Clients.ClientId = field.NewString(tableName, "client_id")
	_oauth2Clients.ClientSecret = field.NewString(tableName, "client_secret")
	_oauth2Clients.RedirectUri = field.NewString(tableName, "redirect_uri")
	_oauth2Clients.Privatekey = field.NewString(tableName, "privatekey")
	_oauth2Clients.Publickey = field.NewString(tableName, "publickey")
	_oauth2Clients.CreatedAt = field.NewTime(tableName, "created_at")
	_oauth2Clients.UpdatedAt = field.NewTime(tableName, "updated_at")
	_oauth2Clients.InternalSecret = field.NewString(tableName, "internal_secret")
	_oauth2Clients.fillFieldMap()

	return _oauth2Clients
}

func (c oauth2Clients) Table(newTableName string) *oauth2Clients {
	c.oauth2ClientsDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c oauth2Clients) As(alias string) *oauth2Clients {
	c.oauth2ClientsDo.DO = *(c.oauth2ClientsDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *oauth2Clients) updateTableName(table string) *oauth2Clients {
	c.ALL = field.NewAsterisk(table)
	c.Id = field.NewUint64(table, "id")
	c.ClientId = field.NewString(table, "client_id")
	c.ClientSecret = field.NewString(table, "client_secret")
	c.RedirectUri = field.NewString(table, "redirect_uri")
	c.Privatekey = field.NewString(table, "privatekey")
	c.Publickey = field.NewString(table, "publickey")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.InternalSecret = field.NewString(table, "internal_secret")
	c.fillFieldMap()
	return c
}

func (c *oauth2Clients) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 9)
	c.fieldMap["id"] = c.Id
	c.fieldMap["client_id"] = c.ClientId
	c.fieldMap["client_secret"] = c.ClientSecret
	c.fieldMap["redirect_uri"] = c.RedirectUri
	c.fieldMap["privatekey"] = c.Privatekey
	c.fieldMap["publickey"] = c.Publickey
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["internal_secret"] = c.InternalSecret
}

func (c *oauth2Clients) WithContext(ctx context.Context) *oauth2ClientsDo {
	return c.oauth2ClientsDo.WithContext(ctx)
}

func (c *oauth2Clients) CallBackWithContext(ctx context.Context, call func(context.Context, gen.Dao) gen.Dao) *oauth2ClientsDo {
	return c.oauth2ClientsDo.withDO(call(ctx, &c.oauth2ClientsDo.WithContext(ctx).DO))
}

func (c oauth2Clients) TableName() string { return c.oauth2ClientsDo.TableName() }

func (c oauth2Clients) Alias() string { return c.oauth2ClientsDo.Alias() }

func (c oauth2Clients) Columns(cols ...field.Expr) gen.Columns {
	return c.oauth2ClientsDo.Columns(cols...)
}

func (c *oauth2Clients) GetFieldsByName(fieldName []string) []field.OrderExpr {
	_f := []field.OrderExpr{}
	for _, v := range fieldName {
		_rf, ok := c.GetFieldByName(v)
		if ok {
			_f = append(_f, _rf)
		}
	}
	return _f
}

func (c *oauth2Clients) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c oauth2Clients) clone(db *gorm.DB) oauth2Clients {
	c.oauth2ClientsDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c oauth2Clients) replaceDB(db *gorm.DB) oauth2Clients {
	c.oauth2ClientsDo.ReplaceDB(db)
	return c
}

// ----- DO -----
type oauth2ClientsDo struct{ gen.DO }

func (c oauth2ClientsDo) Debug() *oauth2ClientsDo {
	return c.withDO(c.DO.Debug())
}

func (c oauth2ClientsDo) WithContext(ctx context.Context) *oauth2ClientsDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c oauth2ClientsDo) ReadDB() *oauth2ClientsDo {
	return c.Clauses(dbresolver.Read)
}

func (c oauth2ClientsDo) WriteDB() *oauth2ClientsDo {
	return c.Clauses(dbresolver.Write)
}

func (c oauth2ClientsDo) Session(config *gorm.Session) *oauth2ClientsDo {
	return c.withDO(c.DO.Session(config))
}

func (c oauth2ClientsDo) Clauses(conds ...clause.Expression) *oauth2ClientsDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c oauth2ClientsDo) Returning(value interface{}, columns ...string) *oauth2ClientsDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c oauth2ClientsDo) Not(conds ...gen.Condition) *oauth2ClientsDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c oauth2ClientsDo) Or(conds ...gen.Condition) *oauth2ClientsDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c oauth2ClientsDo) Select(conds ...field.Expr) *oauth2ClientsDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c oauth2ClientsDo) Where(conds ...gen.Condition) *oauth2ClientsDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c oauth2ClientsDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *oauth2ClientsDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c oauth2ClientsDo) Order(conds ...field.Expr) *oauth2ClientsDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c oauth2ClientsDo) Distinct(cols ...field.Expr) *oauth2ClientsDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c oauth2ClientsDo) Omit(cols ...field.Expr) *oauth2ClientsDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c oauth2ClientsDo) Join(table schema.Tabler, on ...field.Expr) *oauth2ClientsDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c oauth2ClientsDo) LeftJoin(table schema.Tabler, on ...field.Expr) *oauth2ClientsDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c oauth2ClientsDo) RightJoin(table schema.Tabler, on ...field.Expr) *oauth2ClientsDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c oauth2ClientsDo) Group(cols ...field.Expr) *oauth2ClientsDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c oauth2ClientsDo) Having(conds ...gen.Condition) *oauth2ClientsDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c oauth2ClientsDo) Limit(limit int) *oauth2ClientsDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c oauth2ClientsDo) Offset(offset int) *oauth2ClientsDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c oauth2ClientsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *oauth2ClientsDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c oauth2ClientsDo) Unscoped() *oauth2ClientsDo {
	return c.withDO(c.DO.Unscoped())
}

func (c oauth2ClientsDo) Create(values ...*Oauth2Clients) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c oauth2ClientsDo) CreateInBatches(values []*Oauth2Clients, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c oauth2ClientsDo) Save(values ...*Oauth2Clients) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c oauth2ClientsDo) First() (*Oauth2Clients, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*Oauth2Clients), nil
	}
}

func (c oauth2ClientsDo) Take() (*Oauth2Clients, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*Oauth2Clients), nil
	}
}

func (c oauth2ClientsDo) Last() (*Oauth2Clients, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*Oauth2Clients), nil
	}
}

func (c oauth2ClientsDo) Find() ([]*Oauth2Clients, error) {
	result, err := c.DO.Find()
	return result.([]*Oauth2Clients), err
}

func (c oauth2ClientsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*Oauth2Clients, err error) {
	buf := make([]*Oauth2Clients, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c oauth2ClientsDo) FindInBatches(result *[]*Oauth2Clients, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c oauth2ClientsDo) Attrs(attrs ...field.AssignExpr) *oauth2ClientsDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c oauth2ClientsDo) Assign(attrs ...field.AssignExpr) *oauth2ClientsDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c oauth2ClientsDo) Joins(fields ...field.RelationField) *oauth2ClientsDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c oauth2ClientsDo) Preload(fields ...field.RelationField) *oauth2ClientsDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c oauth2ClientsDo) FirstOrInit() (*Oauth2Clients, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*Oauth2Clients), nil
	}
}

func (c oauth2ClientsDo) FirstOrCreate() (*Oauth2Clients, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*Oauth2Clients), nil
	}
}

func (c oauth2ClientsDo) FindByPage(offset int, limit int) (result []*Oauth2Clients, count int64, err error) {
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

func (c oauth2ClientsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c oauth2ClientsDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c oauth2ClientsDo) Delete(models ...*Oauth2Clients) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *oauth2ClientsDo) withDO(do gen.Dao) *oauth2ClientsDo {
	c.DO = *do.(*gen.DO)
	return c
}
