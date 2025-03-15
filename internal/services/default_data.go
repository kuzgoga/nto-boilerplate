package services

import (
	"app/internal/dialogs"
	"fmt"
	"time"
)

func InsertDefaultData() {
	insertWorkshops()
	insertProductTypes()
	insertCustomers()
	insertOrders()
	insertTasks()
	insertPrepTasks()
	insertShifts()
	insertTeamTypes()
	insertWorkers()
	insertTeamTasks()
}

func InsertDefaultEntityData[T any](service Service[T], entities []T) {
	for _, item := range entities {
		createdItem, err := service.Create(item)
		if err != nil {
			dialogs.ErrorDialog("Ошибка при вставке данных по умолчанию", fmt.Sprintf("Произошла ошибка при вставке значения %#v: %s", createdItem, err))
		}
	}
}

func insertWorkshops() {
	InsertDefaultEntityData(&WorkshopService{}, []Workshop{
		{
			Id:   1,
			Name: "Лесопильный комплекс",
			WorkAreas: []WorkArea{
				{
					Id:          1,
					Name:        "Лесопильная линия №1",
					Description: "Используется для распиловки тонкомеров, например, реек.",
					PrepTasks:   nil,
					Performance: 50,
				},
				{
					Id:          2,
					Name:        "Лесопильная линия №2",
					Description: "Используется для распиловки среднего леса.",
					PrepTasks:   nil,
					Performance: 100,
				},
			},
			Tasks: nil,
		},
		{
			Id:   2,
			Name: "Сушильный комплекс",
			WorkAreas: []WorkArea{
				{Id: 3, Name: "Сушильная камера №1", Performance: 50},
				{Id: 4, Name: "Сушильная камера №2", Performance: 60},
				{Id: 5, Name: "Сушильная камера №3", Performance: 80},
				{Id: 6, Name: "Сушильная камера №4", Performance: 85},
			},
			Tasks: nil,
		},
		{
			Id:   3,
			Name: "Цех строжки и обработки",
			WorkAreas: []WorkArea{
				{
					Id:          7,
					Name:        "Линия строжки №1",
					Description: "Используется для строжки тонкомеров",
					Performance: 50,
				},
				{
					Id:          8,
					Name:        "Линия строжки №2",
					Description: "Используется для строжки среднего леса",
					Performance: 80,
				},
				{
					Id:          9,
					Name:        "Линия строжки №3",
					Description: "Используется для строжки среднего леса",
					Performance: 100,
				},
			},
			Tasks: nil,
		},
		{
			Id:   4,
			Name: "Пеллетный цех",
			WorkAreas: []WorkArea{
				{Id: 10, Name: "Дробилка", Performance: 800},
				{Id: 11, Name: "Сушилка", Performance: 900},
				{Id: 12, Name: "Гранулятор №1", Performance: 900},
				{Id: 13, Name: "Гранулятор №2", Performance: 600},
			},
			Tasks: nil,
		},
	})
}

func insertProductTypes() {
	InsertDefaultEntityData(&ProductTypeService{}, []ProductType{
		{Id: 1, Name: "Сырые пиломатериалы"},
		{Id: 2, Name: "Сухие пиломатериалы"},
		{Id: 3, Name: "Строганные доски"},
		{Id: 4, Name: "Рейки"},
		{Id: 5, Name: "Брус"},
		{Id: 6, Name: "Пеллеты"},
	})
}

func insertCustomers() {
	InsertDefaultEntityData(&CustomerService{}, []Customer{
		{
			Id:      1,
			Title:   "Rivalli",
			Contact: "Телефон: +74955855525",
			Orders:  []Order{},
		},
		{
			Id:      2,
			Title:   "Elemfort",
			Contact: "Телефон: +79270988888",
			Orders:  []Order{},
		},
		{
			Id:      3,
			Title:   "Квазар",
			Contact: "Телефон: +78002342134",
			Orders:  []Order{},
		},
	})
}

func insertOrders() {
	InsertDefaultEntityData(&OrderService{}, []Order{
		{
			Status:        "Согласован клиентом",
			ProductAmount: 400,
			ProductTypeId: 1,
			Description:   "Сырая древесина для Rivalli",
			DeadlineDate:  time.Date(2025, 6, 25, 0, 0, 0, 0, time.Local).Unix(),
			CustomerId:    1,
		},
		{
			Status:        "Согласован клиентом",
			ProductTypeId: 3,
			ProductAmount: 300,
			Description:   "Сухая древесина для производителя мебели. Контракт #574853",
			DeadlineDate:  time.Date(2025, 6, 10, 0, 0, 0, 0, time.Local).Unix(),
			CustomerId:    2,
		},
		{
			Status:        "Согласован клиентом",
			ProductAmount: 100,
			ProductTypeId: 4,
			Description:   "Контракт #234342",
			DeadlineDate:  time.Date(2025, 12, 5, 0, 0, 0, 0, time.Local).Unix(),
			CustomerId:    3,
		},
		{
			Status:        "Согласован клиентом",
			ProductAmount: 800,
			ProductTypeId: 6,
			Description:   "Производство пеллет",
			DeadlineDate:  time.Date(2025, 2, 5, 0, 0, 0, 0, time.Local).Unix(),
			CustomerId:    3,
		},
		{
			Status:        "Черновик",
			ProductAmount: 100,
			ProductTypeId: 6,
			Description:   "Производство пеллет. Заказ на стадии согласования",
			DeadlineDate:  time.Date(2026, 2, 5, 0, 0, 0, 0, time.Local).Unix(),
			CustomerId:    1,
		},
	})
}

