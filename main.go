package main

import (
	"fmt"
	"os"

	"github.com/snowmerak/rest-make/spec"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v3"
)

func main() {
	init := kingpin.Command("init", "Initialize a new project")

	switch kingpin.Parse() {
	case init.FullCommand():
		fmt.Println("Initializing a new project...")
		f, err := os.Create("spec.yaml")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if err := yaml.NewEncoder(f).Encode(spec.Init()); err != nil {
			panic(err)
		}
		fmt.Println("Done!")
	default:
		kingpin.Usage()
	}
}
