package homeTab

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func GetHomeTab(a fyne.App) *fyne.Container {
	blue := color.NRGBA{R: 0, G: 128, B: 191, A: 255}
	//white := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	//red := color.NRGBA{R: 180, G: 0, B: 0, A: 255}

	mainTitle := canvas.NewText("Home", blue)
	mainTitle.Alignment = fyne.TextAlignCenter
	mainTitle.TextSize = 32

	textBox := canvas.NewText("Made with love ❌ Made with gay ✔️", blue)
	textBox.Alignment = fyne.TextAlignCenter
	textBox.TextSize = 18

	//basically just do nothing when the checkbox is checked
	//obfuscateCheckBox := widget.NewCheck("Obfuscate", func(_ bool) {})
	//obfuscateCheckBox.SetChecked(false)
	entryLayout := container.New(layout.NewVBoxLayout(),
		mainTitle,
		//webhookEntry,
		//obfuscateCheckBox,
		textBox,
		layout.NewSpacer(),
		//compileButtonPS1,
		//compileButtonBAT,
		//removeButton,
		//comepileButtonEXE,
	)

	return entryLayout
}