func insertTasks() {
	InsertDefaultEntityData(&TaskService{}, []Task{
		{
			Id:              1,
			Description:     "Обработка сырых пиломатериалов",
			ProductTypeId:   2,
			Workshops:       []*Workshop{{Id: 1}},
			OrderId:         1,
			ProductionStart: time.Date(2025, 11, 25, 0, 0, 0, 0, time.Local).Unix(),
			Amount:          10,
		},
		{
			Id:              2,
			Description:     "Распиловка леса",
			ProductTypeId:   1,
			Workshops:       []*Workshop{{Id: 1}, {Id: 2}},
			OrderId:         2,
			ProductionStart: time.Date(2025, 8, 15, 0, 0, 0, 0, time.Local).Unix(),
			Amount:          1,
		},
		{
			Id:              3,
			Description:     "Строгание реек",
			ProductTypeId:   3,
			Workshops:       []*Workshop{{Id: 1}, {Id: 2}},
			OrderId:         3,
			ProductionStart: time.Date(2025, 7, 10, 0, 0, 0, 0, time.Local).Unix(),
			Amount:          5,
		},
		{
			Id:              4,
			Description:     "Производство пеллет",
			ProductTypeId:   6,
			Workshops:       []*Workshop{{Id: 4}},
			OrderId:         4,
			ProductionStart: time.Date(2025, 4, 5, 0, 0, 0, 0, time.Local).Unix(),
			Amount:          1,
		},
	})
}

func insertPrepTasks() {
	InsertDefaultEntityData(&PrepTaskService{}, []PrepTask{
		{
			Id:          1,
			Description: "Подготовка лесопильной линии. Материал для обработки: сырые пиломатериалы",
			TaskId:      1,
			WorkAreaId:  2,
			Deadline:    time.Date(2025, 11, 10, 0, 0, 0, 0, time.Local).Unix(),
		},
		{
			Id:          2,
			Description: "Подготовка лесопильной линии. Материал для обработки: сухие пиломатериалы",
			TaskId:      2,
			WorkAreaId:  1,
			Deadline:    time.Date(2025, 8, 10, 0, 0, 0, 0, time.Local).Unix(),
		},
		{
			Id:          3,
			Description: "Подготовка сушильной камеры. Сушка распиленных сухих пиломатериалов",
			TaskId:      2,
			WorkAreaId:  3,
			Deadline:    time.Date(2025, 8, 15, 0, 0, 0, 0, time.Local).Unix(),
		},
		{
			Id:          4,
			Description: "Подготовка лесопильного цеха для производства реек. Сушка древесины для производства реек",
			TaskId:      3,
			WorkAreaId:  2,
			Deadline:    time.Date(2025, 8, 15, 0, 0, 0, 0, time.Local).Unix(),
		},
		{
			Id:          5,
			Description: "Подготовка сушильной камеры. Сушка распиленных сухих реек",
			TaskId:      3,
			WorkAreaId:  3,
			Deadline:    time.Date(2025, 8, 15, 0, 0, 0, 0, time.Local).Unix(),
		},
		{
			Id:          6,
			Description: "Подготовка дробилки для производства пеллет",
			TaskId:      4,
			WorkAreaId:  10,
			Deadline:    time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local).Unix(),
		},
		{
			Id:          7,
			Description: "Подготовка сушилки для производства пеллет",
			TaskId:      4,
			WorkAreaId:  11,
			Deadline:    time.Date(2025, 4, 2, 0, 0, 0, 0, time.Local).Unix(),
		},
		{
			Id:          8,
			Description: "Подготовка гранулятора №1 для производства пеллет",
			TaskId:      4,
			WorkAreaId:  12,
			Deadline:    time.Date(2025, 4, 3, 0, 0, 0, 0, time.Local).Unix(),
		},
		{
			Id:          9,
			Description: "Подготовка гранулятора №2 для производства пеллет",
			TaskId:      4,
			WorkAreaId:  13,
			Deadline:    time.Date(2025, 4, 4, 0, 0, 0, 0, time.Local).Unix(),
		},
	})
}

