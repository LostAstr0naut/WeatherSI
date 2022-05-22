package location

import (
	"errors"
)

// Location represents a struct with data of location.
type Location struct {
	X int32
	Y int32
}

func LocationCoordinates(locationName string) (Location, error) {
	location, ok := locations()[locationName]
	if !ok {
		return Location{}, errors.New("invalid location")
	}

	return location, nil
}

func locations() map[string]Location {
	return map[string]Location{
		"BO": {
			X: 224,
			Y: 290,
		},
		"GO": {
			X: 237,
			Y: 374,
		},
		"KP": {
			X: 249,
			Y: 464,
		},
		"ID": {
			X: 296,
			Y: 364,
		},
		"JE": {
			X: 300,
			Y: 268,
		},
		"PO": {
			X: 325,
			Y: 417,
		},
		"KR": {
			X: 346,
			Y: 311,
		},
		"LJ": {
			X: 373,
			Y: 350,
		},
		"KO": {
			X: 424,
			Y: 445,
		},
		"TB": {
			X: 454,
			Y: 332,
		},
		"SG": {
			X: 459,
			Y: 252,
		},
		"NM": {
			X: 474,
			Y: 410,
		},
		"CE": {
			X: 484,
			Y: 310,
		},
		"KK": {
			X: 521,
			Y: 374,
		},
		"MB": {
			X: 545,
			Y: 247,
		},
		"PT": {
			X: 581,
			Y: 270,
		},
		"MS": {
			X: 625,
			Y: 216,
		},
	}
}
