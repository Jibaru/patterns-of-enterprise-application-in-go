package remote

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const RemoteBaseUrl = "http://localhost:8080"

type AddressFacade struct {
	Street string
	City   string
	Zip    string
	loaded bool
}

func (f *AddressFacade) GetAddressData() error {
	if f.loaded {
		return nil
	}

	resp, err := http.Get(RemoteBaseUrl + "/address")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &f); err != nil {
		return err
	}

	f.loaded = true

	return nil
}

func (f *AddressFacade) SetAddress(street, city, zip string) error {
	f.City = city
	f.Street = street
	f.Zip = zip

	jsonData, err := json.Marshal(f)
	if err != nil {
		return err
	}

	_, err = http.Post(RemoteBaseUrl+"/address", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	return nil
}
