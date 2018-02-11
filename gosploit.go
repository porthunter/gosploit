package main
//go build -buildmode=plugin -o modules/test/eng/chi.so modules/test/eng/.go

import (
    "./engine"
)

func main() {

	for {
		engine.RunShell()
	}

}
