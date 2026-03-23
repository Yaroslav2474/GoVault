package screens

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func showErr(msg string, win fyne.Window) {
	dialog.ShowError(fmt.Errorf("%s", msg), win)
}

func tryLogin(password string, win fyne.Window) {
	if password == "" {
		showErr("Please enter a password", win)
		return
	}

	if password == "correct_password" { // Заглушка для проверки пароля
		win.Close()
	} else {
		showErr("Incorrect password", win)
	}

}

func showCreateStorageDialog(win fyne.Window) {
	dialog.ShowConfirm(
		"Initialize account",
		"Create new storage? Previous data will be lost.",
		func(confirm bool) {
			if confirm {
				// Создать логику для создания нового хранилища
				fmt.Println("Creating a new storage...")
			}
		},
		win,
	)
}

func CreateLoginWindow(a fyne.App) fyne.Window {
	win := a.NewWindow("GoVault")

	title := widget.NewLabel("GoVault")
	title.TextStyle.Bold = true
	title.Alignment = fyne.TextAlignCenter

	masterPass := widget.NewPasswordEntry()
	masterPass.PlaceHolder = "Enter master password"
	masterPass.OnSubmitted = func(_ string) {
		tryLogin(masterPass.Text, win)
	}
	masterPass.MinSize(fyne.NewSize(400, 40))

	btnLogin := widget.NewButton("Login", func() {
		tryLogin(masterPass.Text, win)
	})
	btnLogin.Importance = widget.HighImportance
	btnLogin.Resize(fyne.NewSize(300, 40))

	btnCreate := widget.NewButton("Create account", func() {
		showCreateStorageDialog(win)
	})
	btnCreate.Importance = widget.LowImportance

	errorLabel := widget.NewLabel("")
	errorLabel.Hidden = true

	form := container.NewVBox(
		title,
		widget.NewSeparator(),
		masterPass,
		errorLabel,
		container.NewCenter(btnLogin),
		container.NewCenter(btnCreate),
	)

	win.SetContent(container.NewCenter(form))
	win.Resize(fyne.NewSize(500, 300))
	win.SetOnClosed(func() { masterPass.SetText("") })

	return win

}
