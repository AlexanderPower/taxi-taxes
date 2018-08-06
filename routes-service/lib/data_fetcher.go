package lib

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

type Point struct {
	Lat float64
	Lon float64
}

// возвращаемое значение
type Data struct {
	Distance int
	Duration int
}

// Получить значение
func GetData(pointsStr string) (Data, error) {
	points := parsePoints(pointsStr)
	dataResult := Data{0, 0}
	resc, errc := make(chan Data), make(chan error)
	for i := 0; i < len(points)-1; i++ {
		go func(i int) {
			data, err := circularRequest(points[i], points[i+1])
			if err != nil {
				errc <- err
				return
			}
			resc <- data
		}(i)
	}
	for i := 0; i < len(points)-1; i++ {
		select {
		case data := <-resc:
			dataResult.Distance += data.Distance
			dataResult.Duration += data.Duration
		case err := <-errc:
			return Data{0, 0}, err
		}
	}
	return dataResult, nil
}

func parsePoints(str string) []Point {
	var points []Point

	coordsStr := strings.Split(str, "|")
	for i := 0; i < len(coordsStr); i++ {
		coords := strings.Split(coordsStr[i], ",")
		lat, _ := strconv.ParseFloat(coords[0], 64)
		lon, _ := strconv.ParseFloat(coords[1], 64)
		points = append(points, Point{lat, lon})
	}

	return points
}

// делает запрос ко всем указанным источникам данных по очереди
func circularRequest(from Point, to Point) (Data, error) {
	dataSources := config.DataSources
	i := 0
	for {
		data, err := dataSources[i].MakeRequest(from, to)
		log.Printf("response data %v", data)
		if err != nil {
			log.Printf("circularRequest %s", err)
			i++
			if i >= len(dataSources) {
				return Data{0, 0}, errors.New("all data sources are unavailable")
			}
		} else {
			return data, nil
		}
	}
}
