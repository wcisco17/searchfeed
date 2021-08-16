module github.com/wcisco17/searchfeed/main

go 1.16

replace github.com/wcisco17/searchfeed/search => ../search

replace github.com/wcisco17/searchfeed/matchers => ../matchers

replace github.com/wcisco17/searchfeed/prompt => ../prompt

require (
	github.com/wcisco17/searchfeed/matchers v0.0.0-00010101000000-000000000000
	github.com/wcisco17/searchfeed/prompt v0.0.0-00010101000000-000000000000
	github.com/wcisco17/searchfeed/search v0.0.0-00010101000000-000000000000
)
