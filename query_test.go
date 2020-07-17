package openweather

import (
	"testing"

	"github.com/EricNeid/go-openweather/internal/verify"
)

func TestNewQueryForCity(t *testing.T) {
	// arrange
	apiKey := "testKey"
	location := "Berlin,de"
	// action
	q := NewQueryForCity(apiKey, location)
	// verify
	verify.Equals(t, apiKey, q.APIKey)
	verify.Equals(t, location, q.Query)
	verify.Equals(t, "metric", q.Unit)
	verify.Equals(t, queryTypeCity, q.queryType)

	// arrange
	unit := "imperial"
	// action
	q = NewQueryForCity(apiKey, location, unit)
	// verify
	verify.Equals(t, unit, q.Unit)
}

func TestNewQueryForZip(t *testing.T) {
	// arrange
	apiKey := "testKey"
	zip := "12345"
	// action
	q := NewQueryForZip(apiKey, zip)
	// verify
	verify.Equals(t, apiKey, q.APIKey)
	verify.Equals(t, zip, q.Query)
	verify.Equals(t, queryTypeZip, q.queryType)
}

func TestNewQueryForID(t *testing.T) {
	// arrange
	apiKey := "testKey"
	id := "42"
	// action
	q := NewQueryForID(apiKey, id)
	// verify
	verify.Equals(t, apiKey, q.APIKey)
	verify.Equals(t, id, q.Query)
	verify.Equals(t, queryTypeID, q.queryType)
}

func TestNewQueryForLocation(t *testing.T) {
	// arrange
	apiKey := "testKey"
	lat := "51"
	lon := "13"
	// action
	q := NewQueryForLocation(apiKey, lat, lon)
	// verify
	verify.Equals(t, apiKey, q.APIKey)
	verify.Equals(t, lat+"|"+lon, q.Query)
	verify.Equals(t, queryTypeGeo, q.queryType)
}
