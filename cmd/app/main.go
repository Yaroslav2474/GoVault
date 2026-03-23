package main

import (
	"GoVault/ui/screens"

	"fyne.io/fyne/v2/app"
)

func main() {
	App := app.New()
	loginWin := screens.CreateLoginWindow(App)
	loginWin.ShowAndRun()
}
