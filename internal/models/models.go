package models

var Entities = []any{
	&WoodSpecType{},
	&WoodSpec{},
	&Receiver{},
	&Exporter{},
	&WorkResult{},
	&Postav{},
	&SushkaResult{},
	&DryMode{},
}

type WoodSpecType struct {
	Id        uint       `gorm:"primaryKey" ui:"hidden"`
	Name      string     `ui:"label:Название"`
	Receivers []Receiver `ui:"hidden" gorm:"constraint:OnDelete:CASCADE;foreignKey:SpecId;"`
}

type WoodSpec struct {
	Id               uint         `gorm:"primaryKey" ui:"hidden"`
	WoodSpecTypeId   uint         `ui:"hidden"`
	WoodSpecType     WoodSpecType `ui:"label:Материал номеклатуры;field:Name;" gorm:"constraint:OnDelete:CASCADE"`
	Poroda           string       `ui:"label:Порода древесины"`
	Sort             string       `ui:"label:Сорт"`
	DlinaBrevna      string       `ui:"label:Длина бревна"`
	DiameterBrevna   string       `ui:"label:Диаметр бревна"`
	SechenieDoski    string       `ui:"label:Сечение доски"`
	ProcentVlajnosti int          `ui:"label:Процент влажности"`
	DiameterGranyl   string       `ui:"label:Диаметр гранул"`
}

type Exporter struct {
	Id          uint       `gorm:"primaryKey" ui:"hidden"`
	Name        string     `ui:"label:Имя"`
	Description string     `ui:"label:Описание"`
	Receivers   []Receiver `ui:"hidden" gorm:"constraint:OnDelete:CASCADE"`
}

type Receiver struct {
	Id               uint         `gorm:"primaryKey" ui:"hidden"`
	ExporterId       uint         `ui:"hidden"`
	Exporter         Exporter     `ui:"label:Поставщик;field:Name" gorm:"constraint:OnDelete:CASCADE"`
	ExporterQuantity int          `ui:"label:Количество сырья (по данным поставщика), кубометр"`
	FactoryQuantity  int          `ui:"label:Количество сырья (по данным завода), кубометр"`
	Description      string       `ui:"label:Описание"`
	Spec             WoodSpecType `ui:"label:Номеклатура;field:Name;constraint:OnDelete:CASCADE"`
	SpecId           uint         `ui:"hidden"`
	CreatedAt        int64        `gorm:"autoCreateTime" ui:"label:Дата приёмки;readonly;datatype:datetime;"`
}

type WorkResult struct {
	Id               uint         `gorm:"primaryKey" ui:"label:Номер номенклатуры"`
	MaterialId       uint         `ui:"hidden"`
	Material         WoodSpecType `ui:"label:Номенклатура материала;field:Name;" gorm:"constraint:OnDelete:CASCADE"`
	MaterialQuantity int          `ui:"label:Количество материала"`
	ProductSpecId    int          `ui:"hidden"`
	ProductSpec      WoodSpec     `ui:"label:Номенклатура продукта;field:Id;" gorm:"constraint:OnDelete:CASCADE"`
	WorkDate         int64        `ui:"label:Дата работы цеха; datatype:datetime;"`
}

type Postav struct {
	Id                  uint     `gorm:"primaryKey" ui:"hidden"`
	Name                string   `ui:"label:Наименование постава"`
	CenterDoskaSpecId   int      `ui:"hidden"`
	CenterDoskaSpec     WoodSpec `ui:"label:Номеклатура центральной доски;field:Id;" gorm:"constraint:OnDelete:CASCADE"`
	CenterOutPercent    int      `ui:"label:Процент выхода центральной доски"`
	BacksideDoskaSpecId int      `ui:"hidden"`
	BacksideDoskaSpec   WoodSpec `ui:"label:Номенклатура боковой доски;field:Id;" gorm:"constraint:OnDelete:CASCADE"`
	BacksideOutPercent  int      `ui:"label:Процент выхода боковой доски"`
	OpilkiSpecId        int      `ui:"hidden"`
	OpilkiSpec          WoodSpec `ui:"label:Номенклатура опилок;field:Id;" gorm:"constraint:OnDelete:CASCADE;"`
	OpilkiOutPercent    int      `ui:"label:Процент выхода опилок"`
}

type SushkaResult struct {
	Id               uint     `gorm:"primaryKey" ui:"hidden"`
	Description      string   `ui:"label:Описание"`
	MaterialSpecId   uint     `ui:"hidden"`
	MaterialSpec     WoodSpec `ui:"label:Номенклатура материала" gorm:"constraint:OnDelete:CASCADE"`
	MaterialQuantity int      `ui:"label:Количество материала"`
	ProductSpecId    uint     `ui:"hidden"`
	ProductSpec      WoodSpec `ui:"label:Номенклатура продукта" gorm:"constraint:OnDelete:CASCADE"`
	ProductAmount    int      `ui:"label:Количество продукта"`
	WorkDate         int64    `ui:"label:Дата работы;datatype:datetime;"`
}

type DryMode struct {
	Id              uint     `gorm:"primaryKey" ui:"hidden"`
	Name            string   `ui:"label:Имя"`
	WetMaterialsId  uint     `ui:"hidden"`
	WetMaterials    WoodSpec `ui:"label:Номенклатура сырых пиломатериалов" gorm:"constraint:OnDelete:CASCADE"`
	HumidityPercent int      `ui:"label:Итоговый процент влажности сухих пиломатериалов"`
	ProcentYsyshki  int      `ui:"label:Процент усушки"`
}
