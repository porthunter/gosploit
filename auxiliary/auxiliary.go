package auxiliary

import (
	"fmt"
	"github.com/sethgrid/multibar"
	"sync"
	"net/http"
	"../utility"
	"strings"
	"bytes"
	"github.com/fatih/color"
)

func XSS_Scan(target string) {

	lines, err := utility.ReadLines("./payloads/excellent.txt")
	if err != nil {

	}



	// create the multibar container
	// this allows our bars to work together without stomping on one another
	progressBars, _ := multibar.New()

	// some arbitrary totals for our  progress bars
	// in practice, these could be file sizes or similar
	largerTotal := len(lines)-1

	barProgress3 := progressBars.MakeBar(largerTotal, "Payloads:")

	// listen in for changes on the progress bars
	// I should be able to move this into the constructor at some point
	go progressBars.Listen()

	wg := &sync.WaitGroup{}
	wg.Add(1)


	go func() {
		found := 0
		for i := 0; i <= len(lines)-1; i++ {
			barProgress3(i)
			resp, err := http.Get("https://"+target+"/?query="+lines[i])
			defer resp.Body.Close()

			if err != nil {

			}

			buf := new(bytes.Buffer)
		    buf.ReadFrom(resp.Body)
		    body := buf.String()

			if strings.Contains(body, lines[i]) {
				found++
			}
		}
		if found >0 {
			color.Yellow("[i] Payloads were found in responses, the site could be vulnerable")
		}

		wg.Done()
	}()


	wg.Wait()

	// continue doing other work
	fmt.Println("All Tests Complete")


}
