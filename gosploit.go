package main
//go build -buildmode=plugin -o modules/test/eng/chi.so modules/test/eng/.go

import (
    "./engine"
	"github.com/fatih/color"
)

func main() {

	//Print Welcome Info
	multiline := `
	───▄▀▀▀▄▄▄▄▄▄▄▀▀▀▄───
	───█▒▒░░░░░░░░░▒▒█───
	────█░░█░░░░░█░░█────
	─▄▄──█░░░▀█▀░░░█──▄▄─
	█░░█─▀▄░░░░░░░▄▀─█░░█
	`
	color.Red(multiline)

	for {
		engine.RunShell()
	}

}
