package screens

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type fixedWidthLayout struct {
	width float32
}

func (f *fixedWidthLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	for _, obj := range objects {
		objSize := fyne.NewSize(f.width, obj.MinSize().Height)
		obj.Resize(objSize)
		obj.Move(fyne.NewPos((size.Width-f.width)/2, 0))
	}
}

func (f *fixedWidthLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(f.width, 0)
}

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

	masterPassContainer := container.New(&fixedWidthLayout{width: 400}, masterPass)

	buttonsContainer := container.NewHBox(btnLogin, widget.NewLabel("    "), btnCreate)
	centeredButtons := container.NewCenter(buttonsContainer)

	content := container.NewVBox(
		title,
		widget.NewSeparator(),
		widget.NewLabel(""),
		masterPassContainer,
		widget.NewLabel(""),
		centeredButtons,
		widget.NewLabel(""),
	)

	win.SetContent(content)
	win.Resize(fyne.NewSize(500, 350))
	win.SetOnClosed(func() { masterPass.SetText("") })

	return win

}
