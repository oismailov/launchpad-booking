package spacex

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const spacexLaunchesApi = "https://api.spacexdata.com/v4/launches"

type spacexLaunch struct {
	Launchpad string    `json:"launchpad"`
	DateLocal time.Time `json:"date_local"`
}

func GetParsedData() []spacexLaunch {
	response, err := http.Get(spacexLaunchesApi)

	if err != nil {
		log.Fatal(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var spacexData []spacexLaunch

	err = json.Unmarshal(responseData, &spacexData)
	if err != nil {
		log.Fatal(err.Error())
	}

	return spacexData
}
