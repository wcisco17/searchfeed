![Screen Shot 2021-08-16 at 7 19 17 PM](https://user-images.githubusercontent.com/35783824/129640843-34a69ff5-6995-4a16-ac19-a02a063e860b.png)
# searchfeed

Searching a list of news outlet by using keyword, expanded functionality to search and picking search by site using the promptui lib.

### Running
```sh
main && go run main.go
```

### Improvements
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
 
 reference: https://www.manning.com/books/go-in-action
