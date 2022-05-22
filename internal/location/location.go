package location

import (
	"errors"
)

// Supported location names.
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
	Name string
	X    int32
	Y    int32
}

func LocationCoordinates(locationName string) (Location, error) {
	for _, location := range locations() {
		if location.Name == locationName {
			return location, nil
		}
	}

	return Location{}, errors.New("location not found")
}

func locations() []Location {
	return []Location{
		{
			Name: Bohinj,
			X:    224,
			Y:    290,
		},
		{
			Name: Gorica,
			X:    237,
			Y:    374,
		},
		{
			Name: Koper,
			X:    249,
			Y:    464,
		},
		{
			Name: Idrija,
			X:    296,
			Y:    364,
		},
		{
			Name: Jesenice,
			X:    300,
			Y:    268,
		},
		{
			Name: Postojna,
			X:    325,
			Y:    417,
		},
		{
			Name: Kranj,
			X:    346,
			Y:    311,
		},
		{
			Name: Ljubljana,
			X:    373,
			Y:    350,
		},
		{
			Name: Kocevje,
			X:    424,
			Y:    445,
		},
		{
			Name: Trbovlje,
			X:    454,
			Y:    332,
		},
		{
			Name: SlovenjGradec,
			X:    459,
			Y:    252,
		},
		{
			Name: NovoMesto,
			X:    474,
			Y:    410,
		},
		{
			Name: Celje,
			X:    484,
			Y:    310,
		},
		{
			Name: Krsko,
			X:    521,
			Y:    374,
		},
		{
			Name: Maribor,
			X:    545,
			Y:    247,
		},
		{
			Name: Ptuj,
			X:    581,
			Y:    270,
		},
		{
			Name: MurskaSobota,
			X:    625,
			Y:    216,
		},
	}
}
