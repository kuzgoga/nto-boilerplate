package models

var Entities = []any{
	&Customer{}, &Order{}, &PrepTask{}, &ProductType{}, &Shift{}, &Task{}, &TeamTask{}, &WorkArea{}, &Worker{}, &Workshop{}, &TeamType{},
}

type Customer struct {
	Id      uint    `gorm:"primaryKey" ui:"label:ID;readonly"`
	Title   string  `ui:"label:Название"`
	Contact string  `ui:"label:Контакт"`
	Orders  []Order `gorm:"constraint:OnDelete:CASCADE;" ui:"hidden"`
}

type Order struct {
	Id            uint        `gorm:"primaryKey" ui:"label:ID;readonly"`
	Status        string      `ui:"label:Статус"`
	Description   string      `ui:"label:Описание"`
	ProductTypeId int         `gorm:"not null;" ui:"hidden"`
	ProductType   ProductType `gorm:"not null;foreignKey:ProductTypeId;references:Id;constraint:OnDelete:CASCADE;" ui:"label:Тип;field:Name"`
	ProductAmount uint        `ui:"label:Количество продукции"`
	CustomerId    uint        `gorm:"not null;" ui:"hidden"`
	Customer      Customer    `gorm:"not null;foreignKey:CustomerId;references:Id;constraint:OnDelete:CASCADE;" ui:"label:Клиент;field:Title"`
	Tasks         []Task      `gorm:"constraint:OnDelete:CASCADE" ui:"hidden"`
	CreatedAt     int64       `gorm:"autoCreateTime" ui:"label:Дата создания;readonly;datatype:datetime"`
	DeadlineDate  int64       `ui:"label:Крайний срок;datatype:datetime"`
}

type PrepTask struct {
	Id          uint     `gorm:"primaryKey" ui:"label:ID;readonly"`
	Status      string   `gorm:"default:'Создано'" ui:"label:Статус"`
	Description string   `ui:"label:Описание"`
	TaskId      uint     `gorm:"not null;" ui:"hidden"`
	Task        Task     `gorm:"foreignKey:TaskId;references:Id;constraint:OnDelete:CASCADE;" ui:"label:Задача;field:Description"`
	WorkAreaId  uint     `gorm:"not null;" ui:"hidden"`
	WorkArea    WorkArea `gorm:"foreignKey:WorkAreaId;references:Id;constraint:OnDelete:CASCADE;" ui:"label:Рабочая зона;field:Name"`
	CreatedAt   int64    `gorm:"autoCreationTime" ui:"label:Дата создания;readonly;datatype:datetime"`
	Deadline    int64    `ui:"label:Крайний срок;datatype:datetime"`
}

type ProductType struct {
	Id   uint   `gorm:"primaryKey" ui:"label:ID;readonly"`
	Name string `gorm:"not null" ui:"label:Название"`
}

type Shift struct {
	Id            uint        `gorm:"primaryKey" ui:"label:ID;readonly"`
	Description   string      `ui:"label:Описание"`
	ProductTypeId uint        `ui:"hidden"`
	ProductType   ProductType `ui:"field:Name"`
	ProductAmount uint        `ui:"label:Количество продукции"`
	ShiftDate     int64       `ui:"label:Дата смены;datatype:datetime"`
	WorkAreaId    uint        `ui:"hidden"`
	WorkArea      WorkArea    `ui:"field:Name"`
	CreatedAt     int64       `gorm:"autoCreateTime" ui:"label:Дата создания;readonly;datatype:datetime"`
}

type Task struct {
	Id              uint        `gorm:"primaryKey" ui:"label:ID;readonly"`
	Description     string      `ui:"label:Описание"`
	ProductTypeId   uint        `gorm:"not null;" ui:"hidden"`
	ProductType     ProductType `gorm:"foreignKey:ProductTypeId;references:Id;constraint:OnDelete:CASCADE;" ui:"label:Тип;field:Name"`
	Workshops       []*Workshop `gorm:"many2many:workshop_task;constraint:OnDelete:CASCADE;" ui:"hidden"`
	OrderId         uint        `gorm:"not null;" ui:"hidden"`
	Order           Order       `gorm:"foreignKey:OrderId;references:Id;constraint:OnDelete:CASCADE;" ui:"label:Заказ;field:Description"`
	PrepTasks       []PrepTask  `gorm:"constraint:OnDelete:CASCADE;" ui:"hidden"`
	ProductionStart int64       `ui:"label:Дата начала производства;datatype:datetime"`
	CreatedAt       int64       `gorm:"autoCreateTime" ui:"label:Дата создания;readonly;datatype:datetime"`
	Amount          uint        `ui:"label:Количество"`
}

type TeamTask struct {
	Id            uint      `gorm:"primaryKey" ui:"label:ID;readonly"`
	TeamTypeId    uint      `ui:"hidden"`
	TeamType      TeamType  `ui:"field:Name"`
	TeamLeaderId  uint      `ui:"hidden"`
	TeamLeader    Worker    `ui:"field:Name"`
	TeamMembers   []*Worker `gorm:"many2many:worker_team_tasks;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;" ui:"hidden"`
	WorkStartDate int64     `ui:"label:Дата начала работ;datatype:datetime"`
	WorkAreaId    uint      `ui:"hidden"`
	WorkArea      WorkArea  `ui:"field:Name"`
	ShiftDuties   string    `gorm:"check:shift_duties IN ('1/1','2/2')" ui:"label:Обязанности смены"`
}

type TeamType struct {
	Id   uint   `gorm:"primaryKey" ui:"label:ID;readonly"`
	Name string `gorm:"not null" ui:"label:Название"`
}

type WorkArea struct {
	Id          uint       `gorm:"primaryKey" ui:"label:ID;readonly"`
	Name        string     `ui:"label:Наименование"`
	Description string     `ui:"label:Описание"`
	Performance uint       `ui:"label:Производительность"`
	WorkshopId  uint       `gorm:"not null;" ui:"hidden"`
	Workshop    Workshop   `gorm:"foreignKey:WorkshopId;references:Id;" ui:"label:Цех;field:Name"`
	PrepTasks   []PrepTask `gorm:"constraint:OnDelete:CASCADE;" ui:"hidden"`
	Shifts      []Shift    `gorm:"constraint:OnDelete:CASCADE;" ui:"hidden"`
	TeamTasks   []TeamTask `gorm:"foreignKey:Id;constraint:OnDelete:CASCADE;" ui:"hidden"`
}

type Worker struct {
	Id         uint        `gorm:"primaryKey" ui:"label:ID;readonly"`
	Name       string      `ui:"label:Имя"`
	TeamTasks  []*TeamTask `gorm:"many2many:worker_team_tasks;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;" ui:"hidden"`
	Workshop   Workshop    `gorm:"foreignKey:WorkshopId;references:Id;" ui:"field:Name"`
	WorkshopId uint        `ui:"hidden"`
}

type Workshop struct {
	Id        uint       `gorm:"primaryKey" ui:"label:ID;readonly"`
	Name      string     `ui:"label:Наименование"`
	WorkAreas []WorkArea `gorm:"constraint:OnDelete:CASCADE;" ui:"hidden"`
	Tasks     []*Task    `gorm:"many2many:workshop_task;constraint:OnDelete:CASCADE;" ui:"hidden"`
	Workers   []Worker   `gorm:"constraint:OnDelete:CASCADE;" ui:"hidden"`
}
