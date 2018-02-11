package auxiliary

import (
	"fmt"
	"os"
	"bufio"
	"github.com/sethgrid/multibar"
	"sync"
	"net/http"
)

var lines []string


func XSS_Scan(target string) {

	file, err := os.Open("./payloads/excellent.txt")
    if err != nil {

    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      lines = append(lines, scanner.Text())
    }

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
		for i := 0; i <= len(lines)-1; i++ {
			barProgress3(i)
			resp, err := http.Get("http://"+target+"?payload="+lines[i])
		   if err != nil {
			// handle error
		   }
		   if resp != nil {
			   //c.Println(resp)
		   }
		}

		wg.Done()
	}()


	wg.Wait()

	// continue doing other work
	fmt.Println("All Bars Complete")


}
