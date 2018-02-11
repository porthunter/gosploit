package main

import "github.com/gizak/termui"

func main() {
    if err:=termui.Init(); err != nil {
        panic(err)
    }
    defer termui.Close()

	bc := termui.NewBarChart()
	data := []int{3, 2, 5, 3, 9, 5}
	bclabels := []string{"S0", "S1", "S2", "S3", "S4", "S5"}
	bc.BorderLabel = "Bar Chart"
	bc.Data = data
	bc.Width = 26
	bc.Height = 10
	bc.DataLabels = bclabels
	bc.TextColor = termui.ColorGreen
	bc.BarColor = termui.ColorRed
	bc.NumColor = termui.ColorYellow

    termui.Render(bc)

    termui.Loop()
}
