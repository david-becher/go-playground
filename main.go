package main

import (
	"log"
)

type city struct {
	cityName  string
	state     string
	residents struct {
		Year2021 int `json:"2021"`
	}
}

func main() {
	fileName := "cities.json"
	cities, err := jsonToStruct(fileName)
	if err != nil {
		log.Fatal(err)
	}

	cityName, maxResidents := cityWithHighestResidents(cities)
	log.Printf("City with highest residents: %s (%d)", cityName, maxResidents)

	sortedList := sortCitiesByResidents(cities)
	err = writeToFile(sortedList)
	if err != nil {
		log.Fatal(err)
	}
}

func jsonToStruct(filename string) ([]city, error) {
	panic("implement me")
}

func cityWithHighestResidents(cities []city) (string, int) {
	var cityName string
	maxResidents := 0
	for _, c := range cities {
		if maxResidents < c.residents.Year2021 {
			maxResidents = c.residents.Year2021
			cityName = c.cityName
		}
	}
	return cityName, maxResidents
}

func sortCitiesByResidents(cities []city) []city {
	panic("implement me")
}

func writeToFile(cities []city) error {
	filename := "cities_sorted.json"
	panic("implement me")
}
