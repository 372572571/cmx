{{- $model:=. -}} 
{{- $repo:=.UpTableName -}}
{{- $repoName:=printf "%sRepo" .UpTableName }}
{{- $fields:=sortColumn $model -}}
{{- $genstructName:= .FirstLowerTableName -}}
{{- $createrModel := newModel .UpTableName false -}}
{{- $refModel:= model .UpTableName true -}}
{{- $normalModel:= model .UpTableName false -}}
package {{getPkgName}}

import(
    {{- range $k,$v:=getImportPaths}}
    "{{$v}}"
    {{- end}}
)

{{- range $k,$v:=getDefaultRefs}}
{{- if $v}}
{{$v}}
{{- end}}
{{- end}}

// ----- oneof definition -----
{{/* {{getOneofDefinition .}} */}}

// ----- model definition  -----
type {{$model.UpTableName}} struct {
    {{- range $k,$v:=sortColumn $model}}
    {{$v.UpFieldName}} {{$v.GoSchema.Type}} `{{$v.GoSchema.Tag}}`
    {{- end}}
}

func (m *{{$model.UpTableName}}) TableName() string {
	return "{{$model.TableName}}"
}

// ----- repo definition ----- 

type {{$repoName}} struct{
	db *gorm.DB
	{{$genstructName}}
}

func New{{$repoName}}(db *gorm.DB) *{{$repoName}} {
	return &{{$repoName}}{
		db:      db,
		{{$genstructName}}: new{{$repo}}(db),
	}
}

func NewTable{{$repoName}}(db *gorm.DB) *{{$repoName}} {
	return &{{$repoName}}{
		db:      db,
		{{$genstructName}}: *new{{$repo}}(db).Table("{{$model.TableName}}"),
	}
}

{{/* {{getModelToProto $repoName $normalModel $model}} */}}

// ----- gen gorm -----
type {{$genstructName}} struct {
	{{$genstructName}}Do {{$genstructName}}Do	
	ALL       field.Asterisk
	{{- range $index,$value := $fields}}	
	{{$value.UpFieldName}} field.{{genreField $value.GoSchema.Type}}
	{{- end}}	
	
	fieldMap map[string]field.Expr
}

func new{{$repo}}(db *gorm.DB, opts ...gen.DOOption) {{$genstructName}} {
	_{{$genstructName}} := {{$genstructName}}{}

	_{{$genstructName}}.{{$genstructName}}Do.UseDB(db, opts...)
	_{{$genstructName}}.{{$genstructName}}Do.UseModel({{$createrModel}})

	tableName := _{{$genstructName}}.{{$genstructName}}Do.TableName()
	{{- range $index,$value := $fields}}	
	_{{$genstructName}}.{{$value.UpFieldName}} = field.New{{genreField $value.GoSchema.Type}}(tableName, "{{$value.FieldName}}")
	{{- end}}
	_{{$genstructName}}.fillFieldMap()

	return _{{$genstructName}}
}

func (c {{$genstructName}}) Table(newTableName string) *{{$genstructName}} {
	c.{{$genstructName}}Do.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c {{$genstructName}}) As(alias string) *{{$genstructName}} {
	c.{{$genstructName}}Do.DO = *(c.{{$genstructName}}Do.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *{{$genstructName}}) updateTableName(table string) *{{$genstructName}} {
	c.ALL = field.NewAsterisk(table)
	{{- range $index,$value := $fields}}	
	c.{{$value.UpFieldName}} = field.New{{genreField $value.GoSchema.Type}}(table, "{{$value.FieldName}}")
	{{- end}}
	c.fillFieldMap()
	return c
}

func (c *{{$genstructName}}) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, {{len $fields}})
	{{- range $index,$value := $fields}}	
	c.fieldMap["{{$value.FieldName}}"] = c.{{$value.UpFieldName}}
	{{- end}}
}

func (c *{{$genstructName}}) WithContext(ctx context.Context) *{{$genstructName}}Do { return c.{{$genstructName}}Do.WithContext(ctx) }

func (c *{{$genstructName}}) CallBackWithContext(ctx context.Context, call func(context.Context, gen.Dao) gen.Dao) *{{$genstructName}}Do {
	return c.{{$genstructName}}Do.withDO(call(ctx, &c.{{$genstructName}}Do.WithContext(ctx).DO))
}

func (c {{$genstructName}}) TableName() string { return c.{{$genstructName}}Do.TableName() }

func (c {{$genstructName}}) Alias() string { return c.{{$genstructName}}Do.Alias() }

func (c {{$genstructName}}) Columns(cols ...field.Expr) gen.Columns { return c.{{$genstructName}}Do.Columns(cols...) }

func (c *{{$genstructName}}) GetFieldsByName(fieldName []string) ([]field.OrderExpr) {
	_f := []field.OrderExpr{}
	for _, v := range fieldName {
		_rf,ok:=c.GetFieldByName(v)
		if ok {
			_f = append(_f, _rf)
		}
	}
	return _f
}

func (c *{{$genstructName}}) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

{{/* func (c *{{$genstructName}}) DefaultFields() []field.Expr  {
	fieldExprs := []field.Expr{}
	for _, v := range c.fieldMap {
		fieldExprs = append(fieldExprs, v)
	}
	return fieldExprs
} */}}

func (c {{$genstructName}}) clone(db *gorm.DB) {{$genstructName}} {
	c.{{$genstructName}}Do.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c {{$genstructName}}) replaceDB(db *gorm.DB) {{$genstructName}} {
	c.{{$genstructName}}Do.ReplaceDB(db)
	return c
}

