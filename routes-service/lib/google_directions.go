package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Google Directions API
type GoogleDirections struct {
	ApiKey string
}

type googleDirectionsResponseStruct struct {
	Routes []struct {
		Legs []struct {
			Distance struct {
				Value int
			}
			Duration struct {
				Value int
			}
		}
	}
	Status       string
	ErrorMessage string `json:"error_message"`
}

func (gm GoogleDirections) MakeRequest(from Point, to Point) (Data, error) {
	v := url.Values{}
	v.Add("origin", fmt.Sprintf("%v,%v", from.Lat, from.Lon))
	v.Add("destination", fmt.Sprintf("%v,%v", to.Lat, to.Lon))
	v.Add("key", gm.ApiKey)
	// log.Printf("%s\n", v.Encode())

	response, err := http.Get("https://maps.googleapis.com/maps/api/directions/json?" + v.Encode())
	if err != nil {
		return Data{0, 0}, err
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return Data{0, 0}, err
		}
		var responseStruct googleDirectionsResponseStruct
		if err := json.Unmarshal([]byte(contents), &responseStruct); err != nil {
			return Data{0, 0}, err
		}
		if responseStruct.Status != "OK" {
			return Data{0, 0}, errors.New(responseStruct.ErrorMessage)
		}
		return Data{responseStruct.Routes[0].Legs[0].Distance.Value, responseStruct.Routes[0].Legs[0].Duration.Value}, nil
	}
}
