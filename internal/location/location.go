package location

import (
	"errors"
)

// Supported locations.
const (
	Bohinj        = "BO"
	Gorica        = "GO"
	Koper         = "KP"
	Idrija        = "ID"
	Jesenice      = "JE"
	Postojna      = "PO"
	Kranj         = "KR"
	Ljubljana     = "LJ"
	Kocevje       = "KO"
	Trbovlje      = "TB"
	SlovenjGradec = "SG"
	NovoMesto     = "NM"
	Celje         = "CE"
	Krsko         = "KK"
	Maribor       = "MB"
	Ptuj          = "PT"
	MurskaSobota  = "MS"
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
		Bohinj: {
			X: 224,
			Y: 290,
		},
		Gorica: {
			X: 237,
			Y: 374,
		},
		Koper: {
			X: 249,
			Y: 464,
		},
		Idrija: {
			X: 296,
			Y: 364,
		},
		Jesenice: {
			X: 300,
			Y: 268,
		},
		Postojna: {
			X: 325,
			Y: 417,
		},
		Kranj: {
			X: 346,
			Y: 311,
		},
		Ljubljana: {
			X: 373,
			Y: 350,
		},
		Kocevje: {
			X: 424,
			Y: 445,
		},
		Trbovlje: {
			X: 454,
			Y: 332,
		},
		SlovenjGradec: {
			X: 459,
			Y: 252,
		},
		NovoMesto: {
			X: 474,
			Y: 410,
		},
		Celje: {
			X: 484,
			Y: 310,
		},
		Krsko: {
			X: 521,
			Y: 374,
		},
		Maribor: {
			X: 545,
			Y: 247,
		},
		Ptuj: {
			X: 581,
			Y: 270,
		},
		MurskaSobota: {
			X: 625,
			Y: 216,
		},
	}
}
