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
	DetalleLugar, _ := c.PlaceDetails(context.Background(), result)

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

	result, _ := c.FindPlaceFromText(context.Background(), findLugar)

	var placeID string
	for _, can := range result.Candidates {
		placeID = can.PlaceID
	}
	return placeID
}
