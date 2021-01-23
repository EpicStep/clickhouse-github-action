package main

import (
	"testing"
)

func TestSendQuery(t *testing.T) {
	t.Parallel()

	_, err := SendQuery("SELECT version() FORMAT TabSeparated")
	if err != nil {
		t.Error(err)
	}
}