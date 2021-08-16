package main

import (
	"log"
	"os"

	_ "github.com/wcisco17/searchfeed/matchers"
	p "github.com/wcisco17/searchfeed/prompt"
	"github.com/wcisco17/searchfeed/search"
)

// All init() function in any go code will be executed first
func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	term, err := p.FilterSearch()
	if err != nil {
		log.Fatal(err)
	}

	site, errs := p.FilterSelect()
	if errs != nil {
		log.Fatal(errs)
	}

	search.Run(term, site)
}
