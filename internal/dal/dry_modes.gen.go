// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"github.com/kuzgoga/nto-boilerplate/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newDryMode(db *gorm.DB, opts ...gen.DOOption) dryMode {
	_dryMode := dryMode{}

	_dryMode.dryModeDo.UseDB(db, opts...)
	_dryMode.dryModeDo.UseModel(&models.DryMode{})

	tableName := _dryMode.dryModeDo.TableName()
	_dryMode.ALL = field.NewAsterisk(tableName)
	_dryMode.Id = field.NewUint(tableName, "id")
	_dryMode.Name = field.NewString(tableName, "name")
	_dryMode.WetMaterialsId = field.NewUint(tableName, "wet_materials_id")
	_dryMode.HumidityPercent = field.NewInt(tableName, "humidity_percent")
	_dryMode.ProcentYsyshki = field.NewInt(tableName, "procent_ysyshki")
	_dryMode.WetMaterials = dryModeBelongsToWetMaterials{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("WetMaterials", "models.WoodSpec"),
		WoodSpecType: struct {
			field.RelationField
			Receivers struct {
				field.RelationField
				Exporter struct {
					field.RelationField
					Receivers struct {
						field.RelationField
					}
				}
				Spec struct {
					field.RelationField
				}
			}
		}{
			RelationField: field.NewRelation("WetMaterials.WoodSpecType", "models.WoodSpecType"),
			Receivers: struct {
				field.RelationField
				Exporter struct {
					field.RelationField
					Receivers struct {
						field.RelationField
					}
				}
				Spec struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("WetMaterials.WoodSpecType.Receivers", "models.Receiver"),
				Exporter: struct {
					field.RelationField
					Receivers struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("WetMaterials.WoodSpecType.Receivers.Exporter", "models.Exporter"),
					Receivers: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("WetMaterials.WoodSpecType.Receivers.Exporter.Receivers", "models.Receiver"),
					},
				},
				Spec: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("WetMaterials.WoodSpecType.Receivers.Spec", "models.WoodSpecType"),
				},
			},
		},
	}

	_dryMode.fillFieldMap()

	return _dryMode
}

type dryMode struct {
	dryModeDo

	ALL             field.Asterisk
	Id              field.Uint
	Name            field.String
	WetMaterialsId  field.Uint
	HumidityPercent field.Int
	ProcentYsyshki  field.Int
	WetMaterials    dryModeBelongsToWetMaterials

	fieldMap map[string]field.Expr
}

func (d dryMode) Table(newTableName string) *dryMode {
	d.dryModeDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dryMode) As(alias string) *dryMode {
	d.dryModeDo.DO = *(d.dryModeDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dryMode) updateTableName(table string) *dryMode {
	d.ALL = field.NewAsterisk(table)
	d.Id = field.NewUint(table, "id")
	d.Name = field.NewString(table, "name")
	d.WetMaterialsId = field.NewUint(table, "wet_materials_id")
	d.HumidityPercent = field.NewInt(table, "humidity_percent")
	d.ProcentYsyshki = field.NewInt(table, "procent_ysyshki")

	d.fillFieldMap()

	return d
}

func (d *dryMode) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dryMode) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 6)
	d.fieldMap["id"] = d.Id
	d.fieldMap["name"] = d.Name
	d.fieldMap["wet_materials_id"] = d.WetMaterialsId
	d.fieldMap["humidity_percent"] = d.HumidityPercent
	d.fieldMap["procent_ysyshki"] = d.ProcentYsyshki

}

func (d dryMode) clone(db *gorm.DB) dryMode {
	d.dryModeDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d dryMode) replaceDB(db *gorm.DB) dryMode {
	d.dryModeDo.ReplaceDB(db)
	return d
}

type dryModeBelongsToWetMaterials struct {
	db *gorm.DB

	field.RelationField

	WoodSpecType struct {
		field.RelationField
		Receivers struct {
			field.RelationField
			Exporter struct {
				field.RelationField
				Receivers struct {
					field.RelationField
				}
			}
			Spec struct {
				field.RelationField
			}
		}
	}
}

func (a dryModeBelongsToWetMaterials) Where(conds ...field.Expr) *dryModeBelongsToWetMaterials {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a dryModeBelongsToWetMaterials) WithContext(ctx context.Context) *dryModeBelongsToWetMaterials {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a dryModeBelongsToWetMaterials) Session(session *gorm.Session) *dryModeBelongsToWetMaterials {
	a.db = a.db.Session(session)
	return &a
}

func (a dryModeBelongsToWetMaterials) Model(m *models.DryMode) *dryModeBelongsToWetMaterialsTx {
	return &dryModeBelongsToWetMaterialsTx{a.db.Model(m).Association(a.Name())}
}

type dryModeBelongsToWetMaterialsTx struct{ tx *gorm.Association }

func (a dryModeBelongsToWetMaterialsTx) Find() (result *models.WoodSpec, err error) {
	return result, a.tx.Find(&result)
}

func (a dryModeBelongsToWetMaterialsTx) Append(values ...*models.WoodSpec) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a dryModeBelongsToWetMaterialsTx) Replace(values ...*models.WoodSpec) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a dryModeBelongsToWetMaterialsTx) Delete(values ...*models.WoodSpec) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a dryModeBelongsToWetMaterialsTx) Clear() error {
	return a.tx.Clear()
}

