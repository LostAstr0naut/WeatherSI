package rainfallratesi

import (
	"image/gif"
	"log"
	"net/http"
	"rainfallratesi/internal/location"
	"rainfallratesi/internal/rainfallrate"
)

const (
	dataURL = "http://meteo.arso.gov.si/uploads/probase/www/observ/radar/si0-rm-anim.gif"
	radious = 25
)

// GetRainfallRateLevels returns a rainfall rate level based on parameters. Returns highest area level and highest location level.
func GetRainfallRateLevels(locationName string) (*rainfallrate.Level, *rainfallrate.Level, error) {
	rainfallRateSvc := rainfallrate.New()
	locationSvc := location.New()

	foundLocation, err := locationSvc.GetCoordinates(locationName)
	if err != nil {
		log.Println(err)
		return &rainfallrate.Level{}, &rainfallrate.Level{}, err
	}

	xLocation := int(foundLocation.X)
	yLocation := int(foundLocation.Y)
	x1 := int(xLocation - radious)
	y1 := int(yLocation - radious)
	x2 := int(xLocation + radious)
	y2 := int(yLocation + radious)

	resp, err := http.Get(dataURL)
	if err != nil {
		log.Println(err)
		return &rainfallrate.Level{}, &rainfallrate.Level{}, err
	}

	dataGif, err := gif.DecodeAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	dataImages := dataGif.Image

	highestAreaRateLevel := &rainfallrate.Level{}
	highestLocationRateLevel := &rainfallrate.Level{}
	for _, item := range dataImages {
		if item != nil {
			for y := y1; y <= y2; y++ {
				for x := x1; x <= x2; x++ {
					r, g, b, _ := item.At(x, y).RGBA()
					level, err := rainfallRateSvc.GetLevelByRGBA(uint16(r), uint16(g), uint16(b))
					if err != nil {
						log.Println(err)
						return &rainfallrate.Level{}, &rainfallrate.Level{}, err
					}
					if x == xLocation && y == yLocation && highestLocationRateLevel.Value < level.Value {
						highestLocationRateLevel = level
					}
					if highestAreaRateLevel.Value < level.Value {
						highestAreaRateLevel = level
					}
				}
			}
		}
	}
	return highestAreaRateLevel, highestLocationRateLevel, nil
}
