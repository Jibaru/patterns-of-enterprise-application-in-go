package main

import (
	"fmt"

	"github.com/jibaru/remote-facade-client/internal/domain"
	"github.com/jibaru/remote-facade-client/internal/remote"
	"github.com/jibaru/remote-facade-client/internal/utils"
)

func main() {
	address := domain.Address{
		Facade: &remote.AddressFacade{},
	}

	street, err := address.Street()
	if err != nil {
		panic(err)
	}

	city, err := address.City()
	if err != nil {
		panic(err)
	}

	zip, err := address.Zip()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Street: %v, City: %v, Zip: %v\n", street, city, zip)

	street, city, zip = utils.RandomAddressValues()

	err = address.SetStreet(street)
	if err != nil {
		panic(err)
	}

	err = address.SetCity(city)
	if err != nil {
		panic(err)
	}

	err = address.SetZip(zip)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Address set: %v\n", address)
}
