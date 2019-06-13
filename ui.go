package main

import "github.com/andlabs/ui"

func CreateUI() {
	err := ui.Main(func() {
		// input := ui.NewEntry()
		button1 := ui.NewButton("生成文件")
		button2 := ui.NewButton("比对结果")
		combox := ui.NewCombobox()

		greeting := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("选择模版文件:"), false)
		// box.Append(input, false)

		box.Append(combox, false)
		combox.Append("表格一")
		combox.Append("表格二")

		box.Append(button1, false)
		box.Append(button2, false)
		box.Append(greeting, false)
		window := ui.NewWindow("Hello", 200, 100, false)
		window.SetMargined(true)
		window.SetChild(box)
		button1.OnClicked(func(*ui.Button) {
			greeting.SetText("生成文件") // + input.Text() + "!")
		})
		button2.OnClicked(func(*ui.Button) {
			greeting.SetText("比对结果")
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
