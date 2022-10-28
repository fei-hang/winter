package main

import (
	"log"
	"winter/src/mysql"
)

func main() {

	// MyApp := app.New()
	// fanyiW := MyApp.NewWindow("百度翻译")
	// fanyiW.Resize(fyne.Size{Width: 400, Height: 500})
	// input := widget.NewEntry()
	// input.OnChanged = fanyi
	// box := container.NewVBox(input, outbox)
	// fanyiW.SetContent(box)
	// fanyiW.ShowAndRun()

	log.Println(mysql.InitDB())
}
