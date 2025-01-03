package dialogs

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

func InfoDialog(title string, message string) {
	application.InfoDialog().SetTitle(title).SetMessage(message).Show()
}

func WarningDialog(title string, message string) {
	application.WarningDialog().SetTitle(title).SetMessage(message).Show()
}

func ErrorDialog(title string, message string) {
	application.ErrorDialog().SetTitle(title).SetMessage(message).Show()
}
