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
	parse := kingpin.Command("parse", "Parse a project")

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
	case parse.FullCommand():
		fmt.Println("Parsing a project...")
		f, err := os.Open("spec.yaml")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		var s spec.Spec
		if err := yaml.NewDecoder(f).Decode(&s); err != nil {
			panic(err)
		}
		s.Parse()
		fmt.Println("Done!")
	default:
		kingpin.Usage()
	}
}
