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

func newReceiver(db *gorm.DB, opts ...gen.DOOption) receiver {
	_receiver := receiver{}

	_receiver.receiverDo.UseDB(db, opts...)
	_receiver.receiverDo.UseModel(&models.Receiver{})

	tableName := _receiver.receiverDo.TableName()
	_receiver.ALL = field.NewAsterisk(tableName)
	_receiver.Id = field.NewUint(tableName, "id")
	_receiver.ExporterId = field.NewUint(tableName, "exporter_id")
	_receiver.ExporterQuantity = field.NewInt(tableName, "exporter_quantity")
	_receiver.FactoryQuantity = field.NewInt(tableName, "factory_quantity")
	_receiver.Description = field.NewString(tableName, "description")
	_receiver.SpecId = field.NewUint(tableName, "spec_id")
	_receiver.CreatedAt = field.NewInt64(tableName, "created_at")
	_receiver.Exporter = receiverBelongsToExporter{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Exporter", "models.Exporter"),
		Receivers: struct {
			field.RelationField
			Exporter struct {
				field.RelationField
			}
			Spec struct {
				field.RelationField
				Receivers struct {
					field.RelationField
				}
			}
		}{
			RelationField: field.NewRelation("Exporter.Receivers", "models.Receiver"),
			Exporter: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Exporter.Receivers.Exporter", "models.Exporter"),
			},
			Spec: struct {
				field.RelationField
				Receivers struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Exporter.Receivers.Spec", "models.WoodSpecType"),
				Receivers: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Exporter.Receivers.Spec.Receivers", "models.Receiver"),
				},
			},
		},
	}

	_receiver.Spec = receiverBelongsToSpec{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Spec", "models.WoodSpecType"),
	}

	_receiver.fillFieldMap()

	return _receiver
}

type receiver struct {
	receiverDo

	ALL              field.Asterisk
	Id               field.Uint
	ExporterId       field.Uint
	ExporterQuantity field.Int
	FactoryQuantity  field.Int
	Description      field.String
	SpecId           field.Uint
	CreatedAt        field.Int64
	Exporter         receiverBelongsToExporter

	Spec receiverBelongsToSpec

	fieldMap map[string]field.Expr
}

func (r receiver) Table(newTableName string) *receiver {
	r.receiverDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r receiver) As(alias string) *receiver {
	r.receiverDo.DO = *(r.receiverDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *receiver) updateTableName(table string) *receiver {
	r.ALL = field.NewAsterisk(table)
	r.Id = field.NewUint(table, "id")
	r.ExporterId = field.NewUint(table, "exporter_id")
	r.ExporterQuantity = field.NewInt(table, "exporter_quantity")
	r.FactoryQuantity = field.NewInt(table, "factory_quantity")
	r.Description = field.NewString(table, "description")
	r.SpecId = field.NewUint(table, "spec_id")
	r.CreatedAt = field.NewInt64(table, "created_at")

	r.fillFieldMap()

	return r
}

func (r *receiver) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *receiver) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 9)
	r.fieldMap["id"] = r.Id
	r.fieldMap["exporter_id"] = r.ExporterId
	r.fieldMap["exporter_quantity"] = r.ExporterQuantity
	r.fieldMap["factory_quantity"] = r.FactoryQuantity
	r.fieldMap["description"] = r.Description
	r.fieldMap["spec_id"] = r.SpecId
	r.fieldMap["created_at"] = r.CreatedAt

}

func (r receiver) clone(db *gorm.DB) receiver {
	r.receiverDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r receiver) replaceDB(db *gorm.DB) receiver {
	r.receiverDo.ReplaceDB(db)
	return r
}

type receiverBelongsToExporter struct {
	db *gorm.DB

	field.RelationField

	Receivers struct {
		field.RelationField
		Exporter struct {
			field.RelationField
		}
		Spec struct {
			field.RelationField
			Receivers struct {
				field.RelationField
			}
		}
	}
}

