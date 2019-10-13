package main

import (
	"bytes"
	assertions "github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

// Welcome to the October 2019 A2GO Workshop!
//
// We will be building a CLI tool to fetch current and forecasted weather and print
// to Standard Output (the screen).  You might notice in forecast.go that we have the
// shell of a file.  We've implemented the basics for running the code.  Only problem:
// none of the actual behavior has been developed yet.
//
// This is up to you.  We've even written the tests for you!  Consider them to be
// the specs for your functions.  Start with the first test and get it to pass.
// Test your implementation by running `go test` from the command line.
//
// Once you've gotten the first test passing, go to the next test.
// IMPORTANT: you must delete the t.Skip("") lines when you are ready to start working
// on a new test.  We didn't want you to start with ALL the tests yelling at you.
//
// Hints:
// We've structured the hints to help you along the way.
// The last hint in a function will give you the code to make the test pass.
// Try to do it without the hints, but don't feel bad if you need to use them.  We're all learning!
//
// Try it when you are done:
// go run forecast.go -forecast
//
// Stretch goals:
// What else can you do to make this tool more useful?  Some ideas:
// - Better error handling and reporting
// - Use environment variables or CLI input for some hard-coded values (username, password, etc)
// - Display more information to the user
// - Allow searching by keyword
// - Use the existing benchmark and find the fastest string concatenation method
// - Write more integration tests (but keep them separate from unit tests)
//
// If you complete any of the stretch goals, be sure to add tests!!!

func TestGenerateURL(t *testing.T) {
	assert := assertions.New(t)
	t.Run("URL concatenated from base + apipath + ticket", func(t *testing.T) {
		// Hint #1 -- https://gist.github.com/BrianGenisio/59e493e7f791ddba6d5f353d1d5e1051
		// Hint #2 -- https://gist.github.com/BrianGenisio/75f89a244970a3cdfc1e36f56d2e2edd

		// Arrange
		key := "not.important"
		lat := "2"
		long := "-1"

		want := "https://api.forecast.io/forecast/not.important/2,-1"

		// Act
		got := GenerateURL(key, lat, long)

		// Assert
		assert.Equal(want, got, "Expected URL to be concatenated")
	})
}

func TestBuildRequest(t *testing.T) {
	assert := assertions.New(t)
	t.Run("Build Request from url jira UserId and Jira Password", func(t *testing.T) {
		//t.Skip("Delete this when ready to go next")

		// Docs -- https://golang.org/pkg/net/http
		// Hint #1 -- https://gist.github.com/BrianGenisio/a3ee7551c2b6bf0ca89ff3cd09b3c5c5
		// Hint #2 -- https://gist.github.com/BrianGenisio/90f4a514df4b2e75e05ca9fa9f059aee
		// Hint #3 -- https://gist.github.com/BrianGenisio/36c39d6c0ec3fdd4d45bfd9b2ff10777
		// Hint #4 -- https://gist.github.com/BrianGenisio/952c490f42698b952800972b716346df

		// Arrange
		url := "https://nonsense.com"

		wantURL := url
		wantJSONMIMEType := "application/json"

		// Act
		got := BuildRequest(url)

		if assert.NotNil(got, "Expected to get a request back, got nil") {

			gotURL := got.URL.String()
			gotAccept := got.Header.Get("Accept")
			gotContentType := got.Header.Get("Content-Type")

			// Assert
			assert.Equal(wantURL, gotURL)
			assert.Equal(wantJSONMIMEType, gotContentType)
			assert.Equal(wantJSONMIMEType, gotAccept)
		}
	})
}

func TestGetBody(t *testing.T) {
	assert := assertions.New(t)
	t.Run("Test Get Body from Response", func(t *testing.T) {
		//t.Skip("Delete this when ready to go next")

		// Docs -- https://golang.org/pkg/io
		// Hint #1 -- https://gist.github.com/BrianGenisio/cf696a9e29883a5089f8ddd725e20651
		// Hint #2 -- https://gist.github.com/BrianGenisio/85a66b473c17c0face36ce5430502ae9
		// Hint #3 -- https://gist.github.com/BrianGenisio/1154eeb5880a5606fa20ed6f3dd82696
		// Hint #4 -- https://gist.github.com/BrianGenisio/358aec390f3b80b4037bd6a6c3cf9313

		// Arrange
		dummyResponse := &http.Response{
			StatusCode: 200,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
		// Act
		got := GetBody(dummyResponse)
		want := "OK"

		// Assert
		assert.Equal(want, got)
	})
}

func TestParseJiraResponse(t *testing.T) {
	assert := assertions.New(t)

	t.Run("empty response returns empty result", func(t *testing.T) {
		//t.Skip("Delete this when ready to go next")

		// Hint #1 -- https://gist.github.com/BrianGenisio/8634fc8f14c5024f55a9cd6b18adc90c

		// Arrange
		input := "{}"
		want := Forecast{}

		// Act
		got, err := ParseWeatherResponse(input)

		// Assert
		assert.NoError(err, "Expected no error from getJiraResponse")
		assert.Equal(got, want, "Response did not match what we expected")
	})

	t.Run("good response populates result", func(t *testing.T) {
		//t.Skip("Delete this when ready to go next")

		// Docs -- https://golang.org/pkg/encoding/json
		// Hint #2 -- https://gist.github.com/BrianGenisio/f1e0646c5be39da8dd1da77e8aca60d8
		// Hint #3 -- https://gist.github.com/BrianGenisio/0ec105a1031a209953db8f1a94acdfa0

		// Arrange
		input := `{
			"currently": {
				"summary": "This is a test current summary"
			},
			"daily": {
				"summary": "This is a test daily summary"
			}
		}`
		want := Forecast{Currently: CurrentConditions{Summary: "This is a test current summary",}, Daily: WeatherDaily{Summary: "This is a test daily summary"},}
		// Act
		got, err := ParseWeatherResponse(input)

		// Assert
		assert.NoError(err, "Expected no error from getJiraResponse")
		assert.Equal(got, want, "Response did not match what we expected")
	})

	t.Run("bad response returns an error", func(t *testing.T) {
		t.Skip("Delete this when ready to go next")

		// Hint #4 -- https://gist.github.com/BrianGenisio/6bd8dad67dc00e7d647f0e6353a0d486

		// Arrange
		input := "interrupting cow"

		// Act
		_, err := ParseWeatherResponse(input)

		// Assert
		assert.Error(err, "Expected an error from getJiraResponse")
	})
}

func BenchmarkGenerateURL(b *testing.B) {
	var (
		str, longStr string = "my_string", `qwertyuiopqwertyuiopqwertyuio
qwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiop`
	)
	for i := 0; i < b.N; i++ {
		GenerateURL(str, longStr, strconv.Itoa(i))
	}
}
