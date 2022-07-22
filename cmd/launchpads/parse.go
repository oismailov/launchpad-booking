package launches

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const spacexLaunchpadsApi = "https://api.spacexdata.com/v4/launchpads"

type spacexLaunchPad struct {
	ID string `json:"id"`
}

func GetParsedLaunchpads() ([]spacexLaunchPad, error) {
	response, err := http.Get(spacexLaunchpadsApi)
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var spacexData []spacexLaunchPad

	err = json.Unmarshal(responseData, &spacexData)
	if err != nil {
		return nil, err
	}

	return spacexData, nil
}
