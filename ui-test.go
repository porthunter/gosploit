package main

import ui "github.com/gizak/termui" // <- ui shortcut, optional

	func main() {
		err := ui.Init()
		if err != nil {
			panic(err)
		}
		defer ui.Close()

		p := ui.NewPar(":PRESS q TO QUIT DEMO")
		p.Height = 3
		p.Width = 50
		p.TextFgColor = ui.ColorWhite
		p.BorderLabel = "Text Box"
		p.BorderFg = ui.ColorCyan

		g := ui.NewGauge()
		g.Percent = 50
		g.Width = 50
		g.Height = 3
		g.Y = 11
		g.BorderLabel = "Gauge"
		g.BarColor = ui.ColorRed
		g.BorderFg = ui.ColorWhite
		g.BorderLabelFg = ui.ColorCyan

		ui.Render(p, g) // feel free to call Render, it's async and non-block

		// event handler...
	}
