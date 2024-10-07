package main

import (
	"errors"
	"net/http"
)

func CheckClickHouseAvailability() error {
	resp, err := http.Get(`http://localhost:8123/?query=SELECT%201`)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("unexpected status code: " + resp.Status)
	}

	return nil
}
