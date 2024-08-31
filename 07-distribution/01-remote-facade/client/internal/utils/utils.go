package utils

import "math/rand"

func RandomAddressValues() (street string, city string, zip string) {
	var streets = []string{
		"Elm Street", "Maple Avenue", "Oak Drive", "Pine Lane", "Cedar Road",
		"Birch Boulevard", "Cherry Street", "Willow Way", "Ash Court", "Spruce Terrace",
	}

	var cities = []string{
		"Springfield", "Rivertown", "Lakeview", "Hilltop", "Greenwood",
		"Sunnydale", "Riverbend", "Westfield", "Mapleton", "Brookhaven",
	}

	var zips = []string{
		"12345", "67890", "11223", "45678", "98765",
		"54321", "67812", "34567", "23456", "78901",
	}

	street = streets[rand.Intn(len(streets))]
	city = cities[rand.Intn(len(cities))]
	zip = zips[rand.Intn(len(zips))]

	return street, city, zip
}
