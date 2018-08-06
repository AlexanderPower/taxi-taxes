package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Google Distance Matrix API
type GoogleMatrix struct {
	ApiKey string
}

type googleMatrixResponseStruct struct {
	Rows []struct {
		Elements []struct {
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

func (gm GoogleMatrix) MakeRequest(from Point, to Point) (Data, error) {
	v := url.Values{}
	v.Add("origins", fmt.Sprintf("%v,%v", from.Lat, from.Lon))
	v.Add("destinations", fmt.Sprintf("%v,%v", to.Lat, to.Lon))
	v.Add("key", gm.ApiKey)

	response, err := http.Get("https://maps.googleapis.com/maps/api/distancematrix/json?" + v.Encode())
	if err != nil {
		return Data{0, 0}, err
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return Data{0, 0}, err
		}
		var responseStruct googleMatrixResponseStruct
		if err := json.Unmarshal([]byte(contents), &responseStruct); err != nil {
			return Data{0, 0}, err
		}
		if responseStruct.Status != "OK" {
			return Data{0, 0}, errors.New(responseStruct.ErrorMessage)
		}
		return Data{responseStruct.Rows[0].Elements[0].Distance.Value, responseStruct.Rows[0].Elements[0].Duration.Value}, nil
	}
}
