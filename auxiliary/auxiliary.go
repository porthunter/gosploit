package auxiliary

import (
	"fmt"
	"os/exec"
	tm "github.com/buger/goterm"
)

func XSS_Scan(target string) {
	tm.Clear() // Clear current screen

	app := "nmap"
	//app := "buah"

	arg0 := "-sV"
	arg1 := "-sC"
	arg2 := target

	cmd := exec.Command(app, arg0, arg1, arg2)
	stdout, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	// Create Box with 30% width of current screen, and height of 20 lines
	box := tm.NewBox(100|tm.PCT, 6, 0)
	// Add some content to the box
	// Note that you can add ANY content, even tables
	fmt.Fprint(box, string(stdout))
	// Move Box to approx center of the screen
	tm.Print(tm.MoveTo(box.String(), 0|tm.PCT, 40|tm.PCT))

}
