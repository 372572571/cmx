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
type User struct {
	Id        uint64                `gorm:"column:id;type:bigint unsigned;primaryKey;not null;comment:用户ID" json:"id,omitempty" yaml:"id,omitempty"`
	Type      string                `gorm:"column:type;type:varchar(10);not null;default:1;comment:用户类型" json:"type,omitempty" yaml:"type,omitempty"`
	Passwd    string                `gorm:"column:passwd;type:varchar(64);not null;comment:密码" json:"passwd,omitempty" yaml:"passwd,omitempty"`
	Name      string                `gorm:"column:name;type:varchar(50);not null;comment:姓名" json:"name,omitempty" yaml:"name,omitempty"`
	Avatar    string                `gorm:"column:avatar;type:varchar(104);not null;comment:头像" json:"avatar,omitempty" yaml:"avatar,omitempty"`
	Nickname  string                `gorm:"column:nickname;type:varchar(50);not null;comment:昵称" json:"nickname,omitempty" yaml:"nickname,omitempty"`
	Mobile    string                `gorm:"column:mobile;type:varchar(50);uniqueIndex:user_mobile_email;not null;comment:手机号" json:"mobile,omitempty" yaml:"mobile,omitempty"`
	Email     string                `gorm:"column:email;type:varchar(50);uniqueIndex:user_mobile_email;comment:邮箱" json:"email,omitempty" yaml:"email,omitempty"`
	Switch    string                `gorm:"column:switch;type:char(1);not null;default:1;comment:开关" json:"switch,omitempty" yaml:"switch,omitempty"`
	CreatedAt time.Time             `gorm:"column:created_at;type:datetime;not null;comment:创建时间" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt time.Time             `gorm:"column:updated_at;type:datetime;not null;comment:更新时间" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint;not null;default:0;comment:软删标识符" json:"deleted_at,omitempty" yaml:"deleted_at,omitempty"`
}

func (m *User) TableName() string {
	return "user"
}

// ----- repo definition -----

type UserRepo struct {
	db *gorm.DB
	user
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db:   db,
		user: newUser(db),
	}
}

func NewTableUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db:   db,
		user: *newUser(db).Table("user"),
	}
}

// ----- gen gorm -----
type user struct {
	userDo    userDo
	ALL       field.Asterisk
	Id        field.Uint64
	Type      field.String
	Passwd    field.String
	Name      field.String
	Avatar    field.String
	Nickname  field.String
	Mobile    field.String
	Email     field.String
	Switch    field.String
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func newUser(db *gorm.DB, opts ...gen.DOOption) user {
	_user := user{}

	_user.userDo.UseDB(db, opts...)
	_user.userDo.UseModel(User{})

	tableName := _user.userDo.TableName()
	_user.Id = field.NewUint64(tableName, "id")
	_user.Type = field.NewString(tableName, "type")
	_user.Passwd = field.NewString(tableName, "passwd")
	_user.Name = field.NewString(tableName, "name")
	_user.Avatar = field.NewString(tableName, "avatar")
	_user.Nickname = field.NewString(tableName, "nickname")
	_user.Mobile = field.NewString(tableName, "mobile")
	_user.Email = field.NewString(tableName, "email")
	_user.Switch = field.NewString(tableName, "switch")
	_user.CreatedAt = field.NewTime(tableName, "created_at")
	_user.UpdatedAt = field.NewTime(tableName, "updated_at")
	_user.DeletedAt = field.NewField(tableName, "deleted_at")
	_user.fillFieldMap()

	return _user
}

func (c user) Table(newTableName string) *user {
	c.userDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c user) As(alias string) *user {
	c.userDo.DO = *(c.userDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *user) updateTableName(table string) *user {
	c.ALL = field.NewAsterisk(table)
	c.Id = field.NewUint64(table, "id")
	c.Type = field.NewString(table, "type")
	c.Passwd = field.NewString(table, "passwd")
	c.Name = field.NewString(table, "name")
	c.Avatar = field.NewString(table, "avatar")
	c.Nickname = field.NewString(table, "nickname")
	c.Mobile = field.NewString(table, "mobile")
	c.Email = field.NewString(table, "email")
	c.Switch = field.NewString(table, "switch")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.fillFieldMap()
	return c
}

func (c *user) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 12)
	c.fieldMap["id"] = c.Id
	c.fieldMap["type"] = c.Type
	c.fieldMap["passwd"] = c.Passwd
	c.fieldMap["name"] = c.Name
	c.fieldMap["avatar"] = c.Avatar
	c.fieldMap["nickname"] = c.Nickname
	c.fieldMap["mobile"] = c.Mobile
	c.fieldMap["email"] = c.Email
	c.fieldMap["switch"] = c.Switch
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
}