func (a receiverBelongsToExporter) Where(conds ...field.Expr) *receiverBelongsToExporter {
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

func (a receiverBelongsToExporter) WithContext(ctx context.Context) *receiverBelongsToExporter {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a receiverBelongsToExporter) Session(session *gorm.Session) *receiverBelongsToExporter {
	a.db = a.db.Session(session)
	return &a
}

func (a receiverBelongsToExporter) Model(m *models.Receiver) *receiverBelongsToExporterTx {
	return &receiverBelongsToExporterTx{a.db.Model(m).Association(a.Name())}
}

type receiverBelongsToExporterTx struct{ tx *gorm.Association }

func (a receiverBelongsToExporterTx) Find() (result *models.Exporter, err error) {
	return result, a.tx.Find(&result)
}

func (a receiverBelongsToExporterTx) Append(values ...*models.Exporter) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a receiverBelongsToExporterTx) Replace(values ...*models.Exporter) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a receiverBelongsToExporterTx) Delete(values ...*models.Exporter) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a receiverBelongsToExporterTx) Clear() error {
	return a.tx.Clear()
}

func (a receiverBelongsToExporterTx) Count() int64 {
	return a.tx.Count()
}

type receiverBelongsToSpec struct {
	db *gorm.DB

	field.RelationField
}

func (a receiverBelongsToSpec) Where(conds ...field.Expr) *receiverBelongsToSpec {
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

func (a receiverBelongsToSpec) WithContext(ctx context.Context) *receiverBelongsToSpec {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a receiverBelongsToSpec) Session(session *gorm.Session) *receiverBelongsToSpec {
	a.db = a.db.Session(session)
	return &a
}

func (a receiverBelongsToSpec) Model(m *models.Receiver) *receiverBelongsToSpecTx {
	return &receiverBelongsToSpecTx{a.db.Model(m).Association(a.Name())}
}

type receiverBelongsToSpecTx struct{ tx *gorm.Association }

func (a receiverBelongsToSpecTx) Find() (result *models.WoodSpecType, err error) {
	return result, a.tx.Find(&result)
}

func (a receiverBelongsToSpecTx) Append(values ...*models.WoodSpecType) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a receiverBelongsToSpecTx) Replace(values ...*models.WoodSpecType) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a receiverBelongsToSpecTx) Delete(values ...*models.WoodSpecType) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a receiverBelongsToSpecTx) Clear() error {
	return a.tx.Clear()
}

func (a receiverBelongsToSpecTx) Count() int64 {
	return a.tx.Count()
}

type receiverDo struct{ gen.DO }

type IReceiverDo interface {
	gen.SubQuery
	Debug() IReceiverDo
	WithContext(ctx context.Context) IReceiverDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IReceiverDo
	WriteDB() IReceiverDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IReceiverDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IReceiverDo
	Not(conds ...gen.Condition) IReceiverDo
	Or(conds ...gen.Condition) IReceiverDo
	Select(conds ...field.Expr) IReceiverDo
	Where(conds ...gen.Condition) IReceiverDo
	Order(conds ...field.Expr) IReceiverDo
	Distinct(cols ...field.Expr) IReceiverDo
	Omit(cols ...field.Expr) IReceiverDo
	Join(table schema.Tabler, on ...field.Expr) IReceiverDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IReceiverDo
	RightJoin(table schema.Tabler, on ...field.Expr) IReceiverDo
	Group(cols ...field.Expr) IReceiverDo
	Having(conds ...gen.Condition) IReceiverDo
	Limit(limit int) IReceiverDo
	Offset(offset int) IReceiverDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IReceiverDo
	Unscoped() IReceiverDo
	Create(values ...*models.Receiver) error
	CreateInBatches(values []*models.Receiver, batchSize int) error
	Save(values ...*models.Receiver) error
	First() (*models.Receiver, error)
	Take() (*models.Receiver, error)
	Last() (*models.Receiver, error)
	Find() ([]*models.Receiver, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Receiver, err error)
	FindInBatches(result *[]*models.Receiver, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Receiver) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IReceiverDo
	Assign(attrs ...field.AssignExpr) IReceiverDo
	Joins(fields ...field.RelationField) IReceiverDo
	Preload(fields ...field.RelationField) IReceiverDo
	FirstOrInit() (*models.Receiver, error)
	FirstOrCreate() (*models.Receiver, error)
	FindByPage(offset int, limit int) (result []*models.Receiver, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IReceiverDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r receiverDo) Debug() IReceiverDo {
	return r.withDO(r.DO.Debug())
}

func (r receiverDo) WithContext(ctx context.Context) IReceiverDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r receiverDo) ReadDB() IReceiverDo {
	return r.Clauses(dbresolver.Read)
}

