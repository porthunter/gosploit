package main
//go build -buildmode=plugin -o modules/test/eng/chi.so modules/test/eng/.go

import (
	"bufio"
	"fmt"
	"time"
    "net/http"
    "os"
	"github.com/fatih/color"
    "./engine"
)

type GosploitModule interface {
	Exploit()
}

func main() {

    //engine.RunGoSploit()

    fmt.Printf(engine.Reverse("!oG ,olleH"))


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
		color.Yellow("gosploit >")

		//Get Module to Load
		reader := bufio.NewReader(os.Stdin)
		mod, _ := reader.ReadString('\n')

		//Run Module
		engine.LoadModule(mod)
	}

}
