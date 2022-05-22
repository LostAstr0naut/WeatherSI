package arsoweatherimage

import (
	"errors"
	"image"
	"image/gif"
	"net/http"

	"github.com/LostAstr0naut/arsoweatherimage/internal/location"
	"github.com/LostAstr0naut/arsoweatherimage/internal/rainfallrate"
)

const (
	// Supported location names.
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

	// image resource URL
	dataURL = "http://meteo.arso.gov.si/uploads/probase/www/observ/radar/si0-rm-anim.gif"

	// radious represents the square area that will be scanned
	radious = 25

	// radiousInner represents the square area within radious that will be scanned
	radiousInner = 5
)

type RainfallRateLevel struct {
	Value       float64
	Description string
}

// RainfallRate returns rainfall rate levels based on radious and radiousInner parameters.
// Returns rainfall rate on direct location and in general area.
func RainfallRate(locationName string) (RainfallRateLevel, RainfallRateLevel, error) {
	if len(locationName) < 1 {
		return RainfallRateLevel{}, RainfallRateLevel{}, errors.New("invalid location name")
	}

	foundLocation, err := location.LocationCoordinates(locationName)
	if err != nil {
		return RainfallRateLevel{}, RainfallRateLevel{}, err
	}

	dataImages, err := rainfallRateImages(dataURL)
	if err != nil {
		return RainfallRateLevel{}, RainfallRateLevel{}, err
	}

	return locationRainfallRate(foundLocation, dataImages, radious, radiousInner)
}

func rainfallRateImages(dataURL string) ([]*image.Paletted, error) {
	resp, err := http.Get(dataURL)
	if err != nil {
		return []*image.Paletted{}, err
	}

	dataGif, err := gif.DecodeAll(resp.Body)
	if err != nil {
		return []*image.Paletted{}, err
	}

	return dataGif.Image, nil
}

func locationRainfallRate(
	location location.Location,
	dataImages []*image.Paletted,
	radious, radiousInner int,
) (RainfallRateLevel, RainfallRateLevel, error) {
	xLocation := int(location.X)
	yLocation := int(location.Y)

	// general area coordinates
	x1 := int(xLocation - radious)
	y1 := int(yLocation - radious)
	x2 := int(xLocation + radious)
	y2 := int(yLocation + radious)

	// exact location coordinates
	x1Inner := int(xLocation - radiousInner)
	y1Inner := int(yLocation - radiousInner)
	x2Inner := int(xLocation + radiousInner)
	y2Inner := int(yLocation + radiousInner)

	highestInAreaRateLevel := rainfallrate.Level{}
	highestOnLocationRateLevel := rainfallrate.Level{}
	for _, item := range dataImages {
		if item != nil {
			for y := y1; y <= y2; y++ {
				for x := x1; x <= x2; x++ {
					r, g, b, _ := item.At(x, y).RGBA()
					level, err := rainfallrate.LevelByRGBA(uint16(r), uint16(g), uint16(b))
					if err != nil {
						continue
					}
					if x >= x1Inner &&
						y >= y1Inner &&
						x <= x2Inner &&
						y <= y2Inner &&
						highestOnLocationRateLevel.Value < level.Value {
						highestOnLocationRateLevel = level
					}
					if highestInAreaRateLevel.Value < level.Value {
						highestInAreaRateLevel = level
					}
				}
			}
		}
	}

	return RainfallRateLevel{
			Value:       highestOnLocationRateLevel.Value,
			Description: highestOnLocationRateLevel.Description,
		},
		RainfallRateLevel{
			Value:       highestInAreaRateLevel.Value,
			Description: highestInAreaRateLevel.Description,
		},
		nil
}