func (c *user) WithContext(ctx context.Context) *userDo { return c.userDo.WithContext(ctx) }

func (c *user) CallBackWithContext(ctx context.Context, call func(context.Context, gen.Dao) gen.Dao) *userDo {
	return c.userDo.withDO(call(ctx, &c.userDo.WithContext(ctx).DO))
}

func (c user) TableName() string { return c.userDo.TableName() }

func (c user) Alias() string { return c.userDo.Alias() }

func (c user) Columns(cols ...field.Expr) gen.Columns { return c.userDo.Columns(cols...) }

func (c *user) GetFieldsByName(fieldName []string) []field.OrderExpr {
	_f := []field.OrderExpr{}
	for _, v := range fieldName {
		_rf, ok := c.GetFieldByName(v)
		if ok {
			_f = append(_f, _rf)
		}
	}
	return _f
}

func (c *user) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c user) clone(db *gorm.DB) user {
	c.userDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c user) replaceDB(db *gorm.DB) user {
	c.userDo.ReplaceDB(db)
	return c
}

// ----- DO -----
type userDo struct{ gen.DO }

func (c userDo) Debug() *userDo {
	return c.withDO(c.DO.Debug())
}

func (c userDo) WithContext(ctx context.Context) *userDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c userDo) ReadDB() *userDo {
	return c.Clauses(dbresolver.Read)
}

func (c userDo) WriteDB() *userDo {
	return c.Clauses(dbresolver.Write)
}

func (c userDo) Session(config *gorm.Session) *userDo {
	return c.withDO(c.DO.Session(config))
}

func (c userDo) Clauses(conds ...clause.Expression) *userDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c userDo) Returning(value interface{}, columns ...string) *userDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c userDo) Not(conds ...gen.Condition) *userDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c userDo) Or(conds ...gen.Condition) *userDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c userDo) Select(conds ...field.Expr) *userDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c userDo) Where(conds ...gen.Condition) *userDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c userDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *userDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c userDo) Order(conds ...field.Expr) *userDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c userDo) Distinct(cols ...field.Expr) *userDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c userDo) Omit(cols ...field.Expr) *userDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c userDo) Join(table schema.Tabler, on ...field.Expr) *userDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c userDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c userDo) RightJoin(table schema.Tabler, on ...field.Expr) *userDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c userDo) Group(cols ...field.Expr) *userDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c userDo) Having(conds ...gen.Condition) *userDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c userDo) Limit(limit int) *userDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c userDo) Offset(offset int) *userDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c userDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c userDo) Unscoped() *userDo {
	return c.withDO(c.DO.Unscoped())
}

func (c userDo) Create(values ...*User) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c userDo) CreateInBatches(values []*User, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c userDo) Save(values ...*User) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c userDo) First() (*User, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*User), nil
	}
}

func (c userDo) Take() (*User, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*User), nil
	}
}

func (c userDo) Last() (*User, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*User), nil
	}
}

func (c userDo) Find() ([]*User, error) {
	result, err := c.DO.Find()
	return result.([]*User), err
}

func (c userDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*User, err error) {
	buf := make([]*User, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c userDo) FindInBatches(result *[]*User, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c userDo) Attrs(attrs ...field.AssignExpr) *userDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c userDo) Assign(attrs ...field.AssignExpr) *userDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c userDo) Joins(fields ...field.RelationField) *userDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c userDo) Preload(fields ...field.RelationField) *userDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c userDo) FirstOrInit() (*User, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*User), nil
	}
}

func (c userDo) FirstOrCreate() (*User, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*User), nil
	}
}

func (c userDo) FindByPage(offset int, limit int) (result []*User, count int64, err error) {
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

func (c userDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c userDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c userDo) Delete(models ...*User) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *userDo) withDO(do gen.Dao) *userDo {
	c.DO = *do.(*gen.DO)
	return c
}
