package main
//go build -buildmode=plugin -o modules/test/eng/chi.so modules/test/eng/.go

import (
	"bufio"
	"fmt"
	"time"
    "net/http"
    "os"
	"github.com/fatih/color"
	"plugin"
	tm "github.com/buger/goterm"
)

var lines []string

func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

// Here's the worker, of which we'll run several
// concurrent instances. These workers will receive
// work on the `jobs` channel and send the corresponding
// results on `results`. We'll sleep a second per job to
// simulate an expensive task.
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {

		started := j
		finished := 300
		// Based on http://golang.org/pkg/text/tabwriter
		totals := tm.NewTable(0, 10, 5, ' ', 0)
		fmt.Fprintf(totals, "Time\tStarted\tActive\tFinished\n")
		fmt.Fprintf(totals, "%s\t%d\t%d\t%d\n", "All", started, started-finished, finished)
		tm.Println(totals)
		tm.Flush()



        time.Sleep(time.Second)
        resp, err := http.Get("https://"+lines[j])
        if err != nil {
        	// handle error
        }
        if resp != nil {
            fmt.Println(resp)
        } else {
            resp, err := http.Get("http://"+lines[j])
            if err != nil {
            	// handle error
            }
            if resp != nil {
                fmt.Println(resp)
            }
        }
    }
}

type GosploitModule interface {
	Exploit()
}

func main() {
	tm.Clear() // Clear current screen

	// Create Box with 30% width of current screen, and height of 20 lines
	box := tm.NewBox(30|tm.PCT, 20, 0)

	// Add some content to the box
	// Note that you can add ANY content, even tables
	fmt.Fprint(box, "text in a box")

	// Move Box to approx center of the screen
	tm.Print(tm.MoveTo(box.String(), 0|tm.PCT, 40|tm.PCT))

	tm.Flush()

	reader := bufio.NewReader(os.Stdin)

	multiline := `
	───▄▀▀▀▄▄▄▄▄▄▄▀▀▀▄───
	───█▒▒░░░░░░░░░▒▒█───
	────█░░█░░░░░█░░█────
	─▄▄──█░░░▀█▀░░░█──▄▄─
	█░░█─▀▄░░░░░░░▄▀─█░░█
	`

	color.Red(multiline)

	color.Yellow("gosploit >")
	text, _ := reader.ReadString('\n')
	//fmt.Println(text)

	var mod string
	switch text {
	case "english\n":
		mod = "./modules/test/eng/eng.so"
	case "chinese\n":
		mod = "./modules/test/chi/chi.so"
	default:
		fmt.Println("don't speak that language")
		os.Exit(1)
	}

	// load module
	// 1. open the so file to load the symbols
	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 2. look up a symbol (an exported function or variable)
	// in this case, variable GosploitModule
	symGosploitModule, err := plug.Lookup("GosploitModule")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 3. Assert that loaded symbol is of a desired type
	// in this case interface type GosploitModule (defined above)
	var module GosploitModule
	module, ok := symGosploitModule.(GosploitModule)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

    lines, err := readLines("../xs2pwn/domains.txt")
    if err != nil {
		fmt.Println(text)
    }

    // In order to use our pool of workers we need to send
    // them work and collect their results. We make 2
    // channels for this.
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // This starts up 3 workers, initially blocked
    // because there are no jobs yet.
    for w := 1; w <= 200; w++ {
        go worker(w, jobs, results)
    }

    // Here we send 5 `jobs` and then `close` that
    // channel to indicate that's all the work we have.
    for j := 1; j <= len(lines); j++ {
        jobs <- j
    }
    close(jobs)
	
	// 4. use the module
	module.Exploit()

}