// ----- DO -----
type {{$genstructName}}Do struct{ gen.DO }

func (c {{$genstructName}}Do) Debug() *{{$genstructName}}Do {
	return c.withDO(c.DO.Debug())
}

func (c {{$genstructName}}Do) WithContext(ctx context.Context) *{{$genstructName}}Do {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c {{$genstructName}}Do) ReadDB() *{{$genstructName}}Do {
	return c.Clauses(dbresolver.Read)
}

func (c {{$genstructName}}Do) WriteDB() *{{$genstructName}}Do {
	return c.Clauses(dbresolver.Write)
}

func (c {{$genstructName}}Do) Session(config *gorm.Session) *{{$genstructName}}Do {
	return c.withDO(c.DO.Session(config))
}

func (c {{$genstructName}}Do) Clauses(conds ...clause.Expression) *{{$genstructName}}Do {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c {{$genstructName}}Do) Returning(value interface{}, columns ...string) *{{$genstructName}}Do {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c {{$genstructName}}Do) Not(conds ...gen.Condition) *{{$genstructName}}Do {
	return c.withDO(c.DO.Not(conds...))
}

func (c {{$genstructName}}Do) Or(conds ...gen.Condition) *{{$genstructName}}Do {
	return c.withDO(c.DO.Or(conds...))
}

func (c {{$genstructName}}Do) Select(conds ...field.Expr) *{{$genstructName}}Do {
	return c.withDO(c.DO.Select(conds...))
}

func (c {{$genstructName}}Do) Where(conds ...gen.Condition) *{{$genstructName}}Do {
	return c.withDO(c.DO.Where(conds...))
}

func (c {{$genstructName}}Do) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *{{$genstructName}}Do {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c {{$genstructName}}Do) Order(conds ...field.Expr) *{{$genstructName}}Do {
	return c.withDO(c.DO.Order(conds...))
}

func (c {{$genstructName}}Do) Distinct(cols ...field.Expr) *{{$genstructName}}Do {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c {{$genstructName}}Do) Omit(cols ...field.Expr) *{{$genstructName}}Do {
	return c.withDO(c.DO.Omit(cols...))
}

func (c {{$genstructName}}Do) Join(table schema.Tabler, on ...field.Expr) *{{$genstructName}}Do {
	return c.withDO(c.DO.Join(table, on...))
}

func (c {{$genstructName}}Do) LeftJoin(table schema.Tabler, on ...field.Expr) *{{$genstructName}}Do {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c {{$genstructName}}Do) RightJoin(table schema.Tabler, on ...field.Expr) *{{$genstructName}}Do {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c {{$genstructName}}Do) Group(cols ...field.Expr) *{{$genstructName}}Do {
	return c.withDO(c.DO.Group(cols...))
}

func (c {{$genstructName}}Do) Having(conds ...gen.Condition) *{{$genstructName}}Do {
	return c.withDO(c.DO.Having(conds...))
}

func (c {{$genstructName}}Do) Limit(limit int) *{{$genstructName}}Do {
	return c.withDO(c.DO.Limit(limit))
}

func (c {{$genstructName}}Do) Offset(offset int) *{{$genstructName}}Do {
	return c.withDO(c.DO.Offset(offset))
}

func (c {{$genstructName}}Do) Scopes(funcs ...func(gen.Dao) gen.Dao) *{{$genstructName}}Do {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c {{$genstructName}}Do) Unscoped() *{{$genstructName}}Do {
	return c.withDO(c.DO.Unscoped())
}

func (c {{$genstructName}}Do) Create(values ...{{$refModel}}) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c {{$genstructName}}Do) CreateInBatches(values []{{$refModel}}, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c {{$genstructName}}Do) Save(values ...{{$refModel}}) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c {{$genstructName}}Do) First() ({{$refModel}}, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.({{$refModel}}), nil
	}
}

func (c {{$genstructName}}Do) Take() ({{$refModel}}, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.({{$refModel}}), nil
	}
}

func (c {{$genstructName}}Do) Last() ({{$refModel}}, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.({{$refModel}}), nil
	}
}

func (c {{$genstructName}}Do) Find() ([]{{$refModel}}, error) {
	result, err := c.DO.Find()
	return result.([]{{$refModel}}), err
}

func (c {{$genstructName}}Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []{{$refModel}}, err error) {
	buf := make([]{{$refModel}}, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c {{$genstructName}}Do) FindInBatches(result *[]{{$refModel}}, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c {{$genstructName}}Do) Attrs(attrs ...field.AssignExpr) *{{$genstructName}}Do {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c {{$genstructName}}Do) Assign(attrs ...field.AssignExpr) *{{$genstructName}}Do {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c {{$genstructName}}Do) Joins(fields ...field.RelationField) *{{$genstructName}}Do {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c {{$genstructName}}Do) Preload(fields ...field.RelationField) *{{$genstructName}}Do {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c {{$genstructName}}Do) FirstOrInit() ({{$refModel}}, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.({{$refModel}}), nil
	}
}

func (c {{$genstructName}}Do) FirstOrCreate() ({{$refModel}}, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.({{$refModel}}), nil
	}
}

func (c {{$genstructName}}Do) FindByPage(offset int, limit int) (result []{{$refModel}}, count int64, err error) {
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

func (c {{$genstructName}}Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c {{$genstructName}}Do) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c {{$genstructName}}Do) Delete(models ...{{$refModel}}) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}
	
func (c *{{$genstructName}}Do) withDO(do gen.Dao) *{{$genstructName}}Do {
	c.DO = *do.(*gen.DO)
	return c
}