func (a dryModeBelongsToWetMaterialsTx) Count() int64 {
	return a.tx.Count()
}

type dryModeDo struct{ gen.DO }

type IDryModeDo interface {
	gen.SubQuery
	Debug() IDryModeDo
	WithContext(ctx context.Context) IDryModeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDryModeDo
	WriteDB() IDryModeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDryModeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDryModeDo
	Not(conds ...gen.Condition) IDryModeDo
	Or(conds ...gen.Condition) IDryModeDo
	Select(conds ...field.Expr) IDryModeDo
	Where(conds ...gen.Condition) IDryModeDo
	Order(conds ...field.Expr) IDryModeDo
	Distinct(cols ...field.Expr) IDryModeDo
	Omit(cols ...field.Expr) IDryModeDo
	Join(table schema.Tabler, on ...field.Expr) IDryModeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDryModeDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDryModeDo
	Group(cols ...field.Expr) IDryModeDo
	Having(conds ...gen.Condition) IDryModeDo
	Limit(limit int) IDryModeDo
	Offset(offset int) IDryModeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDryModeDo
	Unscoped() IDryModeDo
	Create(values ...*models.DryMode) error
	CreateInBatches(values []*models.DryMode, batchSize int) error
	Save(values ...*models.DryMode) error
	First() (*models.DryMode, error)
	Take() (*models.DryMode, error)
	Last() (*models.DryMode, error)
	Find() ([]*models.DryMode, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.DryMode, err error)
	FindInBatches(result *[]*models.DryMode, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.DryMode) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDryModeDo
	Assign(attrs ...field.AssignExpr) IDryModeDo
	Joins(fields ...field.RelationField) IDryModeDo
	Preload(fields ...field.RelationField) IDryModeDo
	FirstOrInit() (*models.DryMode, error)
	FirstOrCreate() (*models.DryMode, error)
	FindByPage(offset int, limit int) (result []*models.DryMode, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDryModeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dryModeDo) Debug() IDryModeDo {
	return d.withDO(d.DO.Debug())
}

func (d dryModeDo) WithContext(ctx context.Context) IDryModeDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dryModeDo) ReadDB() IDryModeDo {
	return d.Clauses(dbresolver.Read)
}

func (d dryModeDo) WriteDB() IDryModeDo {
	return d.Clauses(dbresolver.Write)
}

func (d dryModeDo) Session(config *gorm.Session) IDryModeDo {
	return d.withDO(d.DO.Session(config))
}

func (d dryModeDo) Clauses(conds ...clause.Expression) IDryModeDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dryModeDo) Returning(value interface{}, columns ...string) IDryModeDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dryModeDo) Not(conds ...gen.Condition) IDryModeDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dryModeDo) Or(conds ...gen.Condition) IDryModeDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dryModeDo) Select(conds ...field.Expr) IDryModeDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dryModeDo) Where(conds ...gen.Condition) IDryModeDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dryModeDo) Order(conds ...field.Expr) IDryModeDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dryModeDo) Distinct(cols ...field.Expr) IDryModeDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dryModeDo) Omit(cols ...field.Expr) IDryModeDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dryModeDo) Join(table schema.Tabler, on ...field.Expr) IDryModeDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dryModeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDryModeDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dryModeDo) RightJoin(table schema.Tabler, on ...field.Expr) IDryModeDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dryModeDo) Group(cols ...field.Expr) IDryModeDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dryModeDo) Having(conds ...gen.Condition) IDryModeDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dryModeDo) Limit(limit int) IDryModeDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dryModeDo) Offset(offset int) IDryModeDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dryModeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDryModeDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dryModeDo) Unscoped() IDryModeDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dryModeDo) Create(values ...*models.DryMode) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dryModeDo) CreateInBatches(values []*models.DryMode, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dryModeDo) Save(values ...*models.DryMode) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dryModeDo) First() (*models.DryMode, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.DryMode), nil
	}
}

func (d dryModeDo) Take() (*models.DryMode, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.DryMode), nil
	}
}

func (d dryModeDo) Last() (*models.DryMode, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.DryMode), nil
	}
}

func (d dryModeDo) Find() ([]*models.DryMode, error) {
	result, err := d.DO.Find()
	return result.([]*models.DryMode), err
}

func (d dryModeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.DryMode, err error) {
	buf := make([]*models.DryMode, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dryModeDo) FindInBatches(result *[]*models.DryMode, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dryModeDo) Attrs(attrs ...field.AssignExpr) IDryModeDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dryModeDo) Assign(attrs ...field.AssignExpr) IDryModeDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dryModeDo) Joins(fields ...field.RelationField) IDryModeDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dryModeDo) Preload(fields ...field.RelationField) IDryModeDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dryModeDo) FirstOrInit() (*models.DryMode, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.DryMode), nil
	}
}

func (d dryModeDo) FirstOrCreate() (*models.DryMode, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.DryMode), nil
	}
}

func (d dryModeDo) FindByPage(offset int, limit int) (result []*models.DryMode, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d dryModeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dryModeDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d dryModeDo) Delete(models ...*models.DryMode) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *dryModeDo) withDO(do gen.Dao) *dryModeDo {
	d.DO = *do.(*gen.DO)
	return d
}
