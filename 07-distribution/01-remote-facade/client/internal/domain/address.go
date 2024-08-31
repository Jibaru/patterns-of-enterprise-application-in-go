package domain

import "github.com/jibaru/remote-facade-client/internal/remote"

type Address struct {
	street string
	city   string
	zip    string
	Facade *remote.AddressFacade
}

func (a *Address) load() error {
	err := a.Facade.GetAddressData()
	if err != nil {
		return err
	}
	a.city = a.Facade.City
	a.street = a.Facade.Street
	a.zip = a.Facade.Zip

	return nil
}

func (a *Address) Street() (string, error) {
	err := a.load()
	if err != nil {
		return "", err
	}

	return a.street, nil
}

func (a *Address) City() (string, error) {
	err := a.load()
	if err != nil {
		return "", err
	}

	return a.city, nil
}

func (a *Address) Zip() (string, error) {
	err := a.load()
	if err != nil {
		return "", err
	}

	return a.zip, nil
}

func (a *Address) SetStreet(val string) error {
	a.street = val
	return a.Facade.SetAddress(a.street, a.city, a.zip)
}

func (a *Address) SetCity(val string) error {
	a.city = val
	return a.Facade.SetAddress(a.street, a.city, a.zip)
}

func (a *Address) SetZip(val string) error {
	a.zip = val
	return a.Facade.SetAddress(a.street, a.city, a.zip)
}
