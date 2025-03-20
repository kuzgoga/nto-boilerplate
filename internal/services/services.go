package services

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

var ExportedServices = []application.Service{
	application.NewService(&WoodSpecTypeService{}),
	application.NewService(&WoodSpecService{}),
	application.NewService(&ExporterService{}),
	application.NewService(&ReceiverService{}),
	application.NewService(&WorkResultService{}),
	application.NewService(&PostavService{}),
	application.NewService(&ReportService{}),
	application.NewService(&SushkaResultService{}),
	application.NewService(&DryModeService{}),
	application.NewService(&DryReportService{}),
}
