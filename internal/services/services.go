package services

import (
	"fmt"
	"github.com/wailsapp/wails/v3/pkg/application"
)

var ExportedServices = []application.Service{
	application.NewService(&CustomerService{}),
	application.NewService(&OrderService{}),
	application.NewService(&PrepTaskService{}),
	application.NewService(&ProductTypeService{}),
	application.NewService(&ShiftService{}),
	application.NewService(&TaskService{}),
	application.NewService(&TeamTaskService{}),
	application.NewService(&TeamTypeService{}),
	application.NewService(&WorkAreaService{}),
	application.NewService(&WorkerService{}),
	application.NewService(&WorkshopService{}),
}

func CreateServices(items ...any) []application.Service {
	var services = make([]application.Service, len(items))
	for i, item := range items {
		services[i] = application.NewService(&item)
	}
	fmt.Printf("%#+v", services)
	return services
}
