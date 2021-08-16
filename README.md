# searchfeed

Searching a list of news outlet by using keyword, expanded functionality to search and Picking search by site using the promptui lib.

- Improvements convert func to accept any array duplicate.
- Improve the time complexity of func to run Time = O(n), Space = O(1) - Considering remove the additional array.
  - Current time complexity Time = O(n^2) (Due to looping over the data once then again with the result arr)
 ```go
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
 ```
