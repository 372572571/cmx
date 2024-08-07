package newrepo

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
type UserSubjects struct {
	Id              int64                 `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement;not null;comment:自增id" json:"id,omitempty" yaml:"id,omitempty"`
	UserId          int64                 `gorm:"column:user_id;type:bigint(20);index:idx_user_id_company_name;not null;comment:user_id" json:"user_id,omitempty" yaml:"user_id,omitempty"`
	CompanyName     string                `gorm:"column:company_name;type:varchar(24);index:idx_user_id_company_name;not null;comment:公司名称" json:"company_name,omitempty" yaml:"company_name,omitempty"`
	JuridicalPerson string                `gorm:"column:juridical_person;type:varchar(24);not null;comment:公司法人名称" json:"juridical_person,omitempty" yaml:"juridical_person,omitempty"`
	IdCard          string                `gorm:"column:id_card;type:varchar(32);not null;comment:身份证号" json:"id_card,omitempty" yaml:"id_card,omitempty"`
	UniformCode     string                `gorm:"column:uniform_code;type:varchar(32);not null;comment:统一信用代码" json:"uniform_code,omitempty" yaml:"uniform_code,omitempty"`
	Address         string                `gorm:"column:address;type:varchar(108);not null;comment:公司地址" json:"address,omitempty" yaml:"address,omitempty"`
	Capital         int                   `gorm:"column:capital;type:int(10) unsigned;not null;comment:注册资金" json:"capital,omitempty" yaml:"capital,omitempty"`
	CreationTime    int64                 `gorm:"column:creation_time;type:bigint(20);not null;comment:公司创建时间" json:"creation_time,omitempty" yaml:"creation_time,omitempty"`
	Scopes          string                `gorm:"column:scopes;type:varchar(512);not null;comment:经营范围" json:"scopes,omitempty" yaml:"scopes,omitempty"`
	License         string                `gorm:"column:license;type:varchar(124);not null;comment:营业执照" json:"license,omitempty" yaml:"license,omitempty"`
	Photo           string                `gorm:"column:photo;type:varchar(2048);not null;comment:公司环境" json:"photo,omitempty" yaml:"photo,omitempty"`
	Status          int                   `gorm:"column:status;type:tinyint(4);not null;comment:是否当前选择主体[1默认 2选中]" json:"status,omitempty" yaml:"status,omitempty"`
	DeletedAt       soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint(20);not null" json:"deleted_at,omitempty" yaml:"deleted_at,omitempty"`
	UpdatedAt       time.Time             `gorm:"column:updated_at;type:datetime;not null;comment:更新时间" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	CreatedAt       time.Time             `gorm:"column:created_at;type:datetime;not null" json:"created_at,omitempty" yaml:"created_at,omitempty"`
}

func (m *UserSubjects) TableName() string {
	return "user_subjects"
}

// ----- repo definition -----

type UserSubjectsRepo struct {
	db *gorm.DB
	userSubjects
}

func NewUserSubjectsRepo(db *gorm.DB) *UserSubjectsRepo {
	return &UserSubjectsRepo{
		db:           db,
		userSubjects: newUserSubjects(db),
	}
}

// not enable model to proto

// ----- gen gorm -----
type userSubjects struct {
	userSubjectsDo  userSubjectsDo
	ALL             field.Asterisk
	Id              field.Int64
	UserId          field.Int64
	CompanyName     field.String
	JuridicalPerson field.String
	IdCard          field.String
	UniformCode     field.String
	Address         field.String
	Capital         field.Int
	CreationTime    field.Int64
	Scopes          field.String
	License         field.String
	Photo           field.String
	Status          field.Int
	DeletedAt       field.Field
	UpdatedAt       field.Time
	CreatedAt       field.Time

	fieldMap map[string]field.Expr
}

func newUserSubjects(db *gorm.DB, opts ...gen.DOOption) userSubjects {
	_userSubjects := userSubjects{}

	_userSubjects.userSubjectsDo.UseDB(db, opts...)
	_userSubjects.userSubjectsDo.UseModel(UserSubjects{})

	tableName := _userSubjects.userSubjectsDo.TableName()
	_userSubjects.Id = field.NewInt64(tableName, "id")
	_userSubjects.UserId = field.NewInt64(tableName, "user_id")
	_userSubjects.CompanyName = field.NewString(tableName, "company_name")
	_userSubjects.JuridicalPerson = field.NewString(tableName, "juridical_person")
	_userSubjects.IdCard = field.NewString(tableName, "id_card")
	_userSubjects.UniformCode = field.NewString(tableName, "uniform_code")
	_userSubjects.Address = field.NewString(tableName, "address")
	_userSubjects.Capital = field.NewInt(tableName, "capital")
	_userSubjects.CreationTime = field.NewInt64(tableName, "creation_time")
	_userSubjects.Scopes = field.NewString(tableName, "scopes")
	_userSubjects.License = field.NewString(tableName, "license")
	_userSubjects.Photo = field.NewString(tableName, "photo")
	_userSubjects.Status = field.NewInt(tableName, "status")
	_userSubjects.DeletedAt = field.NewField(tableName, "deleted_at")
	_userSubjects.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userSubjects.CreatedAt = field.NewTime(tableName, "created_at")
	_userSubjects.fillFieldMap()

	return _userSubjects
}

func (c userSubjects) Table(newTableName string) *userSubjects {
	c.userSubjectsDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c userSubjects) As(alias string) *userSubjects {
	c.userSubjectsDo.DO = *(c.userSubjectsDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *userSubjects) updateTableName(table string) *userSubjects {
	c.ALL = field.NewAsterisk(table)
	c.Id = field.NewInt64(table, "id")
	c.UserId = field.NewInt64(table, "user_id")
	c.CompanyName = field.NewString(table, "company_name")
	c.JuridicalPerson = field.NewString(table, "juridical_person")
	c.IdCard = field.NewString(table, "id_card")
	c.UniformCode = field.NewString(table, "uniform_code")
	c.Address = field.NewString(table, "address")
	c.Capital = field.NewInt(table, "capital")
	c.CreationTime = field.NewInt64(table, "creation_time")
	c.Scopes = field.NewString(table, "scopes")
	c.License = field.NewString(table, "license")
	c.Photo = field.NewString(table, "photo")
	c.Status = field.NewInt(table, "status")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.fillFieldMap()
	return c
}

func (c *userSubjects) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 16)
	c.fieldMap["id"] = c.Id
	c.fieldMap["user_id"] = c.UserId
	c.fieldMap["company_name"] = c.CompanyName
	c.fieldMap["juridical_person"] = c.JuridicalPerson
	c.fieldMap["id_card"] = c.IdCard
	c.fieldMap["uniform_code"] = c.UniformCode
	c.fieldMap["address"] = c.Address
	c.fieldMap["capital"] = c.Capital
	c.fieldMap["creation_time"] = c.CreationTime
	c.fieldMap["scopes"] = c.Scopes
	c.fieldMap["license"] = c.License
	c.fieldMap["photo"] = c.Photo
	c.fieldMap["status"] = c.Status
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["created_at"] = c.CreatedAt
}

func (c *userSubjects) WithContext(ctx context.Context) *userSubjectsDo {
	return c.userSubjectsDo.WithContext(ctx)
}

func (c userSubjects) TableName() string { return c.userSubjectsDo.TableName() }

func (c userSubjects) Alias() string { return c.userSubjectsDo.Alias() }

func (c userSubjects) Columns(cols ...field.Expr) gen.Columns {
	return c.userSubjectsDo.Columns(cols...)
}

func (c *userSubjects) GetFieldsByName(fieldName []string) []field.OrderExpr {
	_f := []field.OrderExpr{}
	for _, v := range fieldName {
		_rf, ok := c.GetFieldByName(v)
		if ok {
			_f = append(_f, _rf)
		}
	}
	return _f
}

func (c *userSubjects) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c userSubjects) clone(db *gorm.DB) userSubjects {
	c.userSubjectsDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c userSubjects) replaceDB(db *gorm.DB) userSubjects {
	c.userSubjectsDo.ReplaceDB(db)
	return c
}

// ----- DO -----
type userSubjectsDo struct{ gen.DO }

func (c userSubjectsDo) Debug() *userSubjectsDo {
	return c.withDO(c.DO.Debug())
}

func (c userSubjectsDo) WithContext(ctx context.Context) *userSubjectsDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c userSubjectsDo) ReadDB() *userSubjectsDo {
	return c.Clauses(dbresolver.Read)
}

func (c userSubjectsDo) WriteDB() *userSubjectsDo {
	return c.Clauses(dbresolver.Write)
}

func (c userSubjectsDo) Session(config *gorm.Session) *userSubjectsDo {
	return c.withDO(c.DO.Session(config))
}

func (c userSubjectsDo) Clauses(conds ...clause.Expression) *userSubjectsDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c userSubjectsDo) Returning(value interface{}, columns ...string) *userSubjectsDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c userSubjectsDo) Not(conds ...gen.Condition) *userSubjectsDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c userSubjectsDo) Or(conds ...gen.Condition) *userSubjectsDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c userSubjectsDo) Select(conds ...field.Expr) *userSubjectsDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c userSubjectsDo) Where(conds ...gen.Condition) *userSubjectsDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c userSubjectsDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *userSubjectsDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c userSubjectsDo) Order(conds ...field.Expr) *userSubjectsDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c userSubjectsDo) Distinct(cols ...field.Expr) *userSubjectsDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c userSubjectsDo) Omit(cols ...field.Expr) *userSubjectsDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c userSubjectsDo) Join(table schema.Tabler, on ...field.Expr) *userSubjectsDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c userSubjectsDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userSubjectsDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c userSubjectsDo) RightJoin(table schema.Tabler, on ...field.Expr) *userSubjectsDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c userSubjectsDo) Group(cols ...field.Expr) *userSubjectsDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c userSubjectsDo) Having(conds ...gen.Condition) *userSubjectsDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c userSubjectsDo) Limit(limit int) *userSubjectsDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c userSubjectsDo) Offset(offset int) *userSubjectsDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c userSubjectsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userSubjectsDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c userSubjectsDo) Unscoped() *userSubjectsDo {
	return c.withDO(c.DO.Unscoped())
}

func (c userSubjectsDo) Create(values ...*UserSubjects) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c userSubjectsDo) CreateInBatches(values []*UserSubjects, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c userSubjectsDo) Save(values ...*UserSubjects) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c userSubjectsDo) First() (*UserSubjects, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*UserSubjects), nil
	}
}

func (c userSubjectsDo) Take() (*UserSubjects, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*UserSubjects), nil
	}
}

func (c userSubjectsDo) Last() (*UserSubjects, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*UserSubjects), nil
	}
}

func (c userSubjectsDo) Find() ([]*UserSubjects, error) {
	result, err := c.DO.Find()
	return result.([]*UserSubjects), err
}

func (c userSubjectsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*UserSubjects, err error) {
	buf := make([]*UserSubjects, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c userSubjectsDo) FindInBatches(result *[]*UserSubjects, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c userSubjectsDo) Attrs(attrs ...field.AssignExpr) *userSubjectsDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c userSubjectsDo) Assign(attrs ...field.AssignExpr) *userSubjectsDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c userSubjectsDo) Joins(fields ...field.RelationField) *userSubjectsDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c userSubjectsDo) Preload(fields ...field.RelationField) *userSubjectsDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c userSubjectsDo) FirstOrInit() (*UserSubjects, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*UserSubjects), nil
	}
}

func (c userSubjectsDo) FirstOrCreate() (*UserSubjects, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*UserSubjects), nil
	}
}

func (c userSubjectsDo) FindByPage(offset int, limit int) (result []*UserSubjects, count int64, err error) {
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

func (c userSubjectsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c userSubjectsDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c userSubjectsDo) Delete(models ...*UserSubjects) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *userSubjectsDo) withDO(do gen.Dao) *userSubjectsDo {
	c.DO = *do.(*gen.DO)
	return c
}
