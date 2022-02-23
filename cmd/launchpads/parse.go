package launches

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const spacexLaunchpadsApi = "https://api.spacexdata.com/v4/launchpads"

type spacexLaunchPad struct {
	ID string `json:"id"`
}

func GetParsedLaunchpads() []spacexLaunchPad {
	response, err := http.Get(spacexLaunchpadsApi)

	if err != nil {
		log.Fatal(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var spacexData []spacexLaunchPad

	err = json.Unmarshal(responseData, &spacexData)
	if err != nil {
		log.Fatal(err.Error())
	}

	return spacexData
}