func insertShifts() {
	InsertDefaultEntityData(&ShiftService{}, []Shift{
		{
			Id:            1,
			Description:   "Распилка лесоматериалов",
			ProductTypeId: 1,
			ProductAmount: 45,
			ShiftDate:     time.Date(2025, 11, 20, 0, 0, 0, 0, time.Local).Unix(),
			WorkAreaId:    2,
		},
		{
			Id:            2,
			Description:   "Обработка сухой древесины",
			ProductTypeId: 3,
			ProductAmount: 39,
			ShiftDate:     time.Date(2025, 8, 13, 0, 0, 0, 0, time.Local).Unix(),
			WorkAreaId:    1,
		},
		{
			Id:            3,
			Description:   "Сушка реек",
			ProductTypeId: 4,
			ProductAmount: 55,
			ShiftDate:     time.Date(2025, 7, 8, 0, 0, 0, 0, time.Local).Unix(),
			WorkAreaId:    4,
		},
		{
			Id:            4,
			Description:   "Грануляция паллет",
			ProductTypeId: 6,
			ProductAmount: 890,
			ShiftDate:     time.Date(2025, 4, 1, 0, 0, 0, 0, time.Local).Unix(),
			WorkAreaId:    12,
		},
	})
}

func insertTeamTypes() {
	InsertDefaultEntityData(&TeamTypeService{}, []TeamType{
		{Id: 1, Name: "Бригада по распиловке"},
		{Id: 2, Name: "Бригада по сушке"},
		{Id: 3, Name: "Бригада по обработке"},
		{Id: 4, Name: "Бригада пеллетного цеха"},
	})
}

func insertWorkers() {
	InsertDefaultEntityData(&WorkerService{}, []Worker{
		{Id: 1, Name: "Иванов Иван Иванович", WorkshopId: 1},
		{Id: 2, Name: "Петров Петр Петрович", WorkshopId: 1},
		{Id: 3, Name: "Сидоров Сидор Сидорович", WorkshopId: 1},
		{Id: 4, Name: "Пеллетчиков Пеллет Пеллетович", WorkshopId: 4},
		{Id: 5, Name: "Кузнецов Кузьма Кузьмич", WorkshopId: 2},
		{Id: 6, Name: "Смирнов Сергей Сергеевич", WorkshopId: 2},
		{Id: 7, Name: "Васильев Василий Васильевич", WorkshopId: 3},
		{Id: 8, Name: "Михайлов Михаил Михайлович", WorkshopId: 3},
		{Id: 10, Name: "Александров Александр Александрович", WorkshopId: 4},
		{Id: 11, Name: "Николаев Николай Николаевич", WorkshopId: 4},
		{Id: 12, Name: "Дмитриев Дмитрий Дмитриевич", WorkshopId: 4},
		{Id: 13, Name: "Дмитриев Дмитрий Федорович", WorkshopId: 4},
	})
}

func insertTeamTasks() {
	InsertDefaultEntityData(&TeamTaskService{}, []TeamTask{
		{
			Id:         1,
			TeamTypeId: 1,
			TeamMembers: []*Worker{
				{Id: 1},
			},
			WorkAreaId:    2,
			ShiftDuties:   "1/1",
			WorkStartDate: time.Date(2024, 12, 5, 0, 0, 0, 0, time.Local).Unix(),
			TeamLeaderId:  13,
		},
		{
			Id:         2,
			TeamTypeId: 1,
			TeamMembers: []*Worker{
				{Id: 2}, {Id: 3},
			},
			WorkAreaId:    1,
			ShiftDuties:   "1/1",
			WorkStartDate: time.Date(2024, 12, 10, 0, 0, 0, 0, time.Local).Unix(),
			TeamLeaderId:  2,
		},
		{
			Id:         3,
			TeamTypeId: 2,
			TeamMembers: []*Worker{
				{Id: 5}, {Id: 6},
			},
			WorkAreaId:    3,
			ShiftDuties:   "1/1",
			WorkStartDate: time.Date(2025, 1, 7, 0, 0, 0, 0, time.Local).Unix(),
			TeamLeaderId:  5,
		},
		{
			Id:         4,
			TeamTypeId: 3,
			TeamMembers: []*Worker{
				{Id: 7}, {Id: 8},
			},
			WorkAreaId:    8,
			ShiftDuties:   "2/2",
			WorkStartDate: time.Date(2025, 2, 10, 0, 0, 0, 0, time.Local).Unix(),
			TeamLeaderId:  7,
		},
		{
			Id:         5,
			TeamTypeId: 4,
			TeamMembers: []*Worker{
				{Id: 4}, {Id: 10}, {Id: 11}, {Id: 12},
			},
			WorkAreaId:    12,
			ShiftDuties:   "2/2",
			WorkStartDate: time.Date(2025, 3, 10, 0, 0, 0, 0, time.Local).Unix(),
			TeamLeaderId:  10,
		},
	})
}
