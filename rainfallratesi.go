package rainfallratesi

import (
	"image/gif"
	"net/http"
	"rainfallratesi/internal/location"
	"rainfallratesi/internal/rainfallrate"
)

const (
	dataURL      = "http://meteo.arso.gov.si/uploads/probase/www/observ/radar/si0-rm-anim.gif"
	radious      = 25
	radiousInner = 5
)

// GetRainfallRateLevels returns a rainfall rate level based on parameters. Returns highest area level and highest location level.
func GetRainfallRateLevels(locationName string) (rainfallrate.Level, rainfallrate.Level, error) {
	rainfallRateSvc := rainfallrate.New()
	locationSvc := location.New()

	foundLocation, err := locationSvc.GetCoordinates(locationName)
	if err != nil {
		return rainfallrate.Level{}, rainfallrate.Level{}, err
	}

	xLocation := int(foundLocation.X)
	yLocation := int(foundLocation.Y)
	x1 := int(xLocation - radious)
	y1 := int(yLocation - radious)
	x2 := int(xLocation + radious)
	y2 := int(yLocation + radious)
	x1Inner := int(xLocation - radiousInner)
	y1Inner := int(yLocation - radiousInner)
	x2Inner := int(xLocation + radiousInner)
	y2Inner := int(yLocation + radiousInner)

	resp, err := http.Get(dataURL)
	if err != nil {
		return rainfallrate.Level{}, rainfallrate.Level{}, err
	}

	dataGif, err := gif.DecodeAll(resp.Body)
	if err != nil {
	}
	dataImages := dataGif.Image

	highestAreaRateLevel := rainfallrate.Level{}
	highestLocationRateLevel := rainfallrate.Level{}
	for _, item := range dataImages {
		if item != nil {
			for y := y1; y <= y2; y++ {
				for x := x1; x <= x2; x++ {
					r, g, b, _ := item.At(x, y).RGBA()
					level, err := rainfallRateSvc.GetLevelByRGBA(uint16(r), uint16(g), uint16(b))
					if err != nil {
						continue
					}
					if x >= x1Inner &&
						y >= y1Inner &&
						x <= x2Inner &&
						y <= y2Inner &&
						highestLocationRateLevel.Value < level.Value {
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
