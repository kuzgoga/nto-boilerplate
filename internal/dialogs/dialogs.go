package dialogs

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

var currentWindow *application.WebviewWindow

func Init(window *application.WebviewWindow) {
	if window == nil {
		panic("currentWindow is nil")
	}
	currentWindow = window
}

func checkInit() {
	if currentWindow == nil {
		panic("Initialize dialogs package before use")
	}
}

func InfoDialog(title string, message string) {
	checkInit()
	application.InfoDialog().AttachToWindow(currentWindow).SetTitle(title).SetMessage(message).Show()
}

func WarningDialog(title string, message string) {
	checkInit()
	application.WarningDialog().AttachToWindow(currentWindow).SetTitle(title).SetMessage(message).Show()
}

func ErrorDialog(title string, message string) {
	checkInit()
	application.ErrorDialog().AttachToWindow(currentWindow).SetTitle(title).SetMessage(message).Show()
}

func SaveFileDialog(title string, filename string) (string, error) {
	checkInit()
	selection, err := application.SaveFileDialog().AttachToWindow(currentWindow).SetFilename(filename).PromptForSingleSelection()
	if err != nil {
		return "", err
	}
	return selection, nil
}
