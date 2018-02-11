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


func main() {

    engine.RunGoSploit()

    fmt.Printf(engine.Reverse("!oG ,olleH"))

	lines, err := readLines("../xs2pwn/domains.txt")
	if err != nil {

	}

	//Run progress bars
	engine.ProgressBar()

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

	//Print Welcome Info
	multiline := `
	───▄▀▀▀▄▄▄▄▄▄▄▀▀▀▄───
	───█▒▒░░░░░░░░░▒▒█───
	────█░░█░░░░░█░░█────
	─▄▄──█░░░▀█▀░░░█──▄▄─
	█░░█─▀▄░░░░░░░▄▀─█░░█
	`
	color.Red(multiline)
	color.Yellow("gosploit >")

	//Get Module to Load
	reader := bufio.NewReader(os.Stdin)
	mod, _ := reader.ReadString('\n')

	//Run Module
	engine.LoadModule(mod)

}
