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
type User struct {
	Id                   uint64                `gorm:"column:id;type:bigint(20) unsigned;primaryKey;not null;comment:用户id" json:"id,omitempty" yaml:"id,omitempty"`
	Status               string                `gorm:"column:status;type:char(4);not null;comment:状态,[0:未指定,1:启用,2:禁用,3:临时锁定]" json:"status,omitempty" yaml:"status,omitempty"`
	Type                 string                `gorm:"column:type;type:varchar(10);index:idx_user_type;not null;comment:用户类型,[0:未指定,1:个人账号,2:专业用户,3:众包用户,4:个体户]" json:"type,omitempty" yaml:"type,omitempty"`
	LuoboId              string                `gorm:"column:luobo_id;type:varchar(50);not null;comment:萝卜号" json:"luobo_id,omitempty" yaml:"luobo_id,omitempty"`
	Passwd               string                `gorm:"column:passwd;type:varchar(64);not null;comment:密码" json:"passwd,omitempty" yaml:"passwd,omitempty"`
	Name                 string                `gorm:"column:name;type:varchar(50);not null;comment:姓名" json:"name,omitempty" yaml:"name,omitempty"`
	Nickname             string                `gorm:"column:nickname;type:varchar(50);not null;comment:昵称" json:"nickname,omitempty" yaml:"nickname,omitempty"`
	Mobile               string                `gorm:"column:mobile;type:varchar(50);index:user_mobile;not null;comment:手机号" json:"mobile,omitempty" yaml:"mobile,omitempty"`
	Email                string                `gorm:"column:email;type:varchar(50);not null;comment:邮箱" json:"email,omitempty" yaml:"email,omitempty"`
	Avatar               string                `gorm:"column:avatar;type:varchar(256);not null;comment:头像" json:"avatar,omitempty" yaml:"avatar,omitempty"`
	Sex                  int                   `gorm:"column:sex;type:int(10) unsigned;not null;comment:性别,[0:未指定,1:男,2:女,3:未知]" json:"sex,omitempty" yaml:"sex,omitempty"`
	RegisterIp           string                `gorm:"column:register_ip;type:varchar(50);not null;comment:注册IP" json:"register_ip,omitempty" yaml:"register_ip,omitempty"`
	IsAuth               string                `gorm:"column:is_auth;type:char(1);not null;comment:是否已认证" json:"is_auth,omitempty" yaml:"is_auth,omitempty"`
	IsCrowdsourcingAgent int                   `gorm:"column:is_crowdsourcing_agent;type:tinyint(1);not null;comment:是否是众包主体(仅专业用户可设置)" json:"is_crowdsourcing_agent,omitempty" yaml:"is_crowdsourcing_agent,omitempty"`
	PayPasswd            string                `gorm:"column:pay_passwd;type:varchar(64);not null;comment:交易密码" json:"pay_passwd,omitempty" yaml:"pay_passwd,omitempty"`
	Region               string                `gorm:"column:region;type:varchar(32);not null;comment:地区" json:"region,omitempty" yaml:"region,omitempty"`
	Birthday             int64                 `gorm:"column:birthday;type:bigint(20);not null;comment:用户生日" json:"birthday,omitempty" yaml:"birthday,omitempty"`
	LastLoginAt          int64                 `gorm:"column:last_login_at;type:bigint(20);not null;comment:最后登陆时间" json:"last_login_at,omitempty" yaml:"last_login_at,omitempty"`
	CreatedAt            time.Time             `gorm:"column:created_at;type:datetime;not null" json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt            time.Time             `gorm:"column:updated_at;type:datetime;not null" json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	DeletedAt            soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint(20);not null" json:"deleted_at,omitempty" yaml:"deleted_at,omitempty"`
	Score                int64                 `gorm:"column:score;type:bigint(20);not null;comment:用户总评分" json:"score,omitempty" yaml:"score,omitempty"`
	Count                int64                 `gorm:"column:count;type:bigint(20);not null;comment:订单计数" json:"count,omitempty" yaml:"count,omitempty"`
	DeviceType           int                   `gorm:"column:device_type;type:int(10) unsigned;not null;comment:设备类型,[0:未指定,1:Android,2:IOS,3:H5]" json:"device_type,omitempty" yaml:"device_type,omitempty"`
	IsOldPasswd          int                   `gorm:"column:is_old_passwd;type:tinyint(1);not null;comment:是否是旧登陆密码" json:"is_old_passwd,omitempty" yaml:"is_old_passwd,omitempty"`
	IsOldPayPasswd       int                   `gorm:"column:is_old_pay_passwd;type:tinyint(1);not null;comment:是否是旧支付密码" json:"is_old_pay_passwd,omitempty" yaml:"is_old_pay_passwd,omitempty"`
	NewTranNetMemberCode string                `gorm:"column:new_tran_net_member_code;type:varchar(128);not null;comment:平安银行新开户时,用的会员代码(新用户与uid一致,仅注销钱包时会生成一个新的)" json:"new_tran_net_member_code,omitempty" yaml:"new_tran_net_member_code,omitempty"`
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

// not enable model to proto

// ----- gen gorm -----
type user struct {
	userDo               userDo
	ALL                  field.Asterisk
	Id                   field.Uint64
	Status               field.String
	Type                 field.String
	LuoboId              field.String
	Passwd               field.String
	Name                 field.String
	Nickname             field.String
	Mobile               field.String
	Email                field.String
	Avatar               field.String
	Sex                  field.Int
	RegisterIp           field.String
	IsAuth               field.String
	IsCrowdsourcingAgent field.Int
	PayPasswd            field.String
	Region               field.String
	Birthday             field.Int64
	LastLoginAt          field.Int64
	CreatedAt            field.Time
	UpdatedAt            field.Time
	DeletedAt            field.Field
	Score                field.Int64
	Count                field.Int64
	DeviceType           field.Int
	IsOldPasswd          field.Int
	IsOldPayPasswd       field.Int
	NewTranNetMemberCode field.String

	fieldMap map[string]field.Expr
}

func newUser(db *gorm.DB, opts ...gen.DOOption) user {
	_user := user{}

	_user.userDo.UseDB(db, opts...)
	_user.userDo.UseModel(User{})

	tableName := _user.userDo.TableName()
	_user.Id = field.NewUint64(tableName, "id")
	_user.Status = field.NewString(tableName, "status")
	_user.Type = field.NewString(tableName, "type")
	_user.LuoboId = field.NewString(tableName, "luobo_id")
	_user.Passwd = field.NewString(tableName, "passwd")
	_user.Name = field.NewString(tableName, "name")
	_user.Nickname = field.NewString(tableName, "nickname")
	_user.Mobile = field.NewString(tableName, "mobile")
	_user.Email = field.NewString(tableName, "email")
	_user.Avatar = field.NewString(tableName, "avatar")
	_user.Sex = field.NewInt(tableName, "sex")
	_user.RegisterIp = field.NewString(tableName, "register_ip")
	_user.IsAuth = field.NewString(tableName, "is_auth")
	_user.IsCrowdsourcingAgent = field.NewInt(tableName, "is_crowdsourcing_agent")
	_user.PayPasswd = field.NewString(tableName, "pay_passwd")
	_user.Region = field.NewString(tableName, "region")
	_user.Birthday = field.NewInt64(tableName, "birthday")
	_user.LastLoginAt = field.NewInt64(tableName, "last_login_at")
	_user.CreatedAt = field.NewTime(tableName, "created_at")
	_user.UpdatedAt = field.NewTime(tableName, "updated_at")
	_user.DeletedAt = field.NewField(tableName, "deleted_at")
	_user.Score = field.NewInt64(tableName, "score")
	_user.Count = field.NewInt64(tableName, "count")
	_user.DeviceType = field.NewInt(tableName, "device_type")
	_user.IsOldPasswd = field.NewInt(tableName, "is_old_passwd")
	_user.IsOldPayPasswd = field.NewInt(tableName, "is_old_pay_passwd")
	_user.NewTranNetMemberCode = field.NewString(tableName, "new_tran_net_member_code")
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
	c.Status = field.NewString(table, "status")
	c.Type = field.NewString(table, "type")
	c.LuoboId = field.NewString(table, "luobo_id")
	c.Passwd = field.NewString(table, "passwd")
	c.Name = field.NewString(table, "name")
	c.Nickname = field.NewString(table, "nickname")
	c.Mobile = field.NewString(table, "mobile")
	c.Email = field.NewString(table, "email")
	c.Avatar = field.NewString(table, "avatar")
	c.Sex = field.NewInt(table, "sex")
	c.RegisterIp = field.NewString(table, "register_ip")
	c.IsAuth = field.NewString(table, "is_auth")
	c.IsCrowdsourcingAgent = field.NewInt(table, "is_crowdsourcing_agent")
	c.PayPasswd = field.NewString(table, "pay_passwd")
	c.Region = field.NewString(table, "region")
	c.Birthday = field.NewInt64(table, "birthday")
	c.LastLoginAt = field.NewInt64(table, "last_login_at")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.Score = field.NewInt64(table, "score")
	c.Count = field.NewInt64(table, "count")
	c.DeviceType = field.NewInt(table, "device_type")
	c.IsOldPasswd = field.NewInt(table, "is_old_passwd")
	c.IsOldPayPasswd = field.NewInt(table, "is_old_pay_passwd")
	c.NewTranNetMemberCode = field.NewString(table, "new_tran_net_member_code")
	c.fillFieldMap()
	return c
}

func (c *user) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 27)
	c.fieldMap["id"] = c.Id
	c.fieldMap["status"] = c.Status
	c.fieldMap["type"] = c.Type
	c.fieldMap["luobo_id"] = c.LuoboId
	c.fieldMap["passwd"] = c.Passwd
	c.fieldMap["name"] = c.Name
	c.fieldMap["nickname"] = c.Nickname
	c.fieldMap["mobile"] = c.Mobile
	c.fieldMap["email"] = c.Email
	c.fieldMap["avatar"] = c.Avatar
	c.fieldMap["sex"] = c.Sex
	c.fieldMap["register_ip"] = c.RegisterIp
	c.fieldMap["is_auth"] = c.IsAuth
	c.fieldMap["is_crowdsourcing_agent"] = c.IsCrowdsourcingAgent
	c.fieldMap["pay_passwd"] = c.PayPasswd
	c.fieldMap["region"] = c.Region
	c.fieldMap["birthday"] = c.Birthday
	c.fieldMap["last_login_at"] = c.LastLoginAt
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["score"] = c.Score
	c.fieldMap["count"] = c.Count
	c.fieldMap["device_type"] = c.DeviceType
	c.fieldMap["is_old_passwd"] = c.IsOldPasswd
	c.fieldMap["is_old_pay_passwd"] = c.IsOldPayPasswd
	c.fieldMap["new_tran_net_member_code"] = c.NewTranNetMemberCode
}

func (c *user) WithContext(ctx context.Context) *userDo { return c.userDo.WithContext(ctx) }

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
