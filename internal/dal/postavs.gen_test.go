// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"fmt"
	"testing"

	"github.com/kuzgoga/nto-boilerplate/internal/models"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := _gen_test_db.AutoMigrate(&models.Postav{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&models.Postav{}) fail: %s", err)
	}
}

func Test_postavQuery(t *testing.T) {
	postav := newPostav(_gen_test_db)
	postav = *postav.As(postav.TableName())
	_do := postav.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(postav.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <postavs> fail:", err)
		return
	}

	_, ok := postav.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from postav success")
	}

	err = _do.Create(&models.Postav{})
	if err != nil {
		t.Error("create item in table <postavs> fail:", err)
	}

	err = _do.Save(&models.Postav{})
	if err != nil {
		t.Error("create item in table <postavs> fail:", err)
	}

	err = _do.CreateInBatches([]*models.Postav{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <postavs> fail:", err)
	}

	_, err = _do.Select(postav.ALL).Take()
	if err != nil {
		t.Error("Take() on table <postavs> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <postavs> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <postavs> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <postavs> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*models.Postav{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <postavs> fail:", err)
	}

	_, err = _do.Select(postav.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <postavs> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <postavs> fail:", err)
	}

	_, err = _do.Select(postav.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <postavs> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <postavs> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <postavs> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <postavs> fail:", err)
	}

	_, err = _do.ScanByPage(&models.Postav{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <postavs> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <postavs> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <postavs> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <postavs> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <postavs> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <postavs> fail:", err)
	}
}
