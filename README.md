# go-dojo
The scaffolding for the Go Dojo we'll be hosting at the A2Go Meetup on 10/2/2019

##  Welcome to the October 2019 A2GO Workshop!

We will be building a CLI tool to fetch current and forecasted weather and print to Standard Output (the screen).  You might notice in `forecast.go` that we have the shell of a file.  We've implemented the basics for running the code.  Only problem: none of the actual behavior has been developed yet.

*This is up to you.*  We've even written the tests for you!  Consider them to be the specs for your functions.  Start with the first test and get it to pass.

Test your implementation by running `go test` from the command line.

Once you've gotten the first test passing, go to the next test.

*IMPORTANT:* you must delete the `t.Skip("")` lines when you are ready to start working on a new test.  We didn't want you to start with ALL the tests yelling at you.

### Hints:
We've structured the hints to help you along the way. The last hint in a function will give you the code to make the test pass.  Try to do it without the hints, but don't feel bad if you need to use them.  We're all learning!

### Try it when you are done:
```go run forecast.go -forecast```

### Stretch goals:
What else can you do to make this tool more useful?  Some ideas:
- Better error handling and reporting
- Use environment variables or CLI input for some hard-coded values (username, password, etc)
- Display more information to the user
- Allow searching by keyword
- Use the existing benchmark and find the fastest string concatenation method
- Write more integration tests (but keep them separate from unit tests)

If you complete any of the stretch goals, be sure to add tests!!!

##### Test coverage Reports
```
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
go tool cover -func=coverage.out
```

##### Test Benchmarks
If the test suite contains benchmarks, you can run these with the `--bench` and `--benchmem` flags:

```
go test -v --bench . --benchmem
```
Keep in mind that each reviewer will run benchmarks on a different machine, with different specs, so the results from these benchmark tests may vary.

##### Seperating Integration Tests]
You can build the executable and run a simple integration test if you do:
```shell script
go build
PATH=.:$PATH go test -tags=integration
```
It is important to keep slow and possibly destructive integration tests separated from unit tests. 
 
The tag used in the build comment cannot have a dash, although underscores are allowed. For example, `// +build unit-tests` will not work, whereas `// +build unit_tests` will.

##### Create your own API token
In order to use the Dark Sky API, you first need your own API key. Getting an API key is quick and free. 

* https://darksky.net/dev
* Click “Try for Free”.
* Register an account and click the link sent to you in a validation email to activate your account
* Sign in
* You get 1,000 API calls per day with your free Dark Sky Developer account. There is no credit card required unless you want to upgrade to an account that will allow you more than 1,000 API calls per day.

Your secret Dark Sky API key will look something like this: 0123456789abcdef9876543210fedcba. Save it.

<table><tr><td>:bulb: <b>NOTE:</b> You should store the token securely, just as for any password.<br/>

</td></tr></table>

##### Test an API token
A primary use case for API tokens is to allow scripts to access REST APIs for applications using HTTP. Often the tokens will be sent as headers or as part of the URL.

For example, when using curl, you could do something like this:
```shell script
curl -v -L \
https://api.darksky.net/forecast/32772f4b37c5a08eb4488a2ce79155bd/37.8267,-122.4233 
```

<table><tr><td>:bulb: <b>NOTE:</b> this API keys is intended to be replaced with one from your own account.<br/>

</td></tr></table>

Look at all of that glorious weather data!

You can also make an API call to Dark Sky by typing in a URL into your browser in the following format:

    https://api.darksky.net/forecast/[key]/[latitude],[longitude]

for example, to get the weather in Boston, first get the latitude and longitude coordinates (for example from Google Maps):

* https://api.darksky.net/forecast/0123456789abcdef9876543210fedcba/42.3601,-71.0589

Let's look at what is returned (you will have to place your Dark Sky API key in the above link):




##### Resources for future exploration
+ [Prefer table driven tests](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)
+ [Go Walkthrough: encoding/json package](https://medium.com/go-walkthrough/go-walkthrough-encoding-json-package-9681d1d37a8f)
+ [Lesser known features of go test](https://splice.com/blog/lesser-known-features-go-test/)
+ [Unit Testing HTTP Client in Go](http://hassansin.github.io/Unit-Testing-http-client-in-Go)
+ [Peter Bourgon Best Practices](http://peter.bourgon.org/go-best-practices-2016/)
+ [How I write Go after 7 years](https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831)
+ [Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)
+ [Structuring Applications in Go](https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091)
+ [Dave Cheney's High Performance Go Workshop Content](https://dave.cheney.net/high-performance-go-workshop/gophercon-2019.html)
+ [Go database/sql tutorial](http://go-database-sql.org/)
+ [How to work with Postgres in Go
](https://medium.com/avitotech/how-to-work-with-postgres-in-go-bad2dabd13e4)

##### Well-known struct tags

Go offers [struct tags](https://golang.org/ref/spec#Tag). Tags are the backticked strings you sometimes see at the end of structs, which are discoverable via reflection. These enjoy a wide range of use in the standard library in the JSON/XML and other encoding packages. 
```
type User struct {
        Name    string `json:"name"`
        Age     int    `json:"age,omitempty"`
        Zipcode int    `json:"zipcode,string"`
}
```
The json struct tag options include:
+ Renaming the field’s key. A lot of JSON keys are camel cased so it can be important to change the name to match.
+ The omitempty flag can be set which will remove any non-struct fields which have an empty value.
+ The string flag can be used to force a field to encode as a string. For example, forcing an int field to be encoded as a quoted string.

The community welcomed struct tags and has built ORMs, further encodings, flag parsers and much more around them since, especially for these tasks, single-sourcing is beneficial for data structures.

Tag       | Documentation
----------|---------------
xml       | https://godoc.org/encoding/xml
json      | https://godoc.org/encoding/json
asn1      | https://godoc.org/encoding/asn1
reform    | https://godoc.org/gopkg.in/reform.v1
dynamodb  | https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/dynamodbattribute/#Marshal
bigquery  | https://godoc.org/cloud.google.com/go/bigquery
datastore | https://godoc.org/cloud.google.com/go/datastore
spanner   | https://godoc.org/cloud.google.com/go/spanner
gorm      | https://godoc.org/github.com/jinzhu/gorm
yaml      | https://godoc.org/gopkg.in/yaml.v2
validate  | https://github.com/go-playground/validator
mapstructure | https://godoc.org/github.com/mitchellh/mapstructure
protobuf  | https://github.com/golang/protobuf
