package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	assertions "github.com/stretchr/testify/assert"
)

// Welcome to the October 2019 A2GO Workshop!
//
// We will be building a CLI tool to fetch current and forecasted weather and print
// to Standard output (the screen).  You might notice in forecast.go that we have the
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
func init() {
	progress = Level{Current: 1}
}

func TestGenerateURL(t *testing.T) {
	assert := assertions.New(t)
	t.Run("URL concatenated from base + apipath + ticket", func(t *testing.T) {
		// Hint #1 -- https://gist.github.com/BrianGenisio/59e493e7f791ddba6d5f353d1d5e1051
		// Hint #2 -- https://gist.github.com/dbraley/1429aecbab3a9f7ae4a8399dc3e7eaea

		// Arrange
		key := "not.important"
		lat := "2"
		long := "-1"

		want := "https://api.forecast.io/forecast/not.important/2,-1"

		// Act
		got := GenerateURL(key, lat, long)

		// Assert
		if assert.Equal(want, got, "Expected URL to be concatenated") {
			progress.SetMessage()
			fmt.Printf(progress.Msg, progress.MsgValues...)
			progress.Current++
		}
	})
}

func TestBuildRequest(t *testing.T) {
	assert := assertions.New(t)
	t.Run("Build Request from url", func(t *testing.T) {
		if progress.Current < 2 {
			t.Skip("Insufficient progress")
		}

		// Docs -- https://golang.org/pkg/net/http
		// Hint #1 -- https://gist.github.com/BrianGenisio/a3ee7551c2b6bf0ca89ff3cd09b3c5c5
		// Hint #2 -- https://gist.github.com/StevenACoffman/cf75067f8d039a0a193c96bad91ae9f3
		// Hint #3 -- https://gist.github.com/StevenACoffman/74fbf609ace0fbb6d6cf2d0da53a5db0

		// Arrange
		url := "https://nonsense.com"

		wantURL := url
		wantJSONMIMEType := "application/json"
		wantMethod := http.MethodGet

		// Act
		got := BuildRequest(url)

		if assert.NotNil(got, "Expected to get a request back, got nil") {

			gotURL := got.URL.String()
			gotMethod := got.Method
			gotAccept := got.Header.Get("Accept")
			gotContentType := got.Header.Get("Content-Type")

			// Assert
			if assert.Equal(wantURL, gotURL) &&
				assert.Equal(wantMethod, gotMethod) &&
				assert.Equal(wantJSONMIMEType, gotContentType, "Expected `Content-Type` Header to be application/json") &&
				assert.Equal(wantJSONMIMEType, gotAccept, "Expected `Accept` Header to be application/json") {

				progress.SetMessage()
				fmt.Printf(progress.Msg, progress.MsgValues...)
				progress.Current++
			}
		}
	})
}

func TestGetBody(t *testing.T) {
	assert := assertions.New(t)
	t.Run("Test Get Body from Response", func(t *testing.T) {
		if progress.Current < 3 {
			t.Skip("Insufficient progress")
		}

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
		if assert.Equal(want, got) {
			progress.SetMessage()
			fmt.Printf(progress.Msg, progress.MsgValues...)
			progress.Current++
		}
	})
}

func TestParseWeatherResponse(t *testing.T) {
	assert := assertions.New(t)

	t.Run("empty response returns empty result", func(t *testing.T) {
		if progress.Current < 4 {
			t.Skip("Insufficient progress")
		}

		// Hint #1 -- https://gist.github.com/StevenACoffman/f688e1d6b20b218443ffbdce6ec773be

		// Arrange
		input := "{}"
		want := Forecast{}

		// Act
		got, err := ParseWeatherResponse(input)

		// Assert
		if assert.NoError(err, "Expected no error from ParseWeatherResponse") &&
			assert.Equal(got, want, "Response did not match what we expected") {

			progress.SetMessage()
			fmt.Printf(progress.Msg, progress.MsgValues...)
			progress.Current++
		}
	})

	t.Run("good response populates result", func(t *testing.T) {
		if progress.Current < 5 {
			t.Skip("Insufficient progress")
		}

		// Docs -- https://golang.org/pkg/encoding/json
		// Hint #2 -- https://gist.github.com/StevenACoffman/d93ccb05b9a3016ff242f1622edc93ad
		// Hint #3 -- https://gist.github.com/StevenACoffman/19ac39ef5b734d5023a55f91e4bc59db

		// Arrange
		input := `{
			"currently": {
				"summary": "This is a test current summary"
			},
			"daily": {
				"summary": "This is a test daily summary"
			}
		}`
		want := Forecast{Currently: CurrentConditions{Summary: "This is a test current summary"}, Daily: WeatherDaily{Summary: "This is a test daily summary"}}
		// Act
		got, err := ParseWeatherResponse(input)

		// Assert
		if assert.NoError(err, "Expected no error from ParseWeatherResponse") &&
			assert.Equal(got, want, "Response did not match what we expected") {

			progress.SetMessage()
			fmt.Printf(progress.Msg, progress.MsgValues...)
			progress.Current++
		}
	})

	t.Run("bad response returns an error", func(t *testing.T) {
		if progress.Current < 6 {
			t.Skip("Insufficient progress")
		}

		// Hint #4 -- https://gist.github.com/StevenACoffman/d49defdf4147fbf9fdc9c8fcdc09f528

		// Arrange
		input := "interrupting cow"

		// Act
		_, err := ParseWeatherResponse(input)

		// Assert
		if assert.Error(err, "Expected an error from ParseWeatherResponse") {
			progress.SetMessage()
			fmt.Printf(progress.Msg, progress.MsgValues...)
			progress.Current++
		}
	})
}

// This is just an example of a benchmark test. You can run it with:
// go test -v --bench . --benchmem
// as a stretch goal you can use benchmark tests to see where to focus make things faster!
func BenchmarkGenerateURL(b *testing.B) {
	var (
		str, longStr string = "my_string", `qwertyuiopqwertyuiopqwertyuio
qwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiopqwertyuiop`
	)
	for i := 0; i < b.N; i++ {
		GenerateURL(str, longStr, strconv.Itoa(i))
	}
}
