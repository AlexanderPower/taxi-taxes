package lib

import (
	"log"
	"os"
	"strings"
)

// конфигурация
type Config struct {
	// источники данных
	DataSources []DataSourcer
}

var config Config

func init() {
	config = getConfig()
}

// DATA_SOURCES - источники данных через запятую
func getConfig() Config {
	dataSources := []DataSourcer{}

	for _, dataSource := range strings.Split(os.Getenv("DATA_SOURCES"), ",") {
		log.Println(dataSource)
		switch dataSource {
		case "google_directions":
			apiKey, present := os.LookupEnv("GOOGLE_DIRECTIONS_API_KEY")
			if !present {
				panic("GOOGLE_DIRECTIONS_API_KEY not found")
			}
			dataSources = append(dataSources, GoogleDirections{apiKey})
		case "google_matrix":
			apiKey, present := os.LookupEnv("GOOGLE_MATRIX_API_KEY")
			if !present {
				panic("GOOGLE_MATRIX_API_KEY not found")
			}
			dataSources = append(dataSources, GoogleMatrix{apiKey})
		default:
			log.Printf("%s data source not inplemented and will be ignored.", dataSource)
		}
	}

	return Config{dataSources}
}
