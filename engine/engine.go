package engine

import (
	"fmt"
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
}

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
