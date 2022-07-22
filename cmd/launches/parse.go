package launches

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const spacexLaunchesApi = "https://api.spacexdata.com/v4/launches/upcoming"

type spacexLaunch struct {
	LaunchID    string    `json:"id"`
	LaunchpadID string    `json:"launchpad"`
	DateLocal   time.Time `json:"date_local"`
}

func GetParsedLaunches() ([]spacexLaunch, error) {
	response, err := http.Get(spacexLaunchesApi)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var spacexData []spacexLaunch
	err = json.Unmarshal(responseData, &spacexData)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return spacexData, nil
}