func (r receiverDo) WriteDB() IReceiverDo {
	return r.Clauses(dbresolver.Write)
}

func (r receiverDo) Session(config *gorm.Session) IReceiverDo {
	return r.withDO(r.DO.Session(config))
}

func (r receiverDo) Clauses(conds ...clause.Expression) IReceiverDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r receiverDo) Returning(value interface{}, columns ...string) IReceiverDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r receiverDo) Not(conds ...gen.Condition) IReceiverDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r receiverDo) Or(conds ...gen.Condition) IReceiverDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r receiverDo) Select(conds ...field.Expr) IReceiverDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r receiverDo) Where(conds ...gen.Condition) IReceiverDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r receiverDo) Order(conds ...field.Expr) IReceiverDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r receiverDo) Distinct(cols ...field.Expr) IReceiverDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r receiverDo) Omit(cols ...field.Expr) IReceiverDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r receiverDo) Join(table schema.Tabler, on ...field.Expr) IReceiverDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r receiverDo) LeftJoin(table schema.Tabler, on ...field.Expr) IReceiverDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r receiverDo) RightJoin(table schema.Tabler, on ...field.Expr) IReceiverDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r receiverDo) Group(cols ...field.Expr) IReceiverDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r receiverDo) Having(conds ...gen.Condition) IReceiverDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r receiverDo) Limit(limit int) IReceiverDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r receiverDo) Offset(offset int) IReceiverDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r receiverDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IReceiverDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r receiverDo) Unscoped() IReceiverDo {
	return r.withDO(r.DO.Unscoped())
}

func (r receiverDo) Create(values ...*models.Receiver) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r receiverDo) CreateInBatches(values []*models.Receiver, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r receiverDo) Save(values ...*models.Receiver) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r receiverDo) First() (*models.Receiver, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Receiver), nil
	}
}

func (r receiverDo) Take() (*models.Receiver, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Receiver), nil
	}
}

func (r receiverDo) Last() (*models.Receiver, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Receiver), nil
	}
}

func (r receiverDo) Find() ([]*models.Receiver, error) {
	result, err := r.DO.Find()
	return result.([]*models.Receiver), err
}

func (r receiverDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Receiver, err error) {
	buf := make([]*models.Receiver, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r receiverDo) FindInBatches(result *[]*models.Receiver, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r receiverDo) Attrs(attrs ...field.AssignExpr) IReceiverDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r receiverDo) Assign(attrs ...field.AssignExpr) IReceiverDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r receiverDo) Joins(fields ...field.RelationField) IReceiverDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r receiverDo) Preload(fields ...field.RelationField) IReceiverDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r receiverDo) FirstOrInit() (*models.Receiver, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Receiver), nil
	}
}

func (r receiverDo) FirstOrCreate() (*models.Receiver, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Receiver), nil
	}
}

func (r receiverDo) FindByPage(offset int, limit int) (result []*models.Receiver, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r receiverDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r receiverDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r receiverDo) Delete(models ...*models.Receiver) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *receiverDo) withDO(do gen.Dao) *receiverDo {
	r.DO = *do.(*gen.DO)
	return r
}
