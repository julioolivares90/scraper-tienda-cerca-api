package scraper

import (
	"context"
	"log"

	"googlemaps.github.io/maps"
)

var (
	apiKey = "AIzaSyA372lgSayQt_UW6EIb4IAjEABGg2elWrg"
)

func FindLugarByID(id string) maps.PlaceDetailsResult {
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	result := &maps.PlaceDetailsRequest{
		PlaceID:      id,
		Language:     "",
		Fields:       nil,
		SessionToken: maps.PlaceAutocompleteSessionToken{},
		Region:       "",
	}
	DetalleLugar, err := c.PlaceDetails(context.Background(), result)
	if err != nil {
		log.Fatal(err)
	}

	return DetalleLugar
}

func GetIDMunicipio(lugar string) string {
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	findLugar := &maps.FindPlaceFromTextRequest{
		Input:                 lugar,
		InputType:             "textquery",
		Fields:                nil,
		LocationBias:          "",
		LocationBiasPoint:     nil,
		LocationBiasCenter:    nil,
		LocationBiasRadius:    0,
		LocationBiasSouthWest: nil,
		LocationBiasNorthEast: nil,
	}

	result, errorFindPlace := c.FindPlaceFromText(context.Background(), findLugar)
	if errorFindPlace != nil {
		log.Fatal(errorFindPlace)
	}
	var placeID string
	for _, can := range result.Candidates {
		placeID = can.PlaceID
	}
	return placeID
}
