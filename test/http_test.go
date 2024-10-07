package main

import (
	"testing"
)

func TestCheckClickHouseAvailability(t *testing.T) {
	t.Parallel()

	if err := CheckClickHouseAvailability(); err != nil {
		t.Fatal(err)
	}
}
