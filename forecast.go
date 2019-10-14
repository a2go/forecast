///usr/bin/env go run "$0" "$@" ; exit "$?"
package main

import (
	"errors"
	"flag"
	"fmt"
	logger "log"
	"net/http"
	"os"
	"time"
)

const (
	//replace with your personal information as desired
	apiURLFmt = "https://api.forecast.io/forecast/%s/%s,%s"
	key       = "32772f4b37c5a08eb4488a2ce79155bd"
	latitude  = "42.2797"  // Ann Arbor Latitude
	longitude = "-83.7369" // Ann Arbor Longitude
	layoutUS  = "January 2, 2006"
)

func main() {
	log := logger.New(os.Stderr, "", 0)
	forecast := getFlag()
	fc, _ := GetForecast(key, latitude, longitude)
	Output(fc, forecast, log)
}

func getFlag() bool {
	flag.Usage = help
	forecastPtr := flag.Bool("forecast", false, "Show 8 day forecast")
	flag.Parse()
	return *forecastPtr
}

func help() {
	fmt.Printf("Usage: forecast [flags]\n")
	flag.PrintDefaults()
}

// This is a "glue" function.  It takes all of the more testable, behavioral
// functions and "glues" them together without any other inherent behavior
func GetForecast(key, latitude, longitude string) (Forecast, error) {
	url := GenerateURL(key, latitude, longitude)
	weatherClient := http.Client{}
	req := BuildRequest(url)
	response, _ := weatherClient.Do(req)
	if response.StatusCode != http.StatusOK {
		return Forecast{}, errors.New("forbidden - most likely due to invalid token")
	}
	body := GetBody(response)
	return ParseWeatherResponse(body)
}

// GenerateURL will construct the JIRA API call from components
func GenerateURL(key, latitude, longitude string) string {
	return ""
}

// BuildRequest will build a new client and request with the proper
// headers, including basic authentication
func BuildRequest(url string) *http.Request {
	return nil
}

// GetBody will take an httpResponse and extract the body as a string
func GetBody(res *http.Response) string {
	return ""
}

// ParseJiraResponse will parse the Jira response into a JIRAResponse
func ParseWeatherResponse(jsonData string) (Forecast, error) {
	return Forecast{}, nil
}

func Output(fc Forecast, forecast bool, log *logger.Logger) {
	cur := fc.Currently
	daily := fc.Daily
	if !forecast {
		curTime := time.Unix(cur.Time, 0).Format(time.RFC822)

		curWeatherFormat := `
	      Current Weather: %s
	        Summary     %s
	        Temperature %f
	        Humidity    %f
	        WindSpeed   %f
	        WindBearing %f
	      `
		log.Printf(curWeatherFormat, curTime, cur.Summary, cur.Temperature, cur.Humidity, cur.WindSpeed, cur.WindBearing)

	} else {
		var dailys string

		for _, v := range daily.Data {
			dTime := time.Unix(v.Time, 0).Format(layoutUS)
			dailyForecastFormat := `
	      Weather for %s
	        Summary         %s
	        Temperature Min %f
	        Temperature Max %f
	        Humidity        %f
	        WindSpeed       %f
	        WindBearing     %f
	      `
			dailys += fmt.Sprintf(dailyForecastFormat, dTime, v.Summary, v.TemperatureMin, v.TemperatureMax, v.Humidity, v.WindSpeed, v.WindBearing)
		}
		log.Println(dailys)
	}
}

type Forecast struct {
	Currently CurrentConditions `json:"currently"`
	Daily     WeatherDaily      `json:"daily"`
}

type CurrentConditions struct {
	Time        int64
	Summary     string
	Temperature float32
	Humidity    float32
	WindSpeed   float32
	WindBearing float32
}

type WeatherDaily struct {
	Summary string
	Data    []struct {
		Time           int64
		Summary        string
		TemperatureMin float32
		TemperatureMax float32
		Humidity       float32
		WindSpeed      float32
		WindBearing    float32
	} `json:"data"`
}
