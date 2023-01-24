package main

import "log"

type city struct {
	cityName  string
	state     string
	residents struct {
		Year2021 int `json:"2021"`
	}
}

func main() {
	fileName := "cities.json"
	cities := jsonToStruct(fileName)

	maxResidents := cityWithHighestResidents(cities)
	log.Printf("City with highest residents: %d", maxResidents)

	sortedList := sortCitiesByResidents(cities)
	_ = sortedList
}

func jsonToStruct(fileName string) []city {
	panic("implement me")
}

func cityWithHighestResidents(cities []city) int {
	maxResidents := 0
	for _, c := range cities {
		if maxResidents < c.residents.Year2021 {
			maxResidents = c.residents.Year2021
		}
	}
	return maxResidents
}

func sortCitiesByResidents(cities []city) []city {
	panic("implement me")

}
