package prompt

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	res "github.com/wcisco17/searchfeed/search"
)

func removedup(feeds []*res.Feed) []string {
	var hastable = make(map[string]bool)
	var input = make([]string, 0)

	for _, feed := range feeds {
		if ok := hastable[feed.Name]; !ok {
			hastable[feed.Name] = false
		}
	}

	for idx := range hastable {
		input = append(input, idx)
	}

	return input
}

func FilterSelect() (string, error) {
	feeds, errs := res.RetrieveFeeds()
	if errs != nil {
		return "", nil
	}

	value := removedup(feeds)

	prompt := promptui.Select{
		Label: "Select news station..",
		Items: value,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", nil
	}

	return result, nil
}

func FilterSearch() (string, error) {
	var input string

	term := promptui.Prompt{
		Label:   "Search...",
		Default: input,
	}

	result, err := term.Run()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return result, nil
}
