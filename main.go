package main

import (
	"app/internal/services"
	"embed"
	"github.com/wailsapp/wails/v3/pkg/application"
	"log"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := application.New(application.Options{
		Name:        "nto_starterkit",
		Description: "NTO toolkit template",
		Services:    append([]application.Service{services.MigratorService}, services.ExportedServices...),
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: "Завод \"Белочка\"",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		URL: "/",
	})

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
