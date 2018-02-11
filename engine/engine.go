package engine

import (
	"fmt"
	"os/exec"
	tm "github.com/buger/goterm"
)

func RunGoSploit() {
	tm.Clear() // Clear current screen
	// Create Box with 30% width of current screen, and height of 20 lines
	box := tm.NewBox(30|tm.PCT, 20, 0)
	// Add some content to the box
	// Note that you can add ANY content, even tables
	fmt.Fprint(box, "text in a box")
	// Move Box to approx center of the screen
	tm.Print(tm.MoveTo(box.String(), 0|tm.PCT, 40|tm.PCT))
	tm.Flush()

	app := "nmap"
    //app := "buah"

    arg0 := "-sV"
    arg1 := "-sC"
    arg2 := "evsec.com"
    arg3 := "-p 443"

    cmd := exec.Command(app, arg0, arg1, arg2, arg3)
    stdout, err := cmd.Output()

    if err != nil {
        println(err.Error())
        return
    }

    print(string(stdout))
}

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
