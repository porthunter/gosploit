package main
//go build -buildmode=plugin -o modules/test/eng/chi.so modules/test/eng/.go

import (
	"fmt"
	"bufio"
    "os"
	"github.com/fatih/color"
	"github.com/abiosoft/ishell"
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
		switch mod {
		case "use test/chi/chi\n":
			engine.LoadModule(mod)
		case "shell":
			// create new shell.
		    // by default, new shell includes 'exit', 'help' and 'clear' commands.
		    shell := ishell.New()

		    // display welcome info.
		    shell.Println("Sample Interactive Shell")

		    // register a function for "greet" command.
		    shell.AddCmd(&ishell.Cmd{
		        Name: "greet",
		        Help: "greet user",
		        Func: func(c *ishell.Context) {
		            c.Println("Hello", strings.Join(c.Args, " "))
		        },
		    })

		    // run shell
		    shell.Run()
		default:
			fmt.Println("can't find module")
		}
	}

}
